package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"

	"github.com/platinumpizza29/medicare/internal/models"
	"github.com/platinumpizza29/medicare/internal/services"
	"github.com/platinumpizza29/medicare/internal/utils"
)

type DoctorHandler struct {
	DoctorService  services.DoctorService
	VisitService   *services.VisitService
	PatientService *services.PatientService
}

func NewDoctorHandler(
	doctorService services.DoctorService,
) *DoctorHandler {
	return &DoctorHandler{
		DoctorService: doctorService,
	}
}

// RegisterDoctorHandler godoc
// @Summary Register a new doctor
// @Tags doctor
// @Accept json
// @Produce json
// @Param request body DoctorRegistration true "Doctor registration"
// @Success 200 {object} AuthToken
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /doctor/auth/register [post]
func (h *DoctorHandler) RegisterDoctorHandler(w http.ResponseWriter, r *http.Request) {
	var doctor models.Doctor
	ctx := r.Context()

	if err := json.NewDecoder(r.Body).Decode(&doctor); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//hash the password and create a jwt token
	hashedPassword, err := utils.HashPasswords(doctor.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	doctor.Password = hashedPassword

	if err := h.DoctorService.Create(ctx, &doctor); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// return jwt token
	token, err := utils.CreateJwt(doctor.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

// LoginDoctorHandler godoc
// @Summary Login as a doctor
// @Tags doctor
// @Accept json
// @Produce json
// @Param request body DoctorLogin true "Doctor login"
// @Success 200 {object} DoctorLoginResponse
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Failure 500 {string} string
// @Router /doctor/auth/login [post]
func (h *DoctorHandler) LoginDoctorHandler(w http.ResponseWriter, r *http.Request) {
	var loginReq models.DoctorRequest
	ctx := r.Context()
	const recentPatientLimit = 5

	if err := json.NewDecoder(r.Body).Decode(&loginReq); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	doctorModel, err := h.DoctorService.GetByEmail(loginReq.Email, ctx)
	if err != nil {
		log.Printf("DB error while fetching doctor: %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if doctorModel == nil {
		http.Error(w, "doctor not found", http.StatusNotFound)
		return
	}

	// TODO: Verify password before generating token
	token, err := utils.CreateJwt(doctorModel.ID)
	if err != nil {
		http.Error(w, "failed to create token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	recentPatients, err := h.VisitService.GetRecentPatientsByDoctorID(
		ctx,
		doctorModel.ID,
		recentPatientLimit,
	)
	if err != nil {
		log.Printf("DB error while fetching recent patients: %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]any{
		"token":           token,
		"recent_patients": recentPatients,
	})
}

// GetPatientProfileByID godoc
// @Summary Get patient profile by ID (doctor view)
// @Tags doctor
// @Produce json
// @Param patientID path int true "Patient ID"
// @Success 200 {object} models.PatientProfile
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Failure 500 {string} string
// @Router /doctor/patients/{patientID} [get]
func (h *DoctorHandler) GetPatientProfileByID(w http.ResponseWriter, r *http.Request) {
	token, err := bearerToken(r)
	if err != nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	if _, err := utils.VerifyJWT(token); err != nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	patientID, err := strconv.Atoi(chi.URLParam(r, "patientID"))
	if err != nil {
		http.Error(w, "invalid patient ID", http.StatusBadRequest)
		return
	}

	patient, err := h.PatientService.GetByID(r.Context(), patientID)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	if patient == nil {
		http.Error(w, "patient not found", http.StatusNotFound)
		return
	}

	profile := models.PatientProfile{
		ID:               patient.ID,
		FirstName:        patient.FirstName,
		LastName:         patient.LastName,
		Address:          patient.Address,
		MobileNumber:     patient.MobileNumber,
		AadharNumber:     patient.AadharNumber,
		DOB:              patient.DOB,
		Gender:           patient.Gender,
		BloodGroup:       patient.BloodGroup,
		EmergencyContact: patient.EmergencyContact,
		Email:            patient.Email,
		CreatedAt:        patient.CreatedAt,
		UpdatedAt:        patient.UpdatedAt,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profile)
}

func bearerToken(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("missing authorization header")
	}

	const prefix = "Bearer "
	if !strings.HasPrefix(authHeader, prefix) {
		return "", errors.New("invalid authorization header")
	}

	token := strings.TrimSpace(strings.TrimPrefix(authHeader, prefix))
	if token == "" {
		return "", errors.New("missing bearer token")
	}

	return token, nil
}

// func LogoutDoctorHandler(w http.ResponseWriter, r *http.Request) {
// 	// Implementation of LogoutDoctorHandler
// }
