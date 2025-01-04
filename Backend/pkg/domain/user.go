package domain

import (
	"mime/multipart"
)

type User struct {
	ID                 uint                  `json:"id"`
	Email              string                `json:"email" binding:"required,email"`
	Password           string                `json:"password" binding:"required"` // Hashed password
	Name               string                `json:"name" binding:"required"`
	ProfilePicture     *multipart.FileHeader `json:"profile_picture" form:"profile_picture"`
	ProfilePicturePath string                ``
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
