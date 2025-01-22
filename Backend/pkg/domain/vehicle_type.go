package domain

// VehicleTypeFormData represents the structure to handle JSON and form data for a vehicle type
type VehicleTypeFormData struct {
	TypeName    string `json:"type_name"`
	Capacity    uint   `json:"capacity"`
	Description string `json:"description"`
}
