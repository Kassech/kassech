package database

import (
	"encoding/json"
	"fmt"
	models "kassech/backend/pkg/model"
	"log"
	"os"

	"golang.org/x/exp/rand"
	"gorm.io/gorm"
)

func SeedDB() {
	// Seed roles
	roles := []models.Role{
		{RoleName: "Admin", Description: "Administrator with full access"},
		{RoleName: "Driver", Description: "Regular user with limited access"},
		{RoleName: "QueueManager", Description: "Manager with elevated permissions"},
		{RoleName: "Owner", Description: "Manager with elevated permissions"},
		{RoleName: "CustomerService", Description: "Manager with elevated permissions"},
	}

	for _, role := range roles {
		if err := DB.FirstOrCreate(&models.Role{}, role).Error; err != nil {
			log.Printf("Failed to seed role %s: %v\n", role.RoleName, err)
		}
	}

	// Seed permissions
	permissions := []models.Permission{
		{PermissionName: "CreateUser", Description: "Permission to create users"},
		{PermissionName: "DeleteUser", Description: "Permission to delete users"},
		{PermissionName: "UpdateUser", Description: "Permission to update users"},
		{PermissionName: "ViewReports", Description: "Permission to view reports"},
	}

	for _, permission := range permissions {
		if err := DB.FirstOrCreate(&models.Permission{}, permission).Error; err != nil {
			log.Printf("Failed to seed permission %s: %v\n", permission.PermissionName, err)
		}
	}

	// Seed role-permission mappings
	rolePermissions := []models.RolePermission{
		{RoleID: 1, PermissionID: 1},
		{RoleID: 1, PermissionID: 2},
		{RoleID: 1, PermissionID: 3},
		{RoleID: 1, PermissionID: 4},
		{RoleID: 2, PermissionID: 4},
	}

	for _, rolePermission := range rolePermissions {
		if err := DB.FirstOrCreate(&models.RolePermission{}, rolePermission).Error; err != nil {
			log.Printf("Failed to seed role-permission mapping: %v\n", err)
		}
	}

	// Seed user
	user := models.User{
		FirstName:   "Abeselom",
		LastName:    "Solomon",
		Email:       "abeselomsolomon106@example.com",
		PhoneNumber: "+251984852481",
		IsOnline:    false,
		Password:    "$2a$10$pkluPLasY7LCXOK25EBkmeUsQDuZwrOhKMhu5EXfN4W0YOZPqST7S", // hashed password
		IsVerified:  true,
	}

	// Check if user exists, if not, create it
	var existingUser models.User
	if err := DB.Where("email = ?", user.Email).First(&existingUser).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// Create the new user
			if err := DB.Create(&user).Error; err != nil {
				log.Printf("Failed to seed user: %v\n", err)
			} else {
				log.Println("User seeded successfully.")
			}
		} else {
			log.Printf("Error checking for existing user: %v\n", err)
		}
	} else {
		log.Println("User already exists.")
		user = existingUser // Use the existing user from the database
	}

	// Set the user's role as 'Admin' (super admin)
	var adminRole models.Role
	if err := DB.Where("role_name = ?", "Admin").First(&adminRole).Error; err != nil {
		log.Printf("Failed to find Admin role: %v\n", err)
	} else {
		// Create a UserRole association
		userRole := models.UserRole{
			UserID: user.ID,
			RoleID: adminRole.ID,
		}
		if err := DB.FirstOrCreate(&userRole).Error; err != nil {
			log.Printf("Failed to assign Admin role to user: %v\n", err)
		} else {
			log.Println("User assigned the Admin role successfully.")
		}
	}

	var owners []models.User
	for i := 0; i < 20; i++ {
		user := models.User{
			FirstName:   fmt.Sprintf("User%d", i+1),
			LastName:    "Test",
			Email:       fmt.Sprintf("user%d@example.com", i+1),
			PhoneNumber: fmt.Sprintf("+2519%08d", rand.Intn(100000000)),
			Password:    "password123",
			Roles:       "owner",
		}
		DB.Create(&user)
		owners = append(owners, user)
	}
	log.Printf("%d owners created and seeded into the database.\n", len(owners))

	// Seed stations
	var stations []models.Station

	// Print the current working directory to debug file path issues
	cwd, err := os.Getwd()
	if err != nil {
		log.Printf("Failed to get current working directory: %v\n", err)
	} else {
		log.Printf("Current working directory: %s\n", cwd)
	}

	// Attempt to read the station JSON file
	data, err := os.ReadFile(cwd + "/pkg/database/data/station.json")
	if err != nil {
		log.Printf("Failed to read station JSON file: %v\n", err)
		return
	}

	// Parse the JSON data into the stations slice
	if err := json.Unmarshal(data, &stations); err != nil {
		log.Printf("Failed to parse station JSON file: %v\n", err)
		return
	}

	// Seed each station into the database
	for _, station := range stations {
		if err := DB.FirstOrCreate(&models.Station{}, station).Error; err != nil {
			log.Printf("Failed to seed station %s: %v\n", station.LocationName, err)
		}
	}

	log.Println("Stations seeded successfully.")

	// Seed vehicle types
	var VehicleTypes []models.VehicleType

	// Attempt to read the vehicle type JSON file
	vehicleTypeData, err := os.ReadFile(cwd + "/pkg/database/data/vehicle_type.json")
	if err != nil {
		log.Printf("Failed to read vehicle type JSON file: %v\n", err)
		return
	}

	// Parse the JSON data into the VehicleTypes slice
	if err := json.Unmarshal(vehicleTypeData, &VehicleTypes); err != nil {
		log.Printf("Failed to parse vehicle type JSON file: %v\n", err)
		return
	}

	// Seed each vehicle type into the database
	for _, VehicleType := range VehicleTypes {
		if err := DB.FirstOrCreate(&models.VehicleType{}, VehicleType).Error; err != nil {
			log.Printf("Failed to seed vehicle type %s: %v\n", VehicleType.TypeName, err)
		} else {
			log.Printf("Vehicle type %s seeded successfully.\n", VehicleType.TypeName)
		}
	}

	log.Println("Vehicle types seeded successfully.")

	// seed vehicle data in to the database

	// Fetch owners
	DB.Find(&owners)

	// Create vehicles
	makes := []string{"Toyota", "Honda", "Ford", "BMW", "Mercedes"}
	colors := []string{"Red", "Blue", "Black", "White", "Silver"}
	statuses := []string{"active", "inactive", "maintenance"}

	for i := 0; i < 50; i++ {
		owner := owners[rand.Intn(len(owners))]

		vehicle := models.Vehicle{
			TypeID:        uint(rand.Intn(2) + 1),
			LicenseNumber: fmt.Sprintf("%03d-%04d", rand.Intn(1000), rand.Intn(10000)),
			VIN:           fmt.Sprintf("%017d", rand.Intn(99999999999999999)),
			Make:          makes[rand.Intn(len(makes))],
			Year:          uint(rand.Intn(23) + 2001),
			Color:         colors[rand.Intn(len(colors))],
			CarPicture:    fmt.Sprintf("https://source.unsplash.com/300x200/?car&random=%d", i),
			Bollo:         fmt.Sprintf("https://files.example.com/bollo/%d.pdf", i),
			Insurance:     fmt.Sprintf("https://files.example.com/insurance/%d.pdf", i),
			Libre:         fmt.Sprintf("https://files.example.com/libre/%d.pdf", i),
			OwnerID:       owner.ID,
			Status:        statuses[rand.Intn(len(statuses))],
		}
		DB.Create(&vehicle)
	}

	log.Println("Vehicles seeded successfully.")

	log.Println("Database seeding completed.")
}
