package services

import "github.com/platinumpizza29/medicare/internal/models"

func Create(doctor *models.Doctor) error {
	return nil
}

func GetByEmail(email string) (*models.Doctor, error) {
	// TODO: Implement compare passwords as well for security
	return nil, nil
}
