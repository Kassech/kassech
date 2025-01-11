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
	userRoleRepo := &repository.UserRoleRepository{}
	sessionRepo := &repository.UserSessionsRepository{}
	stationRepo := &repository.StationRepository{}
	routeRepo := &repository.RouteRepository{}

	// Initialize Services
	userSvc := &service.UserService{Repo: userRepo}
	roleSvc := &service.RoleService{Repo: roleRepo}
	permissionSvc := &service.PermissionService{Repo: permissionRepo}
	rolePermissionSvc := &service.RolePermissionService{Repo: rolePermissionRepo}
	userRoleSvc := &service.UserRoleService{Repo: userRoleRepo}
	sessionService := &service.UserSessionsService{Repo: sessionRepo}
	stationSvc := &service.StationService{Repo: stationRepo}
	routeSvc := &service.RouteService{Repo: routeRepo}
	// Initialize Controllersz
	userCtrl := &controller.UserController{Service: userSvc, SessionService: sessionService}
	roleCtrl := &controller.RoleController{Service: roleSvc}
	permissionCtrl := &controller.PermissionController{Service: permissionSvc}
	rolePermissionCtrl := &controller.RolePermissionController{Service: rolePermissionSvc}
	userRoleCtrl := &controller.UserRoleController{Service: userRoleSvc}
	sessionController := &controller.UserSessionController{Service: sessionService}
	stationCtrl := &controller.StationController{Service: stationSvc}
	routeCtrl := &controller.RouteController{Service: routeSvc}

	// Group API routes
	api := r.Group("/api")
	{
		// User-related routes
		registerUserRoutes(api, userCtrl)

		// Role-related routes
		registerRoleRoutes(api, roleCtrl)

		// Permission-related routes
		RegisterUserRoleRoutes(api, userRoleCtrl)

		RegisterUserSessionRoutes(api, sessionController)

		// Permission-related routes
		registerPermissionRoutes(api, permissionCtrl)

		// Role Permission-related routes
		registerRolePermissionRoutes(api, rolePermissionCtrl)

		// station routes
		stationRoutes(api, stationCtrl)

		// route routes
		routeRoutes(api, routeCtrl)

	}
}

func RegisterUserSessionRoutes(router *gin.RouterGroup, controller *controller.UserSessionController) {
	sessions := router.Group("/sessions")
	{
		sessions.DELETE("/:token", controller.InvalidateToken) // Invalidate a specific token

		sessions.DELETE("/all/:id", controller.InvalidateAllSessions) // Invalidate all sessions for a user
	}
}

func stationRoutes(api *gin.RouterGroup, ctrl *controller.StationController) {

	stationApi := api.Group("/station")
	{

		stationApi.POST("/", ctrl.CreateStation)

		stationApi.PUT("/:id", ctrl.UpdateStation)

		stationApi.DELETE("/:id", ctrl.DeleteStation)

		stationApi.GET("/:id", ctrl.FindStationByID)

		stationApi.GET("/", ctrl.GetAllStations)

		stationApi.GET("", ctrl.GetAllStations)
	}
}

func routeRoutes(api *gin.RouterGroup, ctrl *controller.RouteController) {

	routeApi := api.Group("/routes")
	{

		routeApi.POST("/", ctrl.CreateRoute)      // Create a route
		routeApi.GET("/:id", ctrl.FindRouteByID)  // Get route by ID
		routeApi.GET("/", ctrl.GetAllRoutes)      // Get all routes
		routeApi.PUT("/:id", ctrl.UpdateRoute)    // Update a route by ID
		routeApi.DELETE("/:id", ctrl.DeleteRoute) // Delete a route by ID
	}
}

func registerUserRoutes(api *gin.RouterGroup, ctrl *controller.UserController) {

	api.POST("/register", ctrl.Register)

	api.POST("/login", ctrl.Login)

	api.POST("/refresh", ctrl.RefreshToken)

	api.POST("/validate", ctrl.VerifyAuth)

	api.Use(middleware.AuthMiddleware())

	api.POST("/logout", ctrl.Logout)

	api.POST("/notification", ctrl.SaveNotificationToken)

	api.POST("/users", ctrl.CreateUser)

	api.GET("/users", ctrl.ListUsers)

	api.PUT("/users/:id", ctrl.UpdateUser)

	api.DELETE("/users/:id", ctrl.DeleteUser)

}

func RegisterUserRoleRoutes(api *gin.RouterGroup, ctrl *controller.UserRoleController) {
	userRoleApi := api.Group("/user-roles")
	{
		userRoleApi.POST("/", ctrl.CreateUserRole)

		userRoleApi.GET("/:id", ctrl.GetUserRole)

		userRoleApi.PUT("/:id", ctrl.UpdateUserRole)

		userRoleApi.DELETE("/:id", ctrl.DeleteUserRole)

	}
}
func registerRoleRoutes(api *gin.RouterGroup, ctrl *controller.RoleController) {
	roleApi := api.Group("/roles")
	{

		roleApi.POST("/", ctrl.CreateRole)

		roleApi.PUT("/:id", ctrl.UpdateRole)

		roleApi.DELETE("/:id", ctrl.DeleteRole)

		roleApi.GET("/:id", ctrl.FindRoleByID)

		roleApi.GET("/", ctrl.GetAllRoles)
	}
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
