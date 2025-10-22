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
	DOB           time.Time
	Gender        string
	BloodGroup    string
	Email         string
	Password      string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type DoctorRequest struct {
	Email    string
	Password string
}
