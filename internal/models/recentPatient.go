package models

import "time"

type RecentPatient struct {
	ID            int       `json:"id"`
	FirstName     string    `json:"first_name"`
	LastName      string    `json:"last_name"`
	MobileNumber  string    `json:"mobile_number"`
	Email         string    `json:"email"`
	LastVisitDate time.Time `json:"last_visit_date"`
}
