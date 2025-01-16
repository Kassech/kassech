package domain

import (
	"mime/multipart"
)

type User struct {
	ID                    uint   `json:"id"`
	Email                 string `json:"email" binding:"required,email"`
	Password              string `json:"password" ` // Hashed password
	FirstName             string `json:"first_name" validate:"required"`
	LastName              string `json:"last_name" validate:"required"`
	ProfilePicture        *string
	ProfilePictureFile    *multipart.FileHeader `form:"profile"`
	PhoneNumber           string                `json:"phone_number"`
	IsVerified            bool                  `json:"is_verified"`
	Role                  uint                  `json:"role"`
	DrivingLicenseFile    *multipart.FileHeader `form:"driving_license"`
	NationalIdFile        *multipart.FileHeader `form:"national_id"`
	InsuranceDocumentFile *multipart.FileHeader `form:"insurance_document"`
	OtherFile             *multipart.FileHeader `form:"other_file"`
	DrivingLicense        *string               `json:"driving_license"`
	NationalId            *string               `json:"national_id"`
	InsuranceDocument     *string               `json:"insurance_document"`
	OtherDocument         *string               `json:"other_document"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
