package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/platinumpizza29/medicare/internal/models"
	"github.com/platinumpizza29/medicare/internal/services"
	"github.com/platinumpizza29/medicare/internal/utils"
)

type PatientHandler struct {
	PatientService services.PatientService
}

func NewPatientHandler(patientService services.PatientService) *PatientHandler {
	return &PatientHandler{
		PatientService: patientService,
	}
}

// RegisterPatientHandler handles patient registration
// @Summary Register a new patient
// @Tags patient
// @Accept json
// @Produce json
// @Param request body PatientRegistration true "Patient registration"
// @Success 200 {object} AuthToken
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /patient/auth/register [post]
func (h *PatientHandler) RegisterPatientHandler(w http.ResponseWriter, r *http.Request) {
	var patient models.Patient
	ctx := r.Context()

	if err := json.NewDecoder(r.Body).Decode(&patient); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	// hash password
	hashedPassword, err := utils.HashPasswords(patient.Password)
	if err != nil {
		http.Error(w, "failed to hash password", http.StatusInternalServerError)
		return
	}
	patient.Password = hashedPassword

	// create patient record
	if err := h.PatientService.Create(ctx, &patient); err != nil {
		http.Error(w, "failed to register patient", http.StatusInternalServerError)
		return
	}

	// generate JWT token
	token, err := utils.CreateJwt(patient.ID)
	if err != nil {
		http.Error(w, "failed to create token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
}

// LoginPatientHandler handles patient login
// @Summary Login as a patient
// @Description Also handles /patient/auth/logout (currently routed to the same handler).
// @Tags patient
// @Accept json
// @Produce json
// @Param request body PatientLogin true "Patient login"
// @Success 200 {object} AuthToken
// @Failure 400 {string} string
// @Failure 401 {string} string
// @Failure 404 {string} string
// @Failure 500 {string} string
// @Router /patient/auth/login [post]
// @Router /patient/auth/logout [post]
func (h *PatientHandler) LoginPatientHandler(w http.ResponseWriter, r *http.Request) {
	var loginReq models.PatientRequest
	ctx := r.Context()

	if err := json.NewDecoder(r.Body).Decode(&loginReq); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	patientModel, err := h.PatientService.GetByEmail(ctx, loginReq.Email)
	if err != nil {
		log.Printf("DB error while fetching patient: %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if patientModel == nil {
		http.Error(w, "patient not found", http.StatusNotFound)
		return
	}

	// verify password
	if utils.ComparePasswords(loginReq.Password, patientModel.Password) {
		http.Error(w, "invalid password", http.StatusUnauthorized)
		return
	}

	// generate JWT
	token, err := utils.CreateJwt(patientModel.ID)
	if err != nil {
		http.Error(w, "failed to create token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
}
