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
	vehicleTypeRepo := &repository.VehicleTypeRepository{}
	vehicleRepo := &repository.VehicleRepository{}
	LocationRepo := &repository.LocationRepository{}
	PathRepo := &repository.PathRepository{}
	QueueManagerPathRepo := &repository.QueueManagerRouteRepository{}
	AnalysisRepo := &repository.AnalysisRepository{}
	DriverDeligationRepo := &repository.DriverDeligationRepository{}

	// Initialize Services
	userSvc := &service.UserService{Repo: userRepo}
	roleSvc := &service.RoleService{Repo: roleRepo}
	permissionSvc := &service.PermissionService{Repo: permissionRepo}
	rolePermissionSvc := &service.RolePermissionService{Repo: rolePermissionRepo}
	userRoleSvc := &service.UserRoleService{Repo: userRoleRepo}
	sessionService := &service.UserSessionsService{Repo: sessionRepo}
	stationSvc := &service.StationService{Repo: stationRepo}
	routeSvc := &service.RouteService{Repo: routeRepo}
	vehicleTypeSvc := &service.VehicleTypeService{Repo: vehicleTypeRepo}
	vehicleSvc := &service.VehicleService{Repo: vehicleRepo}
	locationSvc := &service.LocationService{Repo: LocationRepo}
	pathSvc := &service.PathService{Repo: PathRepo}
	queueManagerPathSvc := &service.QueueManagerRouteService{Repo: QueueManagerPathRepo}
	analysesSvc := &service.AnalysisService{Repo: AnalysisRepo}
	driverDeligationSvc := &service.DriverDeligationService{Repo: DriverDeligationRepo}

	// Initialize Controllersz
	userCtrl := &controller.UserController{Service: userSvc, SessionService: sessionService}
	roleCtrl := &controller.RoleController{Service: roleSvc}
	permissionCtrl := &controller.PermissionController{Service: permissionSvc}
	rolePermissionCtrl := &controller.RolePermissionController{Service: rolePermissionSvc}
	userRoleCtrl := &controller.UserRoleController{Service: userRoleSvc}
	sessionController := &controller.UserSessionController{Service: sessionService}
	stationCtrl := &controller.StationController{Service: stationSvc}
	routeCtrl := &controller.RouteController{Service: routeSvc}
	vehicleTypeCtrl := &controller.VehicleTypeController{Service: vehicleTypeSvc}
	vehicleCtrl := &controller.VehicleController{Service: vehicleSvc}
	locationCtrl := &controller.LocationController{Service: locationSvc}
	pathCtrl := &controller.PathController{Service: pathSvc}
	queueManagerPathCtrl := &controller.QueueManagerRouteController{Service: queueManagerPathSvc}
	analysisCtrl := &controller.AnalysisController{Service: analysesSvc}
	driverDeligationCtrl := &controller.DriverDeligationController{Service: driverDeligationSvc}

	// Group API routes
	api := r.Group("/api")
	{
		analysisRoutes(api, analysisCtrl) // Register user-related routes

		registerUserRoutes(api, userCtrl) // Register user-related routes

		registerRoleRoutes(api, roleCtrl) // Register role-related routes

		RegisterUserRoleRoutes(api, userRoleCtrl) // Register user roles-related routes

		RegisterUserSessionRoutes(api, sessionController) // Register user sessions-related routes

		registerPermissionRoutes(api, permissionCtrl) // Register permission-related routes

		registerRolePermissionRoutes(api, rolePermissionCtrl) // Register role permission-related routes

		stationRoutes(api, stationCtrl) // Register station-related routes

		routeRoutes(api, routeCtrl) // Register route-related routes

		pathRoutes(api, pathCtrl) // Register path-related routes

		vehicleTypeRoutes(api, vehicleTypeCtrl) // Register vehicle type-related routes

		vehicleRoutes(api, vehicleCtrl) // Register vehicle-related routes

		locationRoutes(api, locationCtrl) // Register location-related routes

		queueManagerRouteRoutes(api, queueManagerPathCtrl) // Register queue manager path-related routes

		driverDeligationRoutes(api, driverDeligationCtrl) // Register queue manager path-related routes

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

func pathRoutes(api *gin.RouterGroup, ctrl *controller.PathController) {
	pathApi := api.Group("/path")
	{
		pathApi.GET("/", ctrl.GetAllPaths)                     // Get all paths
		pathApi.POST("/", ctrl.CreatePath)                     // Create a path
		pathApi.GET("/station/:id", ctrl.FindPathsByStationID) // Get path by ID
		pathApi.GET("/:id", ctrl.FindPathByID)                 // Get path by ID
		pathApi.DELETE("/:id", ctrl.DeletePath)                // Delete a path by ID

	}
}
func vehicleRoutes(api *gin.RouterGroup, ctrl *controller.VehicleController) {
	vehicleApi := api.Group("/vehicles")
	{
		vehicleApi.POST("/", ctrl.CreateVehicle)      // Create a vehicle
		vehicleApi.GET("/:id", ctrl.FindVehicleByID)  // Get vehicle by ID
		vehicleApi.GET("/", ctrl.GetAllVehicles)      // Get all vehicles
		vehicleApi.PUT("/status/:id", ctrl.UpdateVehicleStatus)    // Update a vehicle by ID
		vehicleApi.PUT("/:id", ctrl.UpdateVehicle)    // Update a vehicle by ID
		vehicleApi.DELETE("/:id", ctrl.DeleteVehicle) // Delete a vehicle by ID

	}
}

func vehicleTypeRoutes(api *gin.RouterGroup, ctrl *controller.VehicleTypeController) {
	vehicleTypeApi := api.Group("/vehicle_types")
	{
		vehicleTypeApi.POST("/", ctrl.CreateVehicleType)      // Create a vehicle type
		vehicleTypeApi.GET("/:id", ctrl.FindVehicleTypeByID)  // Get vehicle type by ID
		vehicleTypeApi.GET("/", ctrl.GetAllVehicleTypes)      // Get all vehicle types
		vehicleTypeApi.PUT("/:id", ctrl.UpdateVehicleType)    // Update a vehicle type by ID
		vehicleTypeApi.DELETE("/:id", ctrl.DeleteVehicleType) // Delete a vehicle type by ID
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

	api.GET("/users/:id", ctrl.GetUserById)

	api.GET("/users/me", ctrl.GetCurrntUser)

	api.GET("/users/verify/:id", ctrl.VerifyUser)

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

func analysisRoutes(api *gin.RouterGroup, ctrl *controller.AnalysisController) {
	analysisApi := api.Group("/analysis")
	{
		// Total number of users
		analysisApi.GET("/total-users", ctrl.GetTotalUsers)

		// Total number of active users
		analysisApi.GET("/active-users", ctrl.GetActiveUsers)

		// Total number of drivers
		analysisApi.GET("/total-drivers", ctrl.GetTotalDrivers)

		// Total number of vehicles
		analysisApi.GET("/total-vehicles", ctrl.GetTotalVehicles)

		// Total number of active vehicles
		analysisApi.GET("/active-vehicles", ctrl.GetActiveVehicles)

		// Total travel logs
		analysisApi.GET("/total-travel-logs", ctrl.GetTotalTravelLogs)

		// Total number of routes
		analysisApi.GET("/total-routes", ctrl.GetTotalRoutes)

		// Total number of stations
		analysisApi.GET("/total-stations", ctrl.GetTotalStations)

		// Total number of login logs
		analysisApi.GET("/login-logs", ctrl.GetLoginLogs)

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

func locationRoutes(api *gin.RouterGroup, ctrl *controller.LocationController) {
	roleApi := api.Group("/location")
	{

		roleApi.POST("/", ctrl.CreateLocation)

		roleApi.DELETE("/:id", ctrl.DeleteLocation)

		roleApi.GET("/:id", ctrl.FindLocationByID)

		roleApi.GET("/", ctrl.GetAllLocations)
	}
}
func queueManagerRouteRoutes(router *gin.RouterGroup, ctrl *controller.QueueManagerRouteController) {
	routes := router.Group("/queue-manager-routes")
	{
		routes.POST("/", ctrl.CreateRoute)
		routes.GET("/", ctrl.GetAllRoutes)
		routes.GET("/:id", ctrl.GetRoute)
		routes.DELETE("/:id", ctrl.DeleteRoute)
	}
}
func driverDeligationRoutes(router *gin.RouterGroup, ctrl *controller.DriverDeligationController) {
	routes := router.Group("/driver-deligation")
	{
		routes.POST("/assign", ctrl.AssignDriver)
		routes.GET("/active-driver/:vehicleId", ctrl.GetActiveDriver)
	}
}
