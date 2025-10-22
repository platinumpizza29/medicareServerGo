package services

import (
	"context"

	"github.com/platinumpizza29/medicare/internal/db"
	"github.com/platinumpizza29/medicare/internal/models"
)

type PatientService struct {
	db *db.PatientDB
}

func NewPatientService(patientDb *db.PatientDB) *PatientService {
	return &PatientService{
		db: patientDb,
	}
}

// Create inserts a new patient record
func (s *PatientService) Create(ctx context.Context, patient *models.Patient) error {
	return s.db.CreatePatient(ctx, patient)
}

// GetByEmail retrieves a patient by email
func (s *PatientService) GetByEmail(ctx context.Context, email string) (*models.Patient, error) {
	patientModel, err := s.db.GetPatientByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return patientModel, nil
}

// GetByID retrieves a patient by ID
func (s *PatientService) GetByID(ctx context.Context, id int) (*models.Patient, error) {
	patientModel, err := s.db.GetPatientByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return patientModel, nil
}

// List retrieves all patients
func (s *PatientService) List(ctx context.Context) ([]models.Patient, error) {
	patients, err := s.db.ListPatients(ctx)
	if err != nil {
		return nil, err
	}
	return patients, nil
}

// Delete removes a patient by ID
func (s *PatientService) Delete(ctx context.Context, id int) error {
	return s.db.DeletePatient(ctx, id)
}
