package domain

import (
	"mime/multipart"
)

type User struct {
	ID                 uint                  `json:"id" form:"ID"`
	Email              string                `json:"email" form:"Email" binding:"required,email"`
	Password           string                `json:"password" form:"Password"` // Hashed password
	FirstName          string                `json:"first_name" form:"FirstName" binding:"required"`
	LastName           string                `json:"last_name" form:"LastName" binding:"required"`
	ProfilePicture     *string               `json:"profile_picture" form:"ProfilePicture"`
	ProfilePictureFile *multipart.FileHeader `form:"Profile"`
	PhoneNumber        string                `json:"phone_number" form:"PhoneNumber"`
	IsVerified         bool                  `json:"is_verified" form:"IsVerified"`
	Role               uint                  `json:"role" form:"Role"`
	Permissions        []string              `json:"permissions,omitempty" form:"Permissions,omitempty"`
	Roles              []string              `json:"roles,omitempty" form:"Roles,omitempty"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
