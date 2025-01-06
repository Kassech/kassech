package domain

import (
	"mime/multipart"
)

type User struct {
	ID                 uint   `json:"id"`
	Email              string `json:"email" binding:"required,email"`
	Password           string `json:"password" binding:"required"` // Hashed password
	FirstName          string `json:"first_name" validate:"required"`
	LastName           string `json:"last_name" validate:"required"`
	ProfilePicture     *string
	ProfilePictureFile *multipart.FileHeader `form:"profile"`
	PhoneNumber        string                `json:"phone_number"`
	IsVerified         bool                  `json:"is_verified"`
	Role               uint                  `json:"role"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
