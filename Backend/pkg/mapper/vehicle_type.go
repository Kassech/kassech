package mapper

import (
	"kassech/backend/pkg/domain"
	models "kassech/backend/pkg/model"
)

// GORM to Domain
func ToDomainVehicleType(dbVehicleType *models.VehicleType) *domain.VehicleTypeFormData {
	return &domain.VehicleTypeFormData{
		TypeName:    dbVehicleType.TypeName,
		Capacity:    dbVehicleType.Capacity,
		Description: dbVehicleType.Description,
	}
}

// Domain to GORM
func ToGormVehicleType(domainVehicleType *domain.VehicleTypeFormData) (*models.VehicleType, error) {
	return &models.VehicleType{
		TypeName:    domainVehicleType.TypeName,
		Capacity:    domainVehicleType.Capacity,
		Description: domainVehicleType.Description,
	}, nil
}
