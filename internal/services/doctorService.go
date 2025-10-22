package services

import (
	"context"

	"github.com/platinumpizza29/medicare/internal/db"
	"github.com/platinumpizza29/medicare/internal/models"
)

type DoctorService struct {
	db *db.DoctorDB
}

func NewDoctorService(doctorDb *db.DoctorDB) *DoctorService {
	return &DoctorService{
		db: doctorDb,
	}
}

func (s *DoctorService) Create(ctx context.Context, doctor *models.Doctor) error {
	return s.db.CreateDoctor(ctx, doctor)
}

func (s *DoctorService) GetByEmail(email string, ctx context.Context) (*models.Doctor, error) {
	doctorModel, err := s.db.GetDoctorByEmail(ctx, email)

	if err != nil {
		return nil, err
	}
	return doctorModel, nil
}

