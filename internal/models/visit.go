package models

import "time"

type Visit struct {
	ID        int       `json:"id"`
	PatientID int       `json:"patient_id"`
	DoctorID  int       `json:"doctor_id"`
	VisitDate time.Time `json:"visit_date"`
	Notes     string    `json:"notes"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
