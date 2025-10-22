package models

import "time"

type Prescription struct {
	ID           int        `json:"id"`
	DoctorID     int        `json:"doctorId"`
	PatientID    int        `json:"patientId"`
	Diagnosis    string     `json:"diagnosis"`
	Medications  string     `json:"medications"`
	Instructions string     `json:"instructions"`
	FollowUpDate *time.Time `json:"followUpDate,omitempty"`
	CreatedAt    time.Time  `json:"createdAt"`
	UpdatedAt    time.Time  `json:"updatedAt"`
}
