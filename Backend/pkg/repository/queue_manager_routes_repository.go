package repository

import (
	"errors"
	"fmt"
	"kassech/backend/pkg/database"
	models "kassech/backend/pkg/model"

	"gorm.io/gorm"
)

type QueueManagerRouteRepository struct{}

func (r *QueueManagerRouteRepository) Create(route *models.QueueManagerRoute, pathIDs []uint) error {
	fmt.Printf("Creating route: %+v\n", route)
	fmt.Printf("Associating paths: %+v\n", pathIDs)

	if route == nil {
		return errors.New("route is nil")
	}

	tx := database.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Create(route).Error; err != nil {
		tx.Rollback()
		return err
	}

	var paths []models.Path
	if err := tx.Where("id IN ?", pathIDs).Find(&paths).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Model(route).Association("Paths").Append(paths); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
func (r *QueueManagerRouteRepository) GetAll(page, perPage int, search string) ([]models.QueueManagerRoute, int64, error) {
	var routes []models.QueueManagerRoute
	var total int64

	query := database.DB.Preload("Paths").
		Preload("Station").
		Preload("User").
		Model(&models.QueueManagerRoute{})

	if search != "" {
		query = query.Where("name ILIKE ?", "%"+search+"%")
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Offset((page - 1) * perPage).
		Limit(perPage).
		Find(&routes).
		Error

	return routes, total, err
}

func (r *QueueManagerRouteRepository) Delete(id uint) error {
	return database.DB.Transaction(func(tx *gorm.DB) error {
		// Clear associations first
		if err := tx.Model(&models.QueueManagerRoute{Model: gorm.Model{ID: id}}).
			Association("Paths").Clear(); err != nil {
			return err
		}

		return tx.Delete(&models.QueueManagerRoute{}, id).Error
	})
}

func (r *QueueManagerRouteRepository) GetAllForUser(userID uint, page, perPage int, search string) ([]models.QueueManagerRoute, int64, error) {
	var routes []models.QueueManagerRoute
	var total int64

	query := database.DB.Preload("Paths").
		Preload("Station").
		Preload("User").
		Model(&models.QueueManagerRoute{}).
		Where("user_id = ?", userID)

	if search != "" {
		query = query.Where("name ILIKE ?", "%"+search+"%")
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Offset((page - 1) * perPage).
		Limit(perPage).
		Find(&routes).
		Error

	return routes, total, err
}
