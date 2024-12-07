package http

import (
	"kassech/backend/pkg/controller"
	"kassech/backend/pkg/middleware"
	"kassech/backend/pkg/repository"
	"kassech/backend/pkg/service"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	// Initialize Repositories
	userRepo := &repository.UserRepository{}
	roleRepo := &repository.RoleRepository{}
	permissionRepo := &repository.PermissionRepository{}
	rolePermissionRepo := &repository.RolePermissionRepository{}

	// Initialize Services
	userSvc := &service.UserService{Repo: userRepo}
	roleSvc := &service.RoleService{Repo: roleRepo}
	permissionSvc := service.NewPermissionService(permissionRepo)
	rolePermissionSvc := service.NewRolePermissionService(rolePermissionRepo)

	// Initialize Controllers
	userCtrl := &controller.UserController{Service: userSvc}
	roleCtrl := &controller.RoleController{Service: roleSvc}
	permissionCtrl := controller.NewPermissionController(permissionSvc)
	rolePermissionCtrl := controller.NewRolePermissionController(rolePermissionSvc)

	// Group API routes
	api := r.Group("/api")
	{
		// User-related routes
		registerUserRoutes(api, userCtrl)

		// Role-related routes
		registerRoleRoutes(api, roleCtrl)

		// Permission-related routes
		registerPermissionRoutes(api, permissionCtrl)

		// Role Permission-related routes
		registerRolePermissionRoutes(api, rolePermissionCtrl)
	}
}

func registerUserRoutes(api *gin.RouterGroup, ctrl *controller.UserController) {

	api.POST("/register", ctrl.Register)

	api.POST("/login", ctrl.Login)

	api.POST("/refresh", ctrl.RefreshToken)

	api.POST("/validate", ctrl.VerifyAuth)

	api.Use(middleware.AuthMiddleware())
	api.GET("/users", ctrl.ListUsers)

	api.POST("/users", ctrl.CreateUser)
	
	api.PUT("/users/:id", ctrl.UpdateUser)

	api.DELETE("/users/:id", ctrl.DeleteUser)

}

func registerRoleRoutes(api *gin.RouterGroup, ctrl *controller.RoleController) {
	api.POST("/roles", ctrl.CreateRole)

	api.PUT("/roles/:id", ctrl.UpdateRole)

	api.DELETE("/roles/:id", ctrl.DeleteRole)

	api.GET("/roles/:id", ctrl.FindRoleByID)

	api.GET("/roles", ctrl.GetAllRoles)
}

func registerPermissionRoutes(api *gin.RouterGroup, ctrl *controller.PermissionController) {
	api.POST("/permissions", ctrl.CreatePermission)

	api.GET("/permissions", ctrl.GetAllPermissions)

	api.GET("/permissions/:id", ctrl.GetPermissionByID)

	api.PUT("/permissions/:id", ctrl.UpdatePermission)

	api.DELETE("/permissions/:id", ctrl.DeletePermission)

	api.POST("/permissions/permission/:permission_id/roles/:role_id", ctrl.AttachRoleToPermission)

	api.DELETE("/permissions/permission/:permission_id/roles/:role_id", ctrl.DetachRoleFromPermission)
}

func registerRolePermissionRoutes(api *gin.RouterGroup, ctrl *controller.RolePermissionController) {
	api.POST("/role_permissions", ctrl.CreateRolePermission)

	api.GET("/role_permissions", ctrl.GetAllRolePermissions)

	api.GET("/role_permissions/:id", ctrl.GetRolePermissionByID)

	api.DELETE("/role_permissions/:id", ctrl.DeleteRolePermission)

	api.DELETE("/role_permissions/role/:role_id/permission/:permission_id", ctrl.DeleteRolePermissionByRoleAndPermission)
}
