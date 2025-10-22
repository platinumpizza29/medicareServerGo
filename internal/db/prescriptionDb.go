package db

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/platinumpizza29/medicare/internal/models"
)

type PrescriptionDB struct {
	Pool *pgxpool.Pool
}

func NewPrescriptionDB(pool *pgxpool.Pool) *PrescriptionDB {
	return &PrescriptionDB{Pool: pool}
}

// CreatePrescription inserts a new prescription record
func (db *PrescriptionDB) CreatePrescription(ctx context.Context, p *models.Prescription) error {
	query := `
		INSERT INTO prescriptions (
			doctor_id, patient_id, diagnosis, medications, instructions, follow_up_date, created_at, updated_at
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id, created_at, updated_at
	`

	return db.Pool.QueryRow(ctx, query,
		p.DoctorID,
		p.PatientID,
		p.Diagnosis,
		p.Medications,
		p.Instructions,
		p.FollowUpDate,
		time.Now(),
		time.Now(),
	).Scan(&p.ID, &p.CreatedAt, &p.UpdatedAt)
}

// GetPrescriptionByID returns a prescription by ID
func (db *PrescriptionDB) GetPrescriptionByID(ctx context.Context, id int) (*models.Prescription, error) {
	query := `
		SELECT id, doctor_id, patient_id, diagnosis, medications, instructions, follow_up_date, created_at, updated_at
		FROM prescriptions
		WHERE id = $1
	`

	var p models.Prescription
	err := db.Pool.QueryRow(ctx, query, id).Scan(
		&p.ID,
		&p.DoctorID,
		&p.PatientID,
		&p.Diagnosis,
		&p.Medications,
		&p.Instructions,
		&p.FollowUpDate,
		&p.CreatedAt,
		&p.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &p, nil
}

// ListPrescriptionsByPatient returns all prescriptions for a specific patient
func (db *PrescriptionDB) ListPrescriptionsByPatient(ctx context.Context, patientID int) ([]models.Prescription, error) {
	query := `
		SELECT id, doctor_id, patient_id, diagnosis, medications, instructions, follow_up_date, created_at, updated_at
		FROM prescriptions
		WHERE patient_id = $1
		ORDER BY created_at DESC
	`

	rows, err := db.Pool.Query(ctx, query, patientID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var prescriptions []models.Prescription
	for rows.Next() {
		var p models.Prescription
		err := rows.Scan(
			&p.ID,
			&p.DoctorID,
			&p.PatientID,
			&p.Diagnosis,
			&p.Medications,
			&p.Instructions,
			&p.FollowUpDate,
			&p.CreatedAt,
			&p.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		prescriptions = append(prescriptions, p)
	}

	return prescriptions, nil
}

// DeletePrescription removes a prescription by ID
func (db *PrescriptionDB) DeletePrescription(ctx context.Context, id int) error {
	_, err := db.Pool.Exec(ctx, "DELETE FROM prescriptions WHERE id = $1", id)
	return err
}
