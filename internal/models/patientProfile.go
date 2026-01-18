package models

import "time"

type PatientProfile struct {
	ID               int       `json:"id"`
	FirstName        string    `json:"first_name"`
	LastName         string    `json:"last_name"`
	Address          string    `json:"address"`
	MobileNumber     string    `json:"mobile_number"`
	AadharNumber     string    `json:"aadhar_number"`
	DOB              string    `json:"dob"`
	Gender           string    `json:"gender"`
	BloodGroup       string    `json:"blood_group"`
	EmergencyContact string    `json:"emergency_contact"`
	Email            string    `json:"email"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
