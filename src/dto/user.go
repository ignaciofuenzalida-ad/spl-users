package dto

import "time"

type UpdateUserDto struct {
	FetchStatus     string    `json:"fetchStatus" validate:"required,oneof=COMPLETED ERROR"`
	Status          string    `json:"status" validate:"omitempty,oneof=NOT_FOUND FOUND"`
	FirstName       string    `json:"firstName" validate:"omitempty"`
	LastName        string    `json:"lastName" validate:"omitempty"`
	PhoneNumber     string    `json:"phoneNumber" validate:"omitempty"`
	Gender          string    `json:"gender" validate:"omitempty,oneof=MALE FEMALE"`
	Email           string    `json:"email" validate:"omitempty"`
	MaritalStatus   string    `json:"maritalStatus" validate:"omitempty"`
	HomeAddress     string    `json:"homeAddress" validate:"omitempty"`
	City            string    `json:"city" validate:"omitempty"`
	BirthDate       time.Time `json:"birthDate" validate:"omitempty"`
	ExpirationDate  time.Time `json:"expirationDate" validate:"omitempty"`
	PlantType       string    `json:"plantType" validate:"omitempty"`
	EmergencyName   string    `json:"emergencyName" validate:"omitempty"`
	EmergencyNumber string    `json:"emergencyNumber" validate:"omitempty"`
	Locations       []string  `json:"locations" validate:"omitempty,dive"`
}
