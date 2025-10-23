package db

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/platinumpizza29/medicare/internal/models"
)

type VisitsDB struct {
	Pool *pgxpool.Pool
}

func NewVisitsDB(pool *pgxpool.Pool) *VisitsDB {
	return &VisitsDB{Pool: pool}
}

func (db *VisitsDB) CreateVisit(ctx context.Context, v *models.Visit) error {
	query := `
		INSERT INTO visits (patient_id, doctor_id, visit_date, notes, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, created_at, updated_at`

	return db.Pool.QueryRow(ctx, query,
		v.PatientID,
		v.DoctorID,
		v.VisitDate,
		v.Notes,
		time.Now(),
		time.Now(),
	).Scan(&v.ID, &v.CreatedAt, &v.UpdatedAt)
}

func (db *VisitsDB) GetVisitByID(ctx context.Context, id int) (*models.Visit, error) {
	query := `
		SELECT id, patient_id, doctor_id, visit_date, notes, created_at, updated_at
		FROM visits
		WHERE id = $1`

	var v models.Visit
	err := db.Pool.QueryRow(ctx, query, id).Scan(
		&v.ID,
		&v.PatientID,
		&v.DoctorID,
		&v.VisitDate,
		&v.Notes,
		&v.CreatedAt,
		&v.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &v, nil
}

func (db *VisitsDB) GetVisitsByPatientID(ctx context.Context, patientID int) ([]*models.Visit, error) {
	query := `
		SELECT id, patient_id, doctor_id, visit_date, notes, created_at, updated_at
		FROM visits
		WHERE patient_id = $1`

	rows, err := db.Pool.Query(ctx, query, patientID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var visits []*models.Visit
	for rows.Next() {
		var v models.Visit
		err := rows.Scan(
			&v.ID,
			&v.PatientID,
			&v.DoctorID,
			&v.VisitDate,
			&v.Notes,
			&v.CreatedAt,
			&v.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		visits = append(visits, &v)
	}

	return visits, nil
}

func (db *VisitsDB) GetVisitsByDoctorID(ctx context.Context, doctorID int) ([]*models.Visit, error) {
	query := `
		SELECT id, patient_id, doctor_id, visit_date, notes, created_at, updated_at
		FROM visits
		WHERE doctor_id = $1`

	rows, err := db.Pool.Query(ctx, query, doctorID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var visits []*models.Visit
	for rows.Next() {
		var v models.Visit
		err := rows.Scan(
			&v.ID,
			&v.PatientID,
			&v.DoctorID,
			&v.VisitDate,
			&v.Notes,
			&v.CreatedAt,
			&v.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		visits = append(visits, &v)
	}

	return visits, nil
}

func (db *VisitsDB) UpdateVisit(ctx context.Context, v *models.Visit) error {
	query := `
		UPDATE visits
		SET patient_id = $1, doctor_id = $2, visit_date = $3, notes = $4, updated_at = $5
		WHERE id = $6`

	_, err := db.Pool.Exec(ctx, query,
		v.PatientID,
		v.DoctorID,
		v.VisitDate,
		v.Notes,
		time.Now(),
		v.ID,
	)

	return err
}

func (db *VisitsDB) DeleteVisit(ctx context.Context, id int) error {
	query := `DELETE FROM visits WHERE id = $1`
	_, err := db.Pool.Exec(ctx, query, id)
	return err
}