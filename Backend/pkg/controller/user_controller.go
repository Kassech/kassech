package controller

import (
	"fmt"
	"kassech/backend/pkg/constants"
	"kassech/backend/pkg/database"
	"kassech/backend/pkg/domain"
	"kassech/backend/pkg/mapper"
	models "kassech/backend/pkg/model"
	"kassech/backend/pkg/service"
	"kassech/backend/pkg/utils"
	"log"
	"mime/multipart"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	Service        *service.UserService
	SessionService *service.UserSessionsService
}

func (uc *UserController) Register(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	uploadAndAssignPath := func(file *multipart.FileHeader, directory string, assign func(path string)) error {
		if file != nil {
			path, err := utils.UploadFile(file, directory)
			if err != nil {
				return err
			}
			assign(path)
		}
		return nil
	}

	if err := uploadAndAssignPath(user.ProfilePictureFile, constants.ProfilePictureDirectory, func(path string) {
		user.ProfilePicture = &path
	}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save profile picture"})
		return
	}

	user.IsVerified = false
	userModel := mapper.ToGormUser(&user)

	insertedUser, err := uc.Service.CreateUser(userModel, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if user.Role == constants.DriverRoleID {
		driverPaths := make(map[string]*string)

		assignPath := func(key string, file *multipart.FileHeader, directory string) error {
			if file != nil {
				path, err := utils.UploadFile(file, directory)
				if err != nil {
					return err
				}
				driverPaths[key] = &path
			}
			return nil
		}

		if err := assignPath("driving_license", user.DrivingLicenseFile, constants.DrivingLicenseDirectory); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save driving license"})
			return
		}

		if err := assignPath("national_id", user.NationalIdFile, constants.NationalIdDirectory); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save national ID"})
			return
		}

		if err := assignPath("insurance_document", user.InsuranceDocumentFile, constants.InsuranceDocumentDirectory); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save insurance document"})
			return
		}

		if err := assignPath("other_document", user.OtherFile, constants.OthersDirectory); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save other document"})
			return
		}

		driver := models.Driver{
			UserID:             insertedUser.ID,
			Status:             "offline",
			DrivingLicensePath: *driverPaths["driving_license"],
			NationalIdPath:     *driverPaths["national_id"],
			InsuranceDocPath:   *driverPaths["insurance_document"],
			OtherFilePath:      *driverPaths["other_document"],
		}

		if _, err := uc.Service.CreateDriver(&driver); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	// Call the service with the user and role
	domainUser, accessToken, refreshToken, err := uc.Service.GenerateAuthentication(insertedUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set the refresh token cookie
	c.SetCookie("refresh_token", refreshToken, 60*60*24*30, "/", "", true, true) // Expires in 30 days
	uc.SessionService.CreateSession(user.ID, refreshToken, time.Now().Add(service.RefreshTokenExpiration))

	// Save the access token to Redis
	redisKey := fmt.Sprintf("session_token:%d", insertedUser.ID)
	err = database.REDIS.Set(c, redisKey, accessToken, service.AccessTokenExpiration).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to store access token"})
		return
	}

	// Return the response with the created user and the token
	c.JSON(http.StatusOK, gin.H{
		"message":     "registration successful",
		"user":        domainUser,
		"accessToken": accessToken,
	})
}

func (uc *UserController) Logout(c *gin.Context) {
	// Get the refresh token from the HTTP-only cookie
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		// If no refresh token is found, return an error response
		c.JSON(http.StatusForbidden, gin.H{"error": "Refresh token is missing"})
		return
	}

	// Invalidate the session using the refresh token (session service)
	err = uc.SessionService.Logout(refreshToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete session"})
		return
	}

	// Remove the refresh token from the client's cookies
	c.SetCookie("refresh_token", "", -1, "/", "", true, true) // Expired immediately to delete

	// Respond with a success message
	c.JSON(http.StatusOK, gin.H{
		"message": "Logout successful",
	})
}

