package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/platinumpizza29/medicare/internal/models"
	"github.com/platinumpizza29/medicare/internal/services"
)

type VisitHandler struct {
	visitService *services.VisitService
}

func NewVisitHandler(visitService *services.VisitService) *VisitHandler {
	return &VisitHandler{
		visitService: visitService,
	}
}

// CreateVisit godoc
// @Summary Create a new visit
// @Tags visits
// @Accept json
// @Produce json
// @Param request body VisitInput true "Visit"
// @Success 201 {string} string
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /visits [post]
func (h *VisitHandler) CreateVisit(w http.ResponseWriter, r *http.Request) {
	var visit models.Visit
	if err := json.NewDecoder(r.Body).Decode(&visit); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.visitService.CreateVisit(r.Context(), &visit); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// GetVisitByID godoc
// @Summary Get a visit by ID
// @Tags visits
// @Produce json
// @Param id path int true "Visit ID"
// @Success 200 {object} models.Visit
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /visits/{id} [get]
func (h *VisitHandler) GetVisitByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	visit, err := h.visitService.GetVisitByID(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(visit)
}

// GetVisitsByPatientID godoc
// @Summary Get visits by patient ID
// @Tags visits
// @Produce json
// @Param patientID path int true "Patient ID"
// @Success 200 {array} models.Visit
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /visits/patient/{patientID} [get]
func (h *VisitHandler) GetVisitsByPatientID(w http.ResponseWriter, r *http.Request) {
	patientID, err := strconv.Atoi(chi.URLParam(r, "patientID"))
	if err != nil {
		http.Error(w, "Invalid Patient ID", http.StatusBadRequest)
		return
	}

	visits, err := h.visitService.GetVisitsByPatientID(r.Context(), patientID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(visits)
}

// GetVisitsByDoctorID godoc
// @Summary Get visits by doctor ID
// @Tags visits
// @Produce json
// @Param doctorID path int true "Doctor ID"
// @Success 200 {array} models.Visit
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /visits/doctor/{doctorID} [get]
func (h *VisitHandler) GetVisitsByDoctorID(w http.ResponseWriter, r *http.Request) {
	doctorID, err := strconv.Atoi(chi.URLParam(r, "doctorID"))
	if err != nil {
		http.Error(w, "Invalid Doctor ID", http.StatusBadRequest)
		return
	}

	visits, err := h.visitService.GetVisitsByDoctorID(r.Context(), doctorID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(visits)
}

// UpdateVisit godoc
// @Summary Update a visit
// @Tags visits
// @Accept json
// @Produce json
// @Param id path int true "Visit ID"
// @Param request body VisitInput true "Visit"
// @Success 200 {string} string
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /visits/{id} [put]
func (h *VisitHandler) UpdateVisit(w http.ResponseWriter, r *http.Request) {
	var visit models.Visit
	if err := json.NewDecoder(r.Body).Decode(&visit); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.visitService.UpdateVisit(r.Context(), &visit); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// DeleteVisit godoc
// @Summary Delete a visit
// @Tags visits
// @Produce json
// @Param id path int true "Visit ID"
// @Success 200 {string} string
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /visits/{id} [delete]
func (h *VisitHandler) DeleteVisit(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if err := h.visitService.DeleteVisit(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
