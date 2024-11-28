package controller

import (
	models "kassech/backend/pkg/model"
	"kassech/backend/pkg/service"
	"net/http"

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

	createdUser, token, err := uc.Service.Register(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Return the response with the created user and the token
	c.JSON(http.StatusOK, gin.H{
		"message": "registration successful",
		"user":    createdUser,
		"token":   token,
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

	user, token, err := uc.Service.Login(input.EmailOrPhone, input.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user":  user,
		"token": token,
	})
}
