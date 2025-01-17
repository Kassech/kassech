package domain

import (
	"mime/multipart"
)

type Driver struct {
	User
	DrivingLicenseFile    *multipart.FileHeader `form:"driving_license"`
	NationalIdFile        *multipart.FileHeader `form:"national_id"`
	InsuranceDocumentFile *multipart.FileHeader `form:"insurance_document"`
	OtherFile             *multipart.FileHeader `form:"other_file"`
	DrivingLicense        *string               `json:"driving_license"`
	NationalId            *string               `json:"national_id"`
	InsuranceDocument     *string               `json:"insurance_document"`
	OtherDocument         *string               `json:"other_document"`
}
