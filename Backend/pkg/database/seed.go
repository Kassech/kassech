package database

import (
	models "kassech/backend/pkg/model"
	"log"
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

	// Seed role-permission mappings (Admin has all permissions, User has limited permissions)
	rolePermissions := []models.RolePermission{
		{RoleID: 1, PermissionID: 1}, // Admin can CreateUser
		{RoleID: 1, PermissionID: 2}, // Admin can DeleteUser
		{RoleID: 1, PermissionID: 3}, // Admin can UpdateUser
		{RoleID: 1, PermissionID: 4}, // Admin can ViewReports
		{RoleID: 2, PermissionID: 4}, // User can ViewReports
	}

	for _, rolePermission := range rolePermissions {
		if err := DB.FirstOrCreate(&models.RolePermission{}, rolePermission).Error; err != nil {
			log.Printf("Failed to seed role-permission mapping: %v\n", err)
		}
	}

	log.Println("Database seeding completed.")
}
