package database

import (
	"encoding/json"
	"errors"
	"fmt"
	models "kassech/backend/pkg/model"
	"log"
	"math"
	"os"
	"path/filepath"
	"time"

	"golang.org/x/exp/rand"
	"gorm.io/gorm"
)

// JSON data structures
type ImageEntry struct {
	Image string `json:"image"`
}

type PermissionConfig struct {
	PermissionName string   `json:"permission_name"`
	Description    string   `json:"description"`
	Roles          []string `json:"roles"`
}

type VehicleTypeEntry struct {
	TypeName    string `json:"type_name"`
	Capacity    uint   `json:"capacity"`
	Description string `json:"description"`
}

func generateVIN() string {
	const letters = "ABCDEFGHJKLMNPRSTUVWXYZ"
	const digits = "0123456789"
	vin := make([]byte, 17)

	for i := range vin {
		if i < 3 || i >= 11 {
			vin[i] = letters[rand.Intn(len(letters))]
		} else {
			vin[i] = digits[rand.Intn(len(digits))]
		}
	}
	return string(vin)
}

func SeedDB() {
	if DB == nil {
		log.Fatal("❌ Database connection is nil")
	}

	log.Println("🚀 Starting database seeding process...")

	// Get current working directory
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("❌ Failed to get working directory: %v", err)
	}

	//--------------------------------------------------
	// 1. Seed Roles
	//--------------------------------------------------
	log.Println("🟢 Starting role seeding...")

	roles := []models.Role{
		{RoleName: "Admin", Description: "System administrator"},
		{RoleName: "Driver", Description: "Vehicle driver"},
		{RoleName: "Owner", Description: "Vehicle owner"},
		{RoleName: "QueueManager", Description: "Manages queues"},
		{RoleName: "CustomerService", Description: "Customer support"},
	}

	for _, role := range roles {
		result := DB.FirstOrCreate(&role, "role_name = ?", role.RoleName)
		if result.Error != nil {
			log.Printf("⚠️ Failed to seed role %s: %v", role.RoleName, result.Error)
		} else if result.RowsAffected > 0 {
			log.Printf("✅ Created role: %s", role.RoleName)
		}
	}
	log.Println("🟢 Role seeding completed")

	//--------------------------------------------------
	// 2. Seed Permissions & Role-Permission Mappings
	//--------------------------------------------------
	log.Println("🟢 Starting permission seeding...")

	var permissionConfigs []PermissionConfig
	permFile, err := os.ReadFile(filepath.Join(cwd, "pkg/database/data/permissions.json"))
	if err != nil {
		log.Fatalf("❌ Failed to read permissions file: %v", err)
	}

	if err := json.Unmarshal(permFile, &permissionConfigs); err != nil {
		log.Fatalf("❌ Failed to parse permissions JSON: %v", err)
	}

	// Additional hardcoded permissions from second code
	basicPermissions := []models.Permission{
		{PermissionName: "CreateUser", Description: "Permission to create users"},
		{PermissionName: "DeleteUser", Description: "Permission to delete users"},
		{PermissionName: "UpdateUser", Description: "Permission to update users"},
		{PermissionName: "ViewReports", Description: "Permission to view reports"},
	}

	for _, perm := range basicPermissions {
		if err := DB.FirstOrCreate(&models.Permission{}, perm).Error; err != nil {
			log.Printf("⚠️ Failed to seed basic permission %s: %v", perm.PermissionName, err)
		}
	}

	for _, config := range permissionConfigs {
		perm := models.Permission{
			PermissionName: config.PermissionName,
			Description:    config.Description,
		}

		result := DB.FirstOrCreate(&perm, "permission_name = ?", config.PermissionName)
		if result.Error != nil {
			log.Printf("⚠️ Failed to seed permission %s: %v", config.PermissionName, result.Error)
			continue
		}

		for _, roleName := range config.Roles {
			var role models.Role
			if err := DB.Where("role_name = ?", roleName).First(&role).Error; err != nil {
				log.Printf("⚠️ Role %s not found for permission %s", roleName, config.PermissionName)
				continue
			}

			rp := models.RolePermission{
				RoleID:       role.ID,
				PermissionID: perm.ID,
			}

			if err := DB.FirstOrCreate(&rp).Error; err != nil {
				log.Printf("⚠️ Failed to map %s to %s: %v", roleName, config.PermissionName, err)
			}
		}
	}
	log.Println("🟢 Permission seeding completed")

	//--------------------------------------------------
	// 3. Seed Admin Users
	//--------------------------------------------------
	log.Println("🟢 Seeding admin users...")

	admins := []models.User{
		{
			FirstName:   "Admin",
			LastName:    "System",
			Email:       "admin@example.com",
			PhoneNumber: "+251900000000",
			Password:    "$2a$10$pkluPLasY7LCXOK25EBkmeUsQDuZwrOhKMhu5EXfN4W0YOZPqST7S",
			IsVerified:  true,
		},
		{
			FirstName:   "Abeselom",
			LastName:    "Solomon",
			Email:       "abeselomsolomon106@example.com",
			PhoneNumber: "+251984852481",
			Password:    "$2a$10$pkluPLasY7LCXOK25EBkmeUsQDuZwrOhKMhu5EXfN4W0YOZPqST7S",
			IsVerified:  true,
		},
	}

	for _, admin := range admins {
		var existingUser models.User
		if err := DB.Where("email = ?", admin.Email).First(&existingUser).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				if err := DB.Create(&admin).Error; err != nil {
					log.Printf("⚠️ Failed to create admin %s: %v", admin.Email, err)
					continue
				}
				log.Printf("✅ Created admin: %s", admin.Email)
			} else {
				log.Printf("⚠️ Error checking admin existence: %v", err)
				continue
			}
		} else {
			admin = existingUser
			log.Printf("ℹ️ Admin already exists: %s", admin.Email)
		}

		var adminRole models.Role
		if err := DB.Where("role_name = ?", "Admin").First(&adminRole).Error; err != nil {
			log.Printf("⚠️ Failed to find Admin role: %v", err)
			continue
		}

		userRole := models.UserRole{
			UserID: admin.ID,
			RoleID: adminRole.ID,
		}
		if err := DB.FirstOrCreate(&userRole).Error; err != nil {
			log.Printf("⚠️ Failed to assign Admin role: %v", err)
		}
	}

	//--------------------------------------------------
	// 4. Seed Vehicle Owners
	//--------------------------------------------------
	log.Println("🟢 Seeding vehicle owners...")

	var owners []models.User
	for i := 0; i < 5000; i++ {
		owner := models.User{
			FirstName:   fmt.Sprintf("Owner%d", i+1),
			LastName:    "Owner",
			Email:       fmt.Sprintf("owner%d@example.com", i+1),
			PhoneNumber: fmt.Sprintf("+2519%08d", rand.Intn(100000000)),
			Password:    "$2a$10$pkluPLasY7LCXOK25EBkmeUsQDuZwrOhKMhu5EXfN4W0YOZPqST7S",
			IsVerified:  true,
		}

		if err := DB.Create(&owner).Error; err != nil {
			log.Printf("⚠️ Failed to create owner %d: %v", i+1, err)
			continue
		}

		var ownerRole models.Role
		if err := DB.Where("role_name = ?", "Owner").First(&ownerRole).Error; err != nil {
			log.Printf("⚠️ Owner role lookup failed: %v", err)
			continue
		}

		ownerUserRole := models.UserRole{
			UserID: owner.ID,
			RoleID: ownerRole.ID,
		}
		if err := DB.FirstOrCreate(&ownerUserRole).Error; err != nil {
			log.Printf("⚠️ Failed to assign Owner role: %v", err)
		}

		owners = append(owners, owner)
	}
	log.Printf("✅ Created %d vehicle owners", len(owners))

	//--------------------------------------------------
	// 5. Seed Drivers with Images
	//--------------------------------------------------
	log.Println("🟢 Seeding drivers...")

	var imageEntries []ImageEntry
	imgFile, err := os.ReadFile(filepath.Join(cwd, "pkg/database/data/images.json"))
	if err != nil {
		log.Fatalf("❌ Failed to read images file: %v", err)
	}

	if err := json.Unmarshal(imgFile, &imageEntries); err != nil {
		log.Fatalf("❌ Failed to parse images JSON: %v", err)
	}

	var driverIDs []uint
	imageIndex := 0
	for i := 0; i < 5000; i++ {
		imgURL := imageEntries[imageIndex%len(imageEntries)].Image
		imageIndex++

		driverUser := models.User{
			FirstName:      fmt.Sprintf("Driver%d", i+1),
			LastName:       "Driver",
			Email:          fmt.Sprintf("driver%d@example.com", i+1),
			PhoneNumber:    fmt.Sprintf("+2519%08d", rand.Intn(100000000)),
			Password:       "$2a$10$pkluPLasY7LCXOK25EBkmeUsQDuZwrOhKMhu5EXfN4W0YOZPqST7S",
			ProfilePicture: &imgURL,
			IsVerified:     true,
		}

		if err := DB.Create(&driverUser).Error; err != nil {
			log.Printf("⚠️ Failed to create driver user %d: %v", i+1, err)
			continue
		}

		var driverRole models.Role
		if err := DB.Where("role_name = ?", "Driver").First(&driverRole).Error; err != nil {
			log.Printf("⚠️ Driver role lookup failed: %v", err)
			continue
		}

		driverUserRole := models.UserRole{
			UserID: driverUser.ID,
			RoleID: driverRole.ID,
		}
		if err := DB.FirstOrCreate(&driverUserRole).Error; err != nil {
			log.Printf("⚠️ Failed to assign Driver role: %v", err)
		}

		driverProfile := models.Driver{
			UserID:             driverUser.ID,
			DriverLicense:      fmt.Sprintf("DL-%09d", rand.Intn(1000000000)),
			Status:             "active",
			DrivingLicensePath: fmt.Sprintf("/docs/drivers/%d/license.pdf", driverUser.ID),
			NationalIdPath:     fmt.Sprintf("/docs/drivers/%d/national-id.pdf", driverUser.ID),
		}

		if err := DB.Create(&driverProfile).Error; err != nil {
			log.Printf("⚠️ Failed to create driver profile: %v", err)
			continue
		}

		driverIDs = append(driverIDs, driverProfile.ID)
	}
	log.Printf("✅ Created %d drivers with profiles", len(driverIDs))

	//--------------------------------------------------
	// 6. Seed Stations
	//--------------------------------------------------
	log.Println("🟢 Seeding stations...")

	var stations []models.Station
	stationFile, err := os.ReadFile(filepath.Join(cwd, "pkg/database/data/station.json"))
	if err != nil {
		log.Fatalf("❌ Failed to read station file: %v", err)
	}

	if err := json.Unmarshal(stationFile, &stations); err != nil {
		log.Fatalf("❌ Failed to parse station JSON: %v", err)
	}

	for _, station := range stations {
		if err := DB.FirstOrCreate(&models.Station{}, station).Error; err != nil {
			log.Printf("⚠️ Failed to seed station %s: %v", station.LocationName, err)
		}
	}
	log.Println("✅ Stations seeded successfully")

	log.Println("🟢 Seeding routes...")

	if err := DB.Order("id").Find(&stations).Error; err != nil {
		log.Fatalf("❌ Failed to fetch stations: %v", err)
	}

	// Pair stations into routes
	for i := 0; i < len(stations); i += 2 {
		if i+1 >= len(stations) {
			log.Println("⚠️ Odd number of stations, skipping the last one")
			break
		}
		stationA := stations[i]
		stationB := stations[i+1]

		// Create route if it doesn't exist
		route := models.Route{
			LocationA: stationA.ID,
			LocationB: stationB.ID,
		}
		result := DB.Where("location_a = ? AND location_b = ?", stationA.ID, stationB.ID).FirstOrCreate(&route)
		if result.Error != nil {
			log.Printf("⚠️ Failed to create route between %s and %s: %v", stationA.LocationName, stationB.LocationName, result.Error)
			continue
		}
		if result.RowsAffected == 0 {
			log.Printf("⏭️ Route between %s and %s already exists, skipping", stationA.LocationName, stationB.LocationName)
		} else {
			log.Printf("✅ Created route %d: %s to %s", route.ID, stationA.LocationName, stationB.LocationName)
		}
	}

	log.Println("✅ Routes seeded successfully")

	log.Println("🟢 Seeding paths...")

	var routes []models.Route
	if err := DB.Preload("StationA").Preload("StationB").Find(&routes).Error; err != nil {
		log.Fatalf("❌ Failed to fetch routes: %v", err)
	}

	const averageSpeed = 50.0 // Assume average speed of 50 km/h

	for _, route := range routes {
		// Check if path already exists for this route
		var existingPath models.Path
		if err := DB.Where("route_id = ?", route.ID).First(&existingPath).Error; err == nil {
			log.Printf("⏭️ Path for route %d already exists, skipping", route.ID)
			continue
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("⚠️ Error checking for existing path for route %d: %v", route.ID, err)
			continue
		}

		stationA := route.StationA
		stationB := route.StationB

		// Calculate distance using Haversine formula
		distance := Haversine(
			stationA.Latitude,
			stationA.Longitude,
			stationB.Latitude,
			stationB.Longitude,
		)

		// Calculate estimated time
		hours := distance / averageSpeed
		estimatedTime := time.Duration(hours * float64(time.Hour))

		// Generate path name
		pathName := fmt.Sprintf("%s -> %s", stationA.LocationName, stationB.LocationName)

		// Create path
		path := models.Path{
			RouteID:       route.ID,
			PathName:      pathName,
			DistanceKM:    distance,
			EstimatedTime: estimatedTime,
			IsActive:      true,
		}

		if err := DB.Create(&path).Error; err != nil {
			log.Printf("⚠️ Failed to create path for route %d: %v", route.ID, err)
		} else {
			log.Printf("✅ Created path for route %d: %s (%.2f km, %s)",
				route.ID, pathName, distance, estimatedTime.String())
		}
	}

	log.Println("✅ Paths seeded successfully")

	//--------------------------------------------------
	// 7. Seed Vehicle Types
	//--------------------------------------------------
	log.Println("🟢 Seeding vehicle types...")

	var vehicleTypeEntries []VehicleTypeEntry
	vtFile, err := os.ReadFile(filepath.Join(cwd, "pkg/database/data/vehicle_types.json"))
	if err != nil {
		log.Fatalf("❌ Failed to read vehicle types file: %v", err)
	}

	if err := json.Unmarshal(vtFile, &vehicleTypeEntries); err != nil {
		log.Fatalf("❌ Failed to parse vehicle types JSON: %v", err)
	}

	var vehicleTypes []models.VehicleType
	for _, entry := range vehicleTypeEntries {
		vt := models.VehicleType{
			TypeName:    entry.TypeName,
			Description: entry.Description,
			Capacity:    entry.Capacity,
		}

		if err := DB.FirstOrCreate(&vt, "type_name = ?", entry.TypeName).Error; err != nil {
			log.Printf("⚠️ Failed to seed vehicle type %s: %v", entry.TypeName, err)
		} else {
			vehicleTypes = append(vehicleTypes, vt)
		}
	}
	log.Println("🟢 Vehicle type seeding completed")

	//--------------------------------------------------
	// 8. Seed Vehicles
	//--------------------------------------------------
	log.Println("🟢 Seeding vehicles...")

	makes := []string{"Toyota", "Honda", "Ford", "BMW", "Mercedes"}
	colors := []string{"Red", "Blue", "Black", "White", "Silver"}
	statuses := []string{"active", "inactive", "maintenance"}

	existingVINs := make(map[string]bool)
	var vehicles []models.Vehicle

	for i := 0; i < 5000; i++ {
		vehicleType := vehicleTypes[rand.Intn(len(vehicleTypes))]
		owner := owners[rand.Intn(len(owners))]

		var vin string
		for {
			vin = generateVIN()
			if !existingVINs[vin] {
				existingVINs[vin] = true
				break
			}
		}

		vehicle := models.Vehicle{
			TypeID:        vehicleType.ID,
			LicenseNumber: fmt.Sprintf("%03d-%04d", rand.Intn(1000), rand.Intn(10000)),
			VIN:           vin,
			Make:          makes[rand.Intn(len(makes))],
			Year:          uint(rand.Intn(23) + 2001),
			Color:         colors[rand.Intn(len(colors))],
			CarPicture:    fmt.Sprintf("https://source.unsplash.com/300x200/?%s", vehicleType.TypeName),
			Bollo:         fmt.Sprintf("/docs/vehicles/%d/bollo.pdf", i),
			Insurance:     fmt.Sprintf("/docs/vehicles/%d/insurance.pdf", i),
			Libre:         fmt.Sprintf("/docs/vehicles/%d/libre.pdf", i),
			OwnerID:       owner.ID,
			Status:        statuses[rand.Intn(len(statuses))],
		}

		if err := DB.Create(&vehicle).Error; err != nil {
			log.Printf("⚠️ Failed to create vehicle %d: %v", i+1, err)
			continue
		}

		vehicles = append(vehicles, vehicle)
	}
	log.Printf("✅ Created %d vehicles with unique VINs", len(vehicles))

	//--------------------------------------------------
	// 9. Assign Drivers to Vehicles
	//--------------------------------------------------
	log.Println("🟢 Assigning drivers to vehicles...")

	vehicleIDs := make([]uint, len(vehicles))
	for i, v := range vehicles {
		vehicleIDs[i] = v.ID
	}

	pairLimit := len(driverIDs)
	if len(vehicleIDs) < pairLimit {
		pairLimit = len(vehicleIDs)
	}

	successfulAssignments := 0
	for i := 0; i < pairLimit; i++ {
		err := DB.Transaction(func(tx *gorm.DB) error {
			// Clear existing driver assignment
			if err := tx.Model(&models.Vehicle{}).
				Where("driver_id = ?", driverIDs[i]).
				Update("driver_id", nil).Error; err != nil {
				return fmt.Errorf("driver cleanup failed: %v", err)
			}

			// Assign to vehicle
			if err := tx.Model(&models.Vehicle{}).
				Where("id = ?", vehicleIDs[i]).
				Update("driver_id", driverIDs[i]).Error; err != nil {
				return fmt.Errorf("assignment failed: %v", err)
			}

			// Create audit log
			logEntry := models.DriverAssignmentLog{
				VehicleID:   vehicleIDs[i],
				DriverID:    driverIDs[i],
				AssignedAt:  time.Now(),
				Action:      "ASSIGNED",
				PerformedBy: "SEED_SYSTEM",
			}

			if err := tx.Create(&logEntry).Error; err != nil {
				return fmt.Errorf("log creation failed: %v", err)
			}

			successfulAssignments++
			return nil
		})

		if err != nil {
			log.Printf("⚠️ Failed to assign driver %d to vehicle %d: %v",
				driverIDs[i], vehicleIDs[i], err)
		}
	}

	// Handle leftovers
	if len(vehicles) > pairLimit {
		log.Printf("ℹ️ %d vehicles remain unassigned", len(vehicles)-pairLimit)
	}
	if len(driverIDs) > pairLimit {
		log.Printf("ℹ️ %d drivers remain unassigned", len(driverIDs)-pairLimit)
	}

	//--------------------------------------------------
	// Final Report
	//--------------------------------------------------
	log.Println("\n📊 Seeding Report:")
	log.Printf("|- Roles: %d", len(roles))
	log.Printf("|- Permissions: %d", len(permissionConfigs)+len(basicPermissions))
	log.Printf("|- Admin Users: %d", 2)
	log.Printf("|- Owners: %d", len(owners))
	log.Printf("|- Drivers: %d", len(driverIDs))
	log.Printf("|- Stations: %d", len(stations))
	log.Printf("|- Vehicle Types: %d", len(vehicleTypes))
	log.Printf("|- Vehicles: %d", len(vehicles))
	log.Printf("└─ Driver Assignments: %d", successfulAssignments)
	log.Println("🎉 Database seeding completed successfully!")
}
func Haversine(lat1, lon1, lat2, lon2 float64) float64 {
	const earthRadius = 6371 // Earth's mean radius in kilometers

	// Convert degrees to radians
	degToRad := func(deg float64) float64 { return deg * math.Pi / 180 }

	lat1Rad := degToRad(lat1)
	lon1Rad := degToRad(lon1)
	lat2Rad := degToRad(lat2)
	lon2Rad := degToRad(lon2)

	// Differences in coordinates
	dLat := lat2Rad - lat1Rad
	dLon := lon2Rad - lon1Rad

	// Haversine formula components
	a := math.Pow(math.Sin(dLat/2), 2) +
		math.Cos(lat1Rad)*math.Cos(lat2Rad)*
			math.Pow(math.Sin(dLon/2), 2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	// Calculate the distance
	distance := earthRadius * c

	return distance
}
