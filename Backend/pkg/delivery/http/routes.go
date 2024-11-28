package http

import (
	"kassech/backend/pkg/controller"
	"kassech/backend/pkg/repository"
	"kassech/backend/pkg/service"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	userRepo := &repository.UserRepository{}
	userService := &service.UserService{Repo: userRepo}
	userController := &controller.UserController{Service: userService}

	api := r.Group("/api")
	{
		api.POST("/register", userController.Register)
		api.POST("/login", userController.Login)
	}
}
