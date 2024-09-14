package model

import (
	"fmt"
	"spl-users/ent"
	"time"
)

type User struct {
	Run             string    `json:"run"`
	FirstName       string    `json:"firstName"`
	LastName        string    `json:"lastName"`
	Email           string    `json:"email"`
	PhoneNumber     string    `json:"phoneNumber"`
	Gender          string    `json:"gender"`
	HomeAddress     string    `json:"homeAddress"`
	City            string    `json:"city"`
	BirthDate       time.Time `json:"birthDate"`
	ExpirationDate  time.Time `json:"expirationDate"`
	PlanType        string    `json:"planType"`
	EmergencyName   string    `json:"emergencyName"`
	EmergencyNumber string    `json:"emergencyNumber"`
	MaritalStatus   string    `json:"maritalStatus"`
	Locations       []string  `json:"locations"`
}

func EntUserToUser(user *ent.User) *User {
	if user == nil {
		return nil
	}
	var locations []string
	for _, location := range user.Edges.Locations {
		locations = append(locations, location.Value)
	}

	return &User{
		Run:             fmt.Sprintf("%d-%s", user.Run, user.VerificationDigit),
		FirstName:       user.FirstName,
		LastName:        user.LastName,
		Email:           user.Email,
		PhoneNumber:     user.PhoneNumber,
		Gender:          string(user.Gender),
		MaritalStatus:   user.MaritalStatus,
		HomeAddress:     user.HomeAddress,
		City:            user.City,
		BirthDate:       user.BirthDate,
		ExpirationDate:  user.ExpirationDate,
		PlanType:        user.PlantType,
		EmergencyName:   user.EmergencyName,
		EmergencyNumber: user.EmergencyNumber,
		Locations:       locations,
	}
}
