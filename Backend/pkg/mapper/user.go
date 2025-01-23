package mapper

import (
	"kassech/backend/pkg/domain"
	models "kassech/backend/pkg/model"
)

// GORM to Domain
func ToDomainUser(dbUser *models.User) *domain.User {
	return &domain.User{
		ID:             dbUser.ID,
		Email:          dbUser.Email,
		Password:       dbUser.Password,
		FirstName:      dbUser.FirstName,
		LastName:       dbUser.LastName,
		ProfilePicture: dbUser.ProfilePicture,
		IsVerified:     dbUser.IsVerified,
	}
}
func ToGormUser(domainUser *domain.User) *models.User {
	password := domainUser.Password
	if password == "" {
		password = domainUser.PhoneNumber
	}
	return &models.User{
		LastName:       domainUser.LastName,
		FirstName:      domainUser.FirstName,
		Email:          domainUser.Email,
		Password:       password,
		ProfilePicture: domainUser.ProfilePicture,
		PhoneNumber:    domainUser.PhoneNumber,
		IsVerified:     domainUser.IsVerified,
	}
}
