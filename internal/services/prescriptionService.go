package services

import (
	"context"

	"github.com/platinumpizza29/medicare/internal/db"
	"github.com/platinumpizza29/medicare/internal/models"
)

type PrescriptionService struct {
	db *db.PrescriptionDB
}

// NewPrescriptionService creates a new service instance
func NewPrescriptionService(db *db.PrescriptionDB) *PrescriptionService {
	return &PrescriptionService{db: db}
}

// Create creates a new prescription
func (s *PrescriptionService) Create(ctx context.Context, p *models.Prescription) error {
	return s.db.CreatePrescription(ctx, p)
}

// GetByID returns a single prescription by ID
func (s *PrescriptionService) GetByID(ctx context.Context, id int) (*models.Prescription, error) {
	return s.db.GetPrescriptionByID(ctx, id)
}

// ListByPatient lists all prescriptions for a given patient ID
func (s *PrescriptionService) ListByPatient(ctx context.Context, patientID int) ([]models.Prescription, error) {
	return s.db.ListPrescriptionsByPatient(ctx, patientID)
}

// Delete removes a prescription by ID
func (s *PrescriptionService) Delete(ctx context.Context, id int) error {
	return s.db.DeletePrescription(ctx, id)
}
