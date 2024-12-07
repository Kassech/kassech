package controller

import (
	models "kassech/backend/pkg/model"
	"kassech/backend/pkg/service"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	Service *service.UserService
}

func (uc *UserController) Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdUser, accessToken, refreshToken, err := uc.Service.Register(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.SetCookie("refresh_token", refreshToken, 60*60*24*30, "/", "", true, true) // Expires in 30 days

	// Return the response with the created user and the token
	c.JSON(http.StatusOK, gin.H{
		"message":     "registration successful",
		"user":        createdUser,
		"accessToken": accessToken,
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
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.SetCookie("refresh_token", refreshToken, 60*60*24*30, "/", "", true, true) // Expires in 30 days

	c.JSON(http.StatusOK, gin.H{
		"user":        user,
		"accessToken": accessToken,
	})
}

func (uc *UserController) RefreshToken(c *gin.Context) {
	// Get the refresh token from the HTTP-only cookie
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Refresh token is missing"})
		return
	}

	// Call the service to refresh the token
	accessToken, err := service.RefreshTokenService(refreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
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

	updatedUser, err := uc.Service.UpdateUser(userId, &user)
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

	err := uc.Service.DeleteUser(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User deleted successfully",
	})
}
