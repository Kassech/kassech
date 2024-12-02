package http

import (
	"kassech/backend/pkg/controller"
	"kassech/backend/pkg/middleware"
	"kassech/backend/pkg/repository"
	"kassech/backend/pkg/service"

	_ "kassech/backend/docs"

	swaggerfiles "github.com/swaggo/files"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// RegisterRoutes initializes and registers all HTTP routes.
// @Summary Initialize routes
// @Description This function initializes and registers all the API routes for users, roles, permissions, and role permissions.
// @Tags Initialization
// @Router / [get]
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

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
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

// registerUserRoutes defines the routes for user-related operations.
// @Summary Register user routes
// @Description Defines all the routes related to user operations such as registration, login, token refresh, etc.
// @Tags User
// @Router /api/user [post]
func registerUserRoutes(api *gin.RouterGroup, ctrl *controller.UserController) {
	// User registration route
	// @Summary Register a new user
	// @Description Register a new user by providing necessary details
	// @Tags User
	// @Accept json
	// @Produce json
	// @Param user body controller.RegisterUserRequest true "User registration request"
	// @Success 201 {object} controller.UserResponse
	// @Failure 400 {object} controller.ErrorResponse
	// @Router /api/register [post]
	api.POST("/register", ctrl.Register)

	// User login route
	// @Summary Login user
	// @Description Login using username and password to get an authentication token
	// @Tags User
	// @Accept json
	// @Produce json
	// @Param login body controller.LoginRequest true "Login request"
	// @Success 200 {object} controller.LoginResponse
	// @Failure 401 {object} controller.ErrorResponse
	// @Router /api/login [post]
	api.POST("/login", ctrl.Login)

	// Refresh token route
	// @Summary Refresh authentication token
	// @Description Refresh an expired authentication token using a valid refresh token
	// @Tags User
	// @Produce json
	// @Success 200 {object} controller.LoginResponse
	// @Failure 400 {object} controller.ErrorResponse
	// @Router /api/refresh [post]
	api.POST("/refresh", ctrl.RefreshToken)

	// Validate user authentication
	// @Summary Validate user authentication
	// @Description Validate if the user is authenticated based on the current token
	// @Tags User
	// @Produce json
	// @Success 200 {string} string "Authenticated"
	// @Failure 401 {object} controller.ErrorResponse
	// @Router /api/validate [post]
	api.POST("/validate", ctrl.VerifyAuth)

	// Middleware to check auth before user-related routes
	api.Use(middleware.AuthMiddleware())

	// User-related route (example)
	// @Summary Get user information
	// @Description Fetch user data based on token
	// @Tags User
	// @Success 200 {object} controller.UserResponse
	// @Failure 400 {object} controller.ErrorResponse
	// @Router /api/user [post]
	api.POST("/user", ctrl.Login)
}

// registerRoleRoutes defines the routes for role-related operations.
// @Summary Register role routes
// @Description Defines all the routes related to role operations such as creating, updating, and deleting roles.
// @Tags Role
// @Router /api/roles [post]
func registerRoleRoutes(api *gin.RouterGroup, ctrl *controller.RoleController) {
	// Create a new role
	// @Summary Create a new role
	// @Description Create a new role with specific permissions
	// @Tags Role
	// @Accept json
	// @Produce json
	// @Param role body controller.CreateRoleRequest true "Role creation request"
	// @Success 201 {object} controller.RoleResponse
	// @Failure 400 {object} controller.ErrorResponse
	// @Router /api/roles [post]
	api.POST("/roles", ctrl.CreateRole)

	// Update an existing role
	// @Summary Update a role
	// @Description Update a role by ID
	// @Tags Role
	// @Accept json
	// @Produce json
	// @Param id path int true "Role ID"
	// @Param role body controller.UpdateRoleRequest true "Role update request"
	// @Success 200 {object} controller.RoleResponse
	// @Failure 400 {object} controller.ErrorResponse
	// @Router /api/roles/{id} [put]
	api.PUT("/roles/:id", ctrl.UpdateRole)

	// Delete a role
	// @Summary Delete a role
	// @Description Delete a role by ID
	// @Tags Role
	// @Param id path int true "Role ID"
	// @Success 204 {object} controller.NoContentResponse
	// @Failure 404 {object} controller.ErrorResponse
	// @Router /api/roles/{id} [delete]
	api.DELETE("/roles/:id", ctrl.DeleteRole)

	// Get role by ID
	// @Summary Get role by ID
	// @Description Fetch a specific role by its ID
	// @Tags Role
	// @Param id path int true "Role ID"
	// @Success 200 {object} controller.RoleResponse
	// @Failure 404 {object} controller.ErrorResponse
	// @Router /api/roles/{id} [get]
	api.GET("/roles/:id", ctrl.FindRoleByID)

	// Get all roles
	// @Summary Get all roles
	// @Description Fetch a list of all roles
	// @Tags Role
	// @Success 200 {array} controller.RoleResponse
	// @Failure 400 {object} controller.ErrorResponse
	// @Router /api/roles [get]
	api.GET("/roles", ctrl.GetAllRoles)
}

// registerPermissionRoutes defines the routes for permission-related operations.
// @Summary Register permission routes
// @Description Defines all the routes related to permission operations such as creating, updating, deleting, and associating roles with permissions.
// @Tags Permission
// @Router /api/permissions [post]
func registerPermissionRoutes(api *gin.RouterGroup, ctrl *controller.PermissionController) {
	// Create a new permission
	// @Summary Create a new permission
	// @Description Create a new permission to assign to roles
	// @Tags Permission
	// @Accept json
	// @Produce json
	// @Param permission body controller.CreatePermissionRequest true "Permission creation request"
	// @Success 201 {object} controller.PermissionResponse
	// @Failure 400 {object} controller.ErrorResponse
	// @Router /api/permissions [post]
	api.POST("/permissions", ctrl.CreatePermission)

	// Get all permissions
	// @Summary Get all permissions
	// @Description Fetch a list of all available permissions
	// @Tags Permission
	// @Success 200 {array} controller.PermissionResponse
	// @Failure 400 {object} controller.ErrorResponse
	// @Router /api/permissions [get]
	api.GET("/permissions", ctrl.GetAllPermissions)

	// Get permission by ID
	// @Summary Get permission by ID
	// @Description Fetch a specific permission by its ID
	// @Tags Permission
	// @Param id path int true "Permission ID"
	// @Success 200 {object} controller.PermissionResponse
	// @Failure 404 {object} controller.ErrorResponse
	// @Router /api/permissions/{id} [get]
	api.GET("/permissions/:id", ctrl.GetPermissionByID)

	// Update a permission
	// @Summary Update a permission
	// @Description Update the details of a permission by its ID
	// @Tags Permission
	// @Accept json
	// @Produce json
	// @Param id path int true "Permission ID"
	// @Param permission body controller.UpdatePermissionRequest true "Permission update request"
	// @Success 200 {object} controller.PermissionResponse
	// @Failure 400 {object} controller.ErrorResponse
	// @Router /api/permissions/{id} [put]
	api.PUT("/permissions/:id", ctrl.UpdatePermission)

	// Delete a permission
	// @Summary Delete a permission
	// @Description Remove a permission by ID
	// @Tags Permission
	// @Param id path int true "Permission ID"
	// @Success 204 {object} controller.NoContentResponse
	// @Failure 404 {object} controller.ErrorResponse
	// @Router /api/permissions/{id} [delete]
	api.DELETE("/permissions/:id", ctrl.DeletePermission)

	// Attach a role to a permission
	// @Summary Attach a role to a permission
	// @Description Attach a role to a specific permission
	// @Tags Permission
	// @Param permission_id path int true "Permission ID"
	// @Param role_id path int true "Role ID"
	// @Success 200 {object} controller.SuccessResponse
	// @Failure 400 {object} controller.ErrorResponse
	// @Router /api/permissions/permission/{permission_id}/roles/{role_id} [post]
	api.POST("/permissions/permission/:permission_id/roles/:role_id", ctrl.AttachRoleToPermission)

	// Detach a role from a permission
	// @Summary Detach a role from a permission
	// @Description Detach a role from a specific permission
	// @Tags Permission
	// @Param permission_id path int true "Permission ID"
	// @Param role_id path int true "Role ID"
	// @Success 204 {object} controller.NoContentResponse
	// @Failure 400 {object} controller.ErrorResponse
	// @Router /api/permissions/permission/{permission_id}/roles/{role_id} [delete]
	api.DELETE("/permissions/permission/:permission_id/roles/:role_id", ctrl.DetachRoleFromPermission)
}

// registerRolePermissionRoutes defines the routes for role-permission-related operations.
// @Summary Register role permission routes
// @Description Defines all the routes related to managing role-permission associations.
// @Tags RolePermission
// @Router /api/role_permissions [post]
func registerRolePermissionRoutes(api *gin.RouterGroup, ctrl *controller.RolePermissionController) {
	// Create a new role-permission association
	// @Summary Create a role-permission association
	// @Description Associate a role with a permission
	// @Tags RolePermission
	// @Accept json
	// @Produce json
	// @Param role_permission body controller.CreateRolePermissionRequest true "Role permission creation request"
	// @Success 201 {object} controller.RolePermissionResponse
	// @Failure 400 {object} controller.ErrorResponse
	// @Router /api/role_permissions [post]
	api.POST("/role_permissions", ctrl.CreateRolePermission)

	// Get all role-permission associations
	// @Summary Get all role-permission associations
	// @Description Fetch a list of all role-permission associations
	// @Tags RolePermission
	// @Success 200 {array} controller.RolePermissionResponse
	// @Failure 400 {object} controller.ErrorResponse
	// @Router /api/role_permissions [get]
	api.GET("/role_permissions", ctrl.GetAllRolePermissions)

	// Get role-permission by ID
	// @Summary Get role-permission by ID
	// @Description Fetch a specific role-permission association by its ID
	// @Tags RolePermission
	// @Param id path int true "Role Permission ID"
	// @Success 200 {object} controller.RolePermissionResponse
	// @Failure 404 {object} controller.ErrorResponse
	// @Router /api/role_permissions/{id} [get]
	api.GET("/role_permissions/:id", ctrl.GetRolePermissionByID)

	// Delete role-permission by ID
	// @Summary Delete role-permission by ID
	// @Description Delete a specific role-permission association by its ID
	// @Tags RolePermission
	// @Param id path int true "Role Permission ID"
	// @Success 204 {object} controller.NoContentResponse
	// @Failure 404 {object} controller.ErrorResponse
	// @Router /api/role_permissions/{id} [delete]
	api.DELETE("/role_permissions/:id", ctrl.DeleteRolePermission)

	// Delete role-permission by role and permission
	// @Summary Delete role-permission by role and permission
	// @Description Remove a role-permission association by role and permission IDs
	// @Tags RolePermission
	// @Param role_id path int true "Role ID"
	// @Param permission_id path int true "Permission ID"
	// @Success 204 {object} controller.NoContentResponse
	// @Failure 404 {object} controller.ErrorResponse
	// @Router /api/role_permissions/role/{role_id}/permission/{permission_id} [delete]
	api.DELETE("/role_permissions/role/:role_id/permission/:permission_id", ctrl.DeleteRolePermissionByRoleAndPermission)
}
