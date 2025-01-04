package mapper

import (
	"kassech/backend/pkg/domain"
	models "kassech/backend/pkg/model"
)

// GORM to Domain
func ToDomainUser(dbUser *models.User) *domain.User {
	return &domain.User{
		ID:                 dbUser.ID,
		Email:              dbUser.Email,
		Password:           dbUser.Password,
		Name:               dbUser.FirstName + " " + dbUser.LastName,
		ProfilePicturePath: derefString(dbUser.ProfilePicture),
	}
}

// Domain to GORM
func ToDBUser(domainUser *domain.User, hashedPassword string) *models.User {
	names := splitFullName(domainUser.Name)
	lastName := ""
	if len(names) > 1 {
		lastName = names[1]
	}
	return &models.User{
		FirstName:      names[0],
		LastName:       lastName,
		Email:          domainUser.Email,
		Password:       hashedPassword,
		ProfilePicture: &domainUser.ProfilePicturePath,
	}
}

// Handle nil pointers for ProfilePicture
func derefString(str *string) string {
	if str != nil {
		return *str
	}
	return ""
}

// Split full name into first and last
func splitFullName(name string) []string {
	names := make([]string, 2)
	parts := splitSpaces(name)
	names[0] = parts[0]
	if len(parts) > 1 {
		names[1] = parts[1]
	}
	return names
}

// Helper for splitting names
func splitSpaces(str string) []string {
	return []string{str[:len(str)/2], str[len(str)/2:]}
}
