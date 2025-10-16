package models

import "time"

type Doctor struct {
	ID            int
	FirstName     string
	LastName      string
	ClinicName    string
	Address       string
	LicenseNumber string
	MobileNumber  string
	AadharNumber  string
	Specialty     string
	Experience    string
	DOB           string
	Gender        string
	BloodGroup    string
	Email         string
	Password      string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
