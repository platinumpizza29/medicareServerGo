package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"

	"github.com/platinumpizza29/medicare/internal/db"
	"github.com/platinumpizza29/medicare/internal/handlers"
	"github.com/platinumpizza29/medicare/internal/services"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error Loading .env file")
	}

	url := os.Getenv("DATABASE_URL")
	var ctx = context.TODO()

	// connect your database
	if err := db.ConnectDB(ctx, url); err != nil {
		log.Fatal("error connecting to the database:", err)
	}

	pool := db.Pool()
	router := chi.NewRouter()

	doctderDb := db.NewDoctorDB(pool)
	docterService := services.NewDoctorService(doctderDb)
	docterHandler := handlers.NewDoctorHandler(*docterService)

	patientDB := db.NewPatientDb(pool)
	patientService := services.NewPatientService(patientDB)
	patientHandler := handlers.NewPatientHandler(*patientService)

	presDB := db.NewPrescriptionDB(pool)
	presService := services.NewPrescriptionService(presDB)
	presHandler := handlers.NewPrescriptionHandler(presService)

	//doctor routes
	router.Route("/v1/doctor", func(r chi.Router) {
		r.Post("/auth/register", docterHandler.RegisterDoctorHandler)
		r.Post("/auth/login", docterHandler.LoginDoctorHandler)
		// r.Post("/auth/logout", handlers.LogoutDoctorHandler)
	})

	//patient routes
	router.Route("/v1/patient", func(r chi.Router) {
		r.Post("/auth/register", patientHandler.RegisterPatientHandler)
		r.Post("/auth/login", patientHandler.LoginPatientHandler)
		r.Post("/auth/logout", patientHandler.LoginPatientHandler)
	})

	//prescription routes
	router.Route("/v1/prescriptions", func(r chi.Router) {
		r.Post("/", presHandler.CreatePrescription)              // POST /v1/prescriptions
		r.Get("/{id}", presHandler.GetPrescriptionByID)          // GET /v1/prescriptions/1
		r.Get("/patient/{patientID}", presHandler.ListByPatient) // GET /v1/prescriptions/patient/5
		r.Delete("/{id}", presHandler.DeletePrescription)        // DELETE /v1/prescriptions/1
	})

	log.Fatal(http.ListenAndServe(":8080", router))
}