func (uc *UserController) Login(c *gin.Context) {
	var input struct {
		EmailOrPhone string `json:"email_or_phone"`
		Password     string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, accessToken, refreshToken, err := uc.Service.Login(input.EmailOrPhone, input.Password, c.Request)
	if err != nil {
		if err.Error() == "invalid credentials" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	uc.SessionService.CreateSession(user.ID, refreshToken, time.Now().Add(service.RefreshTokenExpiration))
	c.SetCookie("refresh_token", refreshToken, 60*60*24*30, "/", "", true, true) // Expires in 30 days
	database.REDIS.Set(c, fmt.Sprintf("session_token:%d", user.ID), accessToken, service.AccessTokenExpiration)

	c.JSON(http.StatusOK, gin.H{
		"user":        user,
		"accessToken": accessToken,
	})
}

// CreateUser handles the creation of a new user. It binds incoming JSON data to a User object,
// uploads and assigns file paths for various user documents, and saves the user information to the database.
// If successful, it returns a success message and the created user object; otherwise, it returns an error response.
func (uc *UserController) CreateUser(c *gin.Context) {
	var user domain.User
	user.Password = user.PhoneNumber
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	uploadAndAssignPath := func(file *multipart.FileHeader, directory string, assign func(path string)) error {
		if file != nil {
			path, err := utils.UploadFile(file, directory)
			if err != nil {
				return err
			}
			assign(path)
		}
		return nil
	}

	if err := uploadAndAssignPath(user.ProfilePictureFile, constants.ProfilePictureDirectory, func(path string) {
		user.ProfilePicture = &path
	}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save profile picture"})
		return
	}

	user.IsVerified = true
	userModel := mapper.ToGormUser(&user)

	insertedUser, err := uc.Service.CreateUser(userModel, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if user.Role == constants.DriverRoleID {
		driverPaths := make(map[string]*string)

		assignPath := func(key string, file *multipart.FileHeader, directory string) error {
			if file != nil {
				path, err := utils.UploadFile(file, directory)
				if err != nil {
					return err
				}
				driverPaths[key] = &path
			}
			return nil
		}

		if err := assignPath("driving_license", user.DrivingLicenseFile, constants.DrivingLicenseDirectory); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save driving license"})
			return
		}

		if err := assignPath("national_id", user.NationalIdFile, constants.NationalIdDirectory); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save national ID"})
			return
		}

		if err := assignPath("insurance_document", user.InsuranceDocumentFile, constants.InsuranceDocumentDirectory); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save insurance document"})
			return
		}

		if err := assignPath("other_document", user.OtherFile, constants.OthersDirectory); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save other document"})
			return
		}

		driver := models.Driver{
			UserID:             insertedUser.ID,
			Status:             "offline",
			DrivingLicensePath: *driverPaths["driving_license"],
			NationalIdPath:     *driverPaths["national_id"],
			InsuranceDocPath:   *driverPaths["insurance_document"],
			OtherFilePath:      *driverPaths["other_document"],
		}

		if _, err := uc.Service.CreateDriver(&driver); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User created successfully",
		"user":    insertedUser,
	})
}

func (uc *UserController) RefreshToken(c *gin.Context) {
	// Get the refresh token from the HTTP-only cookie
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Refresh token is missing"})
		return
	}

	// Call the service to refresh the token
	accessToken, userId, err := service.RefreshTokenService(refreshToken)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("accessToken:", accessToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Save the access token to Redis
	redisKey := "session_token:" + userId
	storedToken, err := database.REDIS.Get(c, redisKey).Result()
	if err != nil || storedToken != accessToken {
		c.JSON(http.StatusForbidden, gin.H{"error": "Session not found or token mismatch"})
		c.Abort()
		return
	}

	// Send the new access token to the client
	c.JSON(http.StatusOK, gin.H{
		"access_token": accessToken,
	})
}

func (uc *UserController) VerifyAuth(c *gin.Context) {
	token := c.GetHeader("Authorization")
	log.Println(token)

	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token missing"})
		return
	}

	// Remove the "Bearer " prefix, if it exists
	if len(token) > 7 && token[:7] == "Bearer " {
		token = token[7:] // Remove "Bearer " from the token string
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
		return
	}

	// Verify the token (e.g., using a JWT library or custom logic)
	user, err := service.VerifyAccessToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user_id": user.UserID})
}

// SaveNotificationToken method (Save Notification Token with DeviceID and IPAddress)
func (uc *UserController) SaveNotificationToken(c *gin.Context) {
	var input struct {
		Token    string `json:"token"`
		DeviceID string `json:"device_id"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found"})
		c.Abort()
		return
	}
	userIDUint64, _ := strconv.ParseUint(fmt.Sprintf("%v", userID), 10, 32)
	userIDUint := uint(userIDUint64)
	fmt.Printf("🚀 ~ func SaveNotificationToken ~ userIDUint: %d, type: %T\n", userIDUint, userIDUint)

	err := uc.Service.SaveNotificationToken(userIDUint, input.Token, input.DeviceID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Notification token saved successfully"})
}

// ListUsers method (Read with Pagination and Search)
func (uc *UserController) ListUsers(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil || limit < 1 {
		limit = 10
	}

	typ := c.DefaultQuery("type", "active")
	log.Println(typ)
	search := c.DefaultQuery("search", "")

	users, total, err := uc.Service.ListUsers(page, limit, search, typ)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"users":       users,
		"total_count": total,
		"page":        page,
		"limit":       limit,
	})
}

// UpdateUser method (Update User by ID)
func (uc *UserController) UpdateUser(c *gin.Context) {
	userId := c.Param("id")

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userIdUint, err := utils.StringToUint(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	updatedUser, err := uc.Service.UpdateUser(userIdUint, &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User updated successfully",
		"user":    updatedUser,
	})
}

// DeleteUser method (Delete User by ID)
func (uc *UserController) DeleteUser(c *gin.Context) {
	userId := c.Param("id")
	userIdUint, err := utils.StringToUint(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	err = uc.Service.DeleteUser(userIdUint)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User deleted successfully",
	})
}
