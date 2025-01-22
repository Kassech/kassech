package domain

import (
	"mime/multipart"
)

// VehicleFormData represents the structure to handle JSON and form data for a vehicle
type VehicleFormData struct {
	CarType        string                `form:"carType"`
	VIN            string                `form:"vin"`
	Make           string                `form:"make"`
	Year           string                `form:"year"`
	Color          string                `form:"color"`
	CarPicture     *multipart.FileHeader `form:"carPicture" json:"-"`
	Bollo          *multipart.FileHeader `form:"bollo" json:"-"`
	Insurance      *multipart.FileHeader `form:"insurance" json:"-"`
	Libre          *multipart.FileHeader `form:"libre" json:"-"`
	CarPicturePath *string
	BolloPath      *string
	InsurancePath  *string
	LibrePath      *string
	OwnerID        uint `form:"ownerID"`
}
