package handlers

import "github.com/platinumpizza29/medicare/internal/models"

// AuthToken is the standard JWT response payload.
type AuthToken struct {
	Token string `json:"token"`
}

// Message is a simple message response.
type Message struct {
	Message string `json:"message"`
}

// DoctorRegistration describes the doctor registration payload.
type DoctorRegistration struct {
	FirstName     string `json:"FirstName"`
	LastName      string `json:"LastName"`
	ClinicName    string `json:"ClinicName"`
	Address       string `json:"Address"`
	LicenseNumber string `json:"LicenseNumber"`
	MobileNumber  string `json:"MobileNumber"`
	AadharNumber  string `json:"AadharNumber"`
	Specialty     string `json:"Specialty"`
	Experience    string `json:"Experience"`
	DOB           string `json:"DOB"`
	Gender        string `json:"Gender"`
	BloodGroup    string `json:"BloodGroup"`
	Email         string `json:"Email"`
	Password      string `json:"Password"`
}

// DoctorLogin describes the doctor login payload.
type DoctorLogin struct {
	Email    string `json:"Email"`
	Password string `json:"Password"`
}

// DoctorLoginResponse describes the doctor login response payload.
type DoctorLoginResponse struct {
	Token          string                 `json:"token"`
	RecentPatients []models.RecentPatient `json:"recent_patients"`
}

// PatientRegistration describes the patient registration payload.
type PatientRegistration struct {
	FirstName        string `json:"FirstName"`
	LastName         string `json:"LastName"`
	Address          string `json:"Address"`
	MobileNumber     string `json:"MobileNumber"`
	AadharNumber     string `json:"AadharNumber"`
	DOB              string `json:"DOB"`
	Gender           string `json:"Gender"`
	BloodGroup       string `json:"BloodGroup"`
	EmergencyContact string `json:"EmergencyContact"`
	Email            string `json:"Email"`
	Password         string `json:"Password"`
}

// PatientLogin describes the patient login payload.
type PatientLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// PrescriptionInput describes create prescription payload.
type PrescriptionInput struct {
	DoctorID     int    `json:"doctorId"`
	PatientID    int    `json:"patientId"`
	Diagnosis    string `json:"diagnosis"`
	Medications  string `json:"medications"`
	Instructions string `json:"instructions"`
	FollowUpDate string `json:"followUpDate"`
}

// PrescriptionCreated describes the create prescription response.
type PrescriptionCreated struct {
	Message      string              `json:"message"`
	Prescription models.Prescription `json:"prescription"`
}

// VisitInput describes create/update visit payload.
type VisitInput struct {
	PatientID int    `json:"patient_id"`
	DoctorID  int    `json:"doctor_id"`
	VisitDate string `json:"visit_date"`
	Notes     string `json:"notes"`
}
