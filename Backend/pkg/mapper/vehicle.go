package mapper

import (
	"fmt"
	"kassech/backend/pkg/domain"
	models "kassech/backend/pkg/model"
	"strconv"
)

// GORM to Domain
func ToDomainVehicle(dbVehicle *models.Vehicle) *domain.VehicleFormData {
	return &domain.VehicleFormData{
		CarType: strconv.Itoa(int(dbVehicle.TypeID)), // Convert uint to string
		VIN:     dbVehicle.VIN,
		Make:    dbVehicle.Make,
		Year:    strconv.Itoa(int(dbVehicle.Year)), // Convert uint to string
		Color:   dbVehicle.Color,
		OwnerID: dbVehicle.OwnerID,
		// File fields can be handled here if needed
		CarPicture: nil,
		Bollo:      nil,
		Insurance:  nil,
		Libre:      nil,
	}
}

// Domain to GORM
func ToGormVehicle(domainVehicle *domain.VehicleFormData) (*models.Vehicle, error) {
	// Convert string to uint for Type with validation
	vehicleType, err := strconv.Atoi(domainVehicle.CarType)
	if err != nil {
		return nil, fmt.Errorf("invalid CarType value: %s", domainVehicle.CarType)
	}

	// Convert string to uint for Year with validation
	vehicleYear, err := strconv.Atoi(domainVehicle.Year)
	if err != nil {
		return nil, fmt.Errorf("invalid Year value: %s", domainVehicle.Year)
	}

	return &models.Vehicle{
		TypeID:        uint(vehicleType),
		LicenseNumber: domainVehicle.VIN, // Map VIN to LicenseNumber if intended
		VIN:           domainVehicle.VIN,
		Make:          domainVehicle.Make,
		Year:          uint(vehicleYear),
		Color:         domainVehicle.Color,
		OwnerID:       domainVehicle.OwnerID,
		CarPicture: func() string {
			if domainVehicle.CarPicturePath != nil {
				return *domainVehicle.CarPicturePath
			}
			return ""
		}(),
		Bollo: func() string {
			if domainVehicle.BolloPath != nil {
				return *domainVehicle.BolloPath
			}
			return ""
		}(),
		Insurance: func() string {
			if domainVehicle.InsurancePath != nil {
				return *domainVehicle.InsurancePath
			}
			return ""
		}(),
		Libre: func() string {
			if domainVehicle.LibrePath != nil {
				return *domainVehicle.LibrePath
			}
			return ""
		}(),
	}, nil
}
