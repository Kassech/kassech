package repository

import (
	"kassech/backend/pkg/database"
	models "kassech/backend/pkg/model"
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
