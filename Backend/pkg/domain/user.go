package domain

import (
	"mime/multipart"
)

type User struct {
	ID                 uint                  `json:"id" form:"id"`
	Email              string                `json:"email" form:"email" binding:"required,email"`
	Password           string                `json:"password" form:"password" binding:"required"` // Hashed password
	FirstName          string                `json:"first_name" form:"first_name" validate:"required"`
	LastName           string                `json:"last_name" form:"last_name" validate:"required"`
	ProfilePicture     *string               `json:"profile_picture" form:"profile_picture"`
	ProfilePictureFile *multipart.FileHeader `form:"profile"`
	PhoneNumber        string                `json:"phone_number" form:"phone_number"`
	IsVerified         bool                  `json:"is_verified" form:"is_verified"`
	Role               uint                  `json:"role" form:"role"`
	Permissions        []string              `json:"permissions,omitempty" form:"permissions,omitempty"`
	Roles              []string              `json:"roles,omitempty" form:"roles,omitempty"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
