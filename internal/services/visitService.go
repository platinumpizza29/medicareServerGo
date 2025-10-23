package services

import (
	"context"

	"github.com/platinumpizza29/medicare/internal/db"
	"github.com/platinumpizza29/medicare/internal/models"
)

type VisitService struct {
	db *db.VisitsDB
}

func NewVisitService(visitDb *db.VisitsDB) *VisitService {
	return &VisitService{
		db: visitDb,
	}
}

func (s *VisitService) CreateVisit(ctx context.Context, visit *models.Visit) error {
	return s.db.CreateVisit(ctx, visit)
}

func (s *VisitService) GetVisitByID(ctx context.Context, id int) (*models.Visit, error) {
	return s.db.GetVisitByID(ctx, id)
}

func (s *VisitService) GetVisitsByPatientID(ctx context.Context, patientID int) ([]*models.Visit, error) {
	return s.db.GetVisitsByPatientID(ctx, patientID)
}

func (s *VisitService) GetVisitsByDoctorID(ctx context.Context, doctorID int) ([]*models.Visit, error) {
	return s.db.GetVisitsByDoctorID(ctx, doctorID)
}

func (s *VisitService) UpdateVisit(ctx context.Context, visit *models.Visit) error {
	return s.db.UpdateVisit(ctx, visit)
}

func (s *VisitService) DeleteVisit(ctx context.Context, id int) error {
	return s.db.DeleteVisit(ctx, id)
}
