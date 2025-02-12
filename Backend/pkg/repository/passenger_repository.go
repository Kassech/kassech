package repository

import (
	"fmt"
	"kassech/backend/pkg/database"
	models "kassech/backend/pkg/model"
	"log"
	"time"
)

type PassengerRepository struct{}

func (pr *PassengerRepository) Create(passenger *models.Passenger) (*models.Passenger, error) {
	return passenger, database.DB.Create(passenger).Error
}
func (pr *PassengerRepository) GetPassengerCount(pathID uint) (int64, error) {
	var passengers []models.Passenger
	var count int64
	query := database.DB.Model(&models.Passenger{}).Where("path_id = ?", pathID)
	if err := query.Find(&passengers).Error; err != nil {
		return 0, err
	}
	if err := query.Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
func (pr *PassengerRepository) FindByID(passengerID uint) (*models.Passenger, error) {
	var passenger models.Passenger
	if err := database.DB.First(&passenger, passengerID).Error; err != nil {
		return nil, err
	}
	return &passenger, nil
}

func (pr *PassengerRepository) GetAll(page, perPage int, status string, pathID uint) ([]models.Passenger, int64, error) {
	var passengers []models.Passenger
	var total int64

	query := database.DB.Model(&models.Passenger{})

	if status != "" {
		query = query.Where("status = ?", status)
	}

	if pathID != 0 {
		query = query.Where("path_id = ?", pathID)
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Offset((page - 1) * perPage).Limit(perPage).Find(&passengers).Error
	if err != nil {
		return nil, 0, err
	}

	return passengers, total, nil
}

func (pr *PassengerRepository) DeleteByID(passengerID uint, forceDelete bool) (*models.Passenger, error) {
	var passenger models.Passenger
	if err := database.DB.First(&passenger, passengerID).Error; err != nil {
		return nil, err
	}
	if forceDelete {
		if err := database.DB.Unscoped().Delete(&passenger).Error; err != nil {
			return nil, err
		}
	} else {
		if err := database.DB.Delete(&passenger).Error; err != nil {
			return nil, err
		}
	}
	return &passenger, nil
}
func (pr *PassengerRepository) IncrementPassengerCountBy(pathID uint, amount int) error {
	for i := 0; i < amount; i++ {
		passenger := models.Passenger{
			PathID: pathID,
			// Add any other fields you need for the passenger row
		}
		if err := database.DB.Create(&passenger).Error; err != nil {
			return err
		}
	}
	return nil
}

func (pr *PassengerRepository) DecrementPassengerCountBy(pathID uint, amount int) error {
	// Fetch the oldest `amount` rows for the given pathID
	var passengers []models.Passenger
	if err := database.DB.Where("path_id = ?", pathID).Order("created_at ASC").Limit(amount).Find(&passengers).Error; err != nil {
		return err
	}

	// Unscoped delete for the fetched rows
	for _, passenger := range passengers {
		if err := database.DB.Unscoped().Delete(&passenger).Error; err != nil {
			return err
		}
	}
	return nil
}

func (pr *PassengerRepository) AssignToCarAndRemove(passengers []models.Passenger, vehicle models.Vehicle, pathID uint) error {
	if len(passengers) == 0 {
		log.Println("No passengers to assign to car")
		return nil
	}

	// Begin a database transaction
	tx := database.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	// Insert into passenger_histories table
	histories := make([]models.PassengerHistory, len(passengers))
	for i, passenger := range passengers {
		histories[i] = models.PassengerHistory{
			PassengerID: passenger.ID,
			VehicleID:   vehicle.ID,
			RouteID:     pathID,
			TravelDate:  time.Now().Format("2006-01-02"),
		}
	}

	if err := tx.Create(&histories).Error; err != nil {
		tx.Rollback()
		log.Printf("Failed to insert into passenger_histories: %v", err)
		return err
	}

	// Delete from passengers table
	for _, passenger := range passengers {
		if err := tx.Unscoped().Delete(&passenger).Error; err != nil {
			tx.Rollback()
			log.Printf("Failed to delete passenger %d: %v", passenger.ID, err)
			return err
		}
	}

	// Log the assignment in the driver_assignment_history table
	assignmentHistory := models.DriverAssignmentHistory{
		DriverID:        *vehicle.DriverID,
		VehicleID:       vehicle.ID,
		AssignedByID:    nil,
		AssignmentDate:  time.Now().Format("2006-01-02"),
		AssignmentNotes: "",
		AlgorithmUsed:   "auto",
		AssignmentType:  "auto",
	}
	if err := tx.Create(&assignmentHistory).Error; err != nil {
		tx.Rollback()
		log.Printf("Failed to insert into driver_assignment_history: %v", err)
		return err
	}

	// Mark the vehicle as assigned
	if err := vehicle.SetStatus("assigned"); err != nil {
		tx.Rollback()
		return err // Return validation error (e.g., invalid status)
	}

	// Save the updated status to the database
	if err := tx.Save(&vehicle).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to update status: %v", err)
	}
	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		log.Printf("Failed to commit transaction: %v", err)
		return err
	}

	log.Printf("Successfully assigned %d passengers to car %d and moved to history", len(passengers), vehicle.ID)
	return nil
}
