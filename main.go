// @title Medicare API
// @version 1.0.0
// @description API documentation for the Medicare service.
// @host localhost:3000
// @BasePath /v1
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

	patientDB := db.NewPatientDb(pool)
	patientService := services.NewPatientService(patientDB)
	patientHandler := handlers.NewPatientHandler(*patientService)

	doctderDb := db.NewDoctorDB(pool)
	docterService := services.NewDoctorService(doctderDb)
	docterHandler := handlers.NewDoctorHandler(*docterService)

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

	router.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	log.Fatal(http.ListenAndServe(":8080", router))
}
