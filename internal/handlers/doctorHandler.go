package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/platinumpizza29/medicare/internal/models"
	"github.com/platinumpizza29/medicare/internal/services"
	"github.com/platinumpizza29/medicare/internal/utils"
)

type DoctorHandler struct {
	DoctorService services.DoctorService
}

func NewDoctorHandler(doctorService services.DoctorService) *DoctorHandler {
	return &DoctorHandler{
		DoctorService: doctorService,
	}
}

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

// func LoginDoctorHandler(w http.ResponseWriter, r *http.Request) {
// 	var doctor models.Doctor
// 	if err := json.NewDecoder(r.Body).Decode(&doctor); err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	if err := services.Login(&doctor); err != nil {
// 		http.Error(w, err.Error(), http.StatusUnauthorized)
// 		return
// 	}

// 	// return jwt token
// 	token, err := utils.CreateJwt(doctor.ID)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	json.NewEncoder(w).Encode(map[string]string{"token": token})
// }

// func LogoutDoctorHandler(w http.ResponseWriter, r *http.Request) {
// 	// Implementation of LogoutDoctorHandler
// }
