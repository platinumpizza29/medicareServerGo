package db

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/platinumpizza29/medicare/internal/models"
)

type DoctorDB struct {
	Pool *pgxpool.Pool
}

func NewDoctorDB(pool *pgxpool.Pool) *DoctorDB {
	return &DoctorDB{Pool: pool}
}

// CreateDoctor inserts a new doctor record
func (db *DoctorDB) CreateDoctor(ctx context.Context, d *models.Doctor) error {
	query := `
        INSERT INTO doctors (
            first_name, last_name, clinic_name, address, license_number,
            mobile_number, aadhar_number, specialty, experience, dob,
            gender, blood_group, email, password, created_at, updated_at
        )
        VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16)
        RETURNING id, created_at, updated_at
    `

	return db.Pool.QueryRow(ctx, query,
		d.FirstName,
		d.LastName,
		d.ClinicName,
		d.Address,
		d.LicenseNumber,
		d.MobileNumber,
		d.AadharNumber,
		d.Specialty,
		d.Experience,
		d.DOB,
		d.Gender,
		d.BloodGroup,
		d.Email,
		d.Password, // hash before calling this function
		time.Now(),
		time.Now(),
	).Scan(&d.ID, &d.CreatedAt, &d.UpdatedAt)
}

// GetDoctorByEmail returns a doctor by email
func (db *DoctorDB) GetDoctorByEmail(ctx context.Context, email string) (*models.Doctor, error) {
	query := `
        SELECT id, first_name, last_name, clinic_name, address, license_number,
               mobile_number, aadhar_number, specialty, experience, dob,
               gender, blood_group, email, password, created_at, updated_at
        FROM doctors
        WHERE email = $1
    `

	var d models.Doctor
	err := db.Pool.QueryRow(ctx, query, email).Scan(
		&d.ID,
		&d.FirstName,
		&d.LastName,
		&d.ClinicName,
		&d.Address,
		&d.LicenseNumber,
		&d.MobileNumber,
		&d.AadharNumber,
		&d.Specialty,
		&d.Experience,
		&d.DOB,
		&d.Gender,
		&d.BloodGroup,
		&d.Email,
		&d.Password,
		&d.CreatedAt,
		&d.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &d, nil
}

// GetDoctorByID returns a doctor by ID
func (db *DoctorDB) GetDoctorByID(ctx context.Context, id int) (*models.Doctor, error) {
	query := `
        SELECT id, first_name, last_name, clinic_name, address, license_number,
               mobile_number, aadhar_number, specialty, experience, dob,
               gender, blood_group, email, password, created_at, updated_at
        FROM doctors
        WHERE id = $1
    `

	var d models.Doctor
	err := db.Pool.QueryRow(ctx, query, id).Scan(
		&d.ID,
		&d.FirstName,
		&d.LastName,
		&d.ClinicName,
		&d.Address,
		&d.LicenseNumber,
		&d.MobileNumber,
		&d.AadharNumber,
		&d.Specialty,
		&d.Experience,
		&d.DOB,
		&d.Gender,
		&d.BloodGroup,
		&d.Email,
		&d.Password,
		&d.CreatedAt,
		&d.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &d, nil
}

// ListDoctors returns all doctors
func (db *DoctorDB) ListDoctors(ctx context.Context) ([]models.Doctor, error) {
	query := `
        SELECT id, first_name, last_name, clinic_name, address, license_number,
               mobile_number, aadhar_number, specialty, experience, dob,
               gender, blood_group, email, password, created_at, updated_at
        FROM doctors
        ORDER BY created_at DESC
    `
	rows, err := db.Pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var doctors []models.Doctor
	for rows.Next() {
		var d models.Doctor
		err := rows.Scan(
			&d.ID,
			&d.FirstName,
			&d.LastName,
			&d.ClinicName,
			&d.Address,
			&d.LicenseNumber,
			&d.MobileNumber,
			&d.AadharNumber,
			&d.Specialty,
			&d.Experience,
			&d.DOB,
			&d.Gender,
			&d.BloodGroup,
			&d.Email,
			&d.Password,
			&d.CreatedAt,
			&d.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		doctors = append(doctors, d)
	}

	return doctors, nil
}

// DeleteDoctor removes a doctor by ID
func (db *DoctorDB) DeleteDoctor(ctx context.Context, id int) error {
	_, err := db.Pool.Exec(ctx, "DELETE FROM doctors WHERE id = $1", id)
	return err
}
