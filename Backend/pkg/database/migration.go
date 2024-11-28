package database

import (
	models "kassech/backend/pkg/model"
	"log"
)

func Migrate() {
	if DB == nil {
		log.Fatal("Database connection is nil. Ensure Connect() is called before Migrate().")
	}

	err := DB.AutoMigrate(
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
	)

	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	log.Println("Migrations completed!")
}
