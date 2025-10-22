package db

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/platinumpizza29/medicare/internal/models"
)

type PatientDB struct {
	Pool *pgxpool.Pool
}

func NewPatientDb(pool *pgxpool.Pool) *PatientDB {
	return &PatientDB{Pool: pool}
}

// CreatePatient inserts a new patient record
func (db *PatientDB) CreatePatient(ctx context.Context, p *models.Patient) error {
	query := `
        INSERT INTO patients (
            first_name, last_name, address, mobile_number, aadhar_number,
            dob, gender, blood_group, emergency_contact, email, password,
            created_at, updated_at
        )
        VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13)
        RETURNING id, created_at, updated_at
    `

	return db.Pool.QueryRow(ctx, query,
		p.FirstName,
		p.LastName,
		p.Address,
		p.MobileNumber,
		p.AadharNumber,
		p.DOB,
		p.Gender,
		p.BloodGroup,
		p.EmergencyContact,
		p.Email,
		p.Password, // hashed before calling this function
		time.Now(),
		time.Now(),
	).Scan(&p.ID, &p.CreatedAt, &p.UpdatedAt)
}

// GetPatientByEmail returns a patient by email
func (db *PatientDB) GetPatientByEmail(ctx context.Context, email string) (*models.Patient, error) {
	query := `
        SELECT id, first_name, last_name, address, mobile_number, aadhar_number,
               dob, gender, blood_group, emergency_contact, email, password,
               created_at, updated_at
        FROM patients
        WHERE email = $1
    `

	var p models.Patient
	err := db.Pool.QueryRow(ctx, query, email).Scan(
		&p.ID,
		&p.FirstName,
		&p.LastName,
		&p.Address,
		&p.MobileNumber,
		&p.AadharNumber,
		&p.DOB,
		&p.Gender,
		&p.BloodGroup,
		&p.EmergencyContact,
		&p.Email,
		&p.Password,
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

// GetPatientByID returns a patient by ID
func (db *PatientDB) GetPatientByID(ctx context.Context, id int) (*models.Patient, error) {
	query := `
        SELECT id, first_name, last_name, address, mobile_number, aadhar_number,
               dob, gender, blood_group, emergency_contact, email, password,
               created_at, updated_at
        FROM patients
        WHERE id = $1
    `

	var p models.Patient
	err := db.Pool.QueryRow(ctx, query, id).Scan(
		&p.ID,
		&p.FirstName,
		&p.LastName,
		&p.Address,
		&p.MobileNumber,
		&p.AadharNumber,
		&p.DOB,
		&p.Gender,
		&p.BloodGroup,
		&p.EmergencyContact,
		&p.Email,
		&p.Password,
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

// ListPatients returns all patients
func (db *PatientDB) ListPatients(ctx context.Context) ([]models.Patient, error) {
	query := `
        SELECT id, first_name, last_name, address, mobile_number, aadhar_number,
               dob, gender, blood_group, emergency_contact, email, password,
               created_at, updated_at
        FROM patients
        ORDER BY created_at DESC
    `
	rows, err := db.Pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var patients []models.Patient
	for rows.Next() {
		var p models.Patient
		err := rows.Scan(
			&p.ID,
			&p.FirstName,
			&p.LastName,
			&p.Address,
			&p.MobileNumber,
			&p.AadharNumber,
			&p.DOB,
			&p.Gender,
			&p.BloodGroup,
			&p.EmergencyContact,
			&p.Email,
			&p.Password,
			&p.CreatedAt,
			&p.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		patients = append(patients, p)
	}

	return patients, nil
}

// DeletePatient removes a patient by ID
func (db *PatientDB) DeletePatient(ctx context.Context, id int) error {
	_, err := db.Pool.Exec(ctx, "DELETE FROM patients WHERE id = $1", id)
	return err
}

