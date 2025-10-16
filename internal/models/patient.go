package models

import "time"

type Patient struct {
	ID               int
	FirstName        string
	LastName         string
	Address          string
	MobileNumber     string
	AadharNumber     string
	DOB              string
	Gender           string
	BloodGroup       string
	EmergencyContact string
	Email            string
	Password         string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
