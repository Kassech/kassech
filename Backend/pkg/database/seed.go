package database

import (
	models "kassech/backend/pkg/model"
	"log"

	"gorm.io/gorm"
)

func SeedDB() {
	// Seed roles
	roles := []models.Role{
		{RoleName: "Admin", Description: "Administrator with full access"},
		{RoleName: "Driver", Description: "Regular user with limited access"},
		{RoleName: "QueueManager", Description: "Manager with elevated permissions"},
		{RoleName: "CS", Description: "Manager with elevated permissions"},
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
		IsVerified:  false,
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

	log.Println("Database seeding completed.")
}
