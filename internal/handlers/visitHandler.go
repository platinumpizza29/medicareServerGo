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
