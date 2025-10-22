package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/platinumpizza29/medicare/internal/models"
	"github.com/platinumpizza29/medicare/internal/services"
)

type PrescriptionHandler struct {
	service *services.PrescriptionService
}

// NewPrescriptionHandler creates a new handler instance
func NewPrescriptionHandler(service *services.PrescriptionService) *PrescriptionHandler {
	return &PrescriptionHandler{service: service}
}

// RegisterRoutes registers prescription-related routes
func (h *PrescriptionHandler) RegisterRoutes(r chi.Router) {
	r.Route("/api/prescriptions", func(r chi.Router) {
		r.Post("/", h.CreatePrescription)
		r.Get("/{id}", h.GetPrescriptionByID)
		r.Get("/patient/{patientID}", h.ListByPatient)
		r.Delete("/{id}", h.DeletePrescription)
	})
}

// CreatePrescription handles POST /api/prescriptions
func (h *PrescriptionHandler) CreatePrescription(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var p models.Prescription

	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.service.Create(ctx, &p); err != nil {
		http.Error(w, "Failed to create prescription: "+err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusCreated, map[string]any{
		"message":      "Prescription created successfully",
		"prescription": p,
	})
}

// GetPrescriptionByID handles GET /api/prescriptions/{id}
func (h *PrescriptionHandler) GetPrescriptionByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid prescription ID", http.StatusBadRequest)
		return
	}

	p, err := h.service.GetByID(ctx, id)
	if err != nil {
		http.Error(w, "Error fetching prescription: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if p == nil {
		http.Error(w, "Prescription not found", http.StatusNotFound)
		return
	}

	writeJSON(w, http.StatusOK, p)
}

// ListByPatient handles GET /api/prescriptions/patient/{patientID}
func (h *PrescriptionHandler) ListByPatient(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	patientIDStr := chi.URLParam(r, "patientID")
	patientID, err := strconv.Atoi(patientIDStr)
	if err != nil {
		http.Error(w, "Invalid patient ID", http.StatusBadRequest)
		return
	}

	prescriptions, err := h.service.ListByPatient(ctx, patientID)
	if err != nil {
		http.Error(w, "Error fetching prescriptions: "+err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusOK, prescriptions)
}

// DeletePrescription handles DELETE /api/prescriptions/{id}
func (h *PrescriptionHandler) DeletePrescription(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid prescription ID", http.StatusBadRequest)
		return
	}

	if err := h.service.Delete(ctx, id); err != nil {
		http.Error(w, "Failed to delete prescription: "+err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{"message": "Prescription deleted successfully"})
}

// Helper function to respond with JSON
func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
