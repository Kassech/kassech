package database

import (
	models "kassech/backend/pkg/model"
	"log"
)

// Migrate performs database migrations based on the specified migration type.
func Migrate(migrationType string) {
	if DB == nil {
		log.Fatal("Database connection is nil. Ensure Connect() is called before Migrate().")
	}

	// List of models to migrate
	modelsToMigrate := []interface{}{
		&models.User{},
		&models.Role{},
		&models.Permission{},
		&models.UserRole{},
		&models.Driver{},
		&models.Owner{},
		&models.QueueManager{},
		&models.Station{},
		&models.Route{},
		&models.Path{},
		&models.Vehicle{},
		&models.Passenger{},
		&models.TravelLog{},
		&models.VehicleGPSLog{},
		&models.Attachment{},
		&models.NotificationLog{},
		&models.QueueManagerRoute{},
		&models.PassengerHistory{},
		&models.DriverVehicle{},
		&models.ManualDriverAssignmentHistory{},
		&models.AutoDriverAssignmentHistory{},
		&models.VehicleType{},
		&models.NotificationToken{},
	}

	switch migrationType {
	case "clean":
		// Drop all tables before running AutoMigrate
		log.Println("Performing clean migration: dropping existing tables...")
		for _, model := range modelsToMigrate {
			if err := DB.Migrator().DropTable(model); err != nil {
				log.Printf("Failed to drop table for model %T: %v", model, err)
			} else {
				log.Printf("Dropped table for model %T successfully.", model)
			}
		}
		// Run AutoMigrate after dropping tables
		fallthrough // Proceed to auto migrate after dropping tables

	case "auto":
		log.Println("Performing auto migration: applying changes...")
		if err := DB.AutoMigrate(modelsToMigrate...); err != nil {
			log.Fatalf("Failed to run migrations: %v", err)
		}
		log.Println("Migrations completed successfully!")
	default:
		log.Printf("Unknown migration type: %s. Use 'auto' or 'clean'.", migrationType)
	}
}
