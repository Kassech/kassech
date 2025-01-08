package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"kassech/backend/pkg/constants"
	"kassech/backend/pkg/domain"
	"kassech/backend/pkg/mapper"
	models "kassech/backend/pkg/model"
	"kassech/backend/pkg/service"
	"kassech/backend/pkg/utils"
	"log"
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
	var user models.User

	// Read the request body
	body, _ := io.ReadAll(c.Request.Body)
	// Extract and upload profile picture
	file, _, err := c.Request.FormFile("profile_picture")
	fmt.Println("file:", file)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Profile picture is required"})
		return
	}
	defer file.Close()

	// profilePictureName, err := utils.UploadFile(c.Request, "profile_picture", constants.ProfilePictureDirectory)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload profile picture"})
	// 	return
	// }

	// // Assign the profile picture location to the user
	// user.ProfilePicture = &profilePictureName

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Parse the body to extract the role
	var requestBody map[string]interface{}
	if err := json.Unmarshal(body, &requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	// Check if "role" exists and is valid
	roleFloat, ok := requestBody["role"].(float64)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Role is required and must be a valid number"})
		return
	}

	// Convert the role to uint
	role := uint(roleFloat)

	// Rebind the body for user struct parsing
	if err := json.Unmarshal(body, &user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the service with the user and role
	createdUser, accessToken, refreshToken, err := uc.Service.Register(&user, role)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set the refresh token cookie
	c.SetCookie("refresh_token", refreshToken, 60*60*24*30, "/", "", true, true) // Expires in 30 days
	uc.SessionService.CreateSession(user.ID, refreshToken, time.Now().Add(service.RefreshTokenExpiration))

	// Return the response with the created user and the token
	c.JSON(http.StatusOK, gin.H{
		"message":     "registration successful",
		"user":        createdUser,
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

	c.JSON(http.StatusOK, gin.H{
		"user":        user,
		"accessToken": accessToken,
	})
}
func (uc *UserController) CreateUser(c *gin.Context) {

	var user domain.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Save the profile picture file
	if user.ProfilePictureFile != nil {
		profilePicturePath, err := utils.UploadFile(user.ProfilePictureFile, constants.ProfilePictureDirectory)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save profile picture"})
			return
		}
		user.ProfilePicture = &profilePicturePath
	}
	// make user is verified to true
	user.IsVerified = true
	// Convert the domain user to GORM user
	var userModel *models.User = mapper.ToGormUser(&user)

	insertedUser, err := uc.Service.CreateUser(userModel, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
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
	accessToken, err := service.RefreshTokenService(refreshToken)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
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
