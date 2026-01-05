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
	httpSwagger "github.com/swaggo/http-swagger/v2"

	"github.com/platinumpizza29/medicare/docs"
	"github.com/platinumpizza29/medicare/internal/db"
	"github.com/platinumpizza29/medicare/internal/handlers"
	"github.com/platinumpizza29/medicare/internal/services"
)

func main() {
	//if err := godotenv.Load(); err != nil {
	//log.Fatal("Error Loading .env file")
	//}

	url := os.Getenv("DATABASE_URL")
	var ctx = context.TODO()

	// connect your database
	if err := db.ConnectDB(ctx, url); err != nil {
		log.Fatal("error connecting to the database:", err)
	}

	pool := db.Pool()
	router := chi.NewRouter()

	docs.SwaggerInfo.Title = "Medicare API"
	docs.SwaggerInfo.Description = "API documentation for the Medicare service."
	docs.SwaggerInfo.Version = "1.0.0"
	docs.SwaggerInfo.BasePath = "/v1"
	if swaggerHost := os.Getenv("SWAGGER_HOST"); swaggerHost != "" {
		docs.SwaggerInfo.Host = swaggerHost
	}

	doctderDb := db.NewDoctorDB(pool)
	docterService := services.NewDoctorService(doctderDb)
	docterHandler := handlers.NewDoctorHandler(*docterService)

	patientDB := db.NewPatientDb(pool)
	patientService := services.NewPatientService(patientDB)
	patientHandler := handlers.NewPatientHandler(*patientService)

	presDB := db.NewPrescriptionDB(pool)
	presService := services.NewPrescriptionService(presDB)
	presHandler := handlers.NewPrescriptionHandler(presService)

	visitsDB := db.NewVisitsDB(pool)
	visitsService := services.NewVisitService(visitsDB)
	visitsHandler := handlers.NewVisitHandler(visitsService)

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

	//visits routes
	router.Route("/v1/visits", func(r chi.Router) {
		r.Post("/", visitsHandler.CreateVisit)
		r.Get("/{id}", visitsHandler.GetVisitByID)
		r.Get("/patient/{patientID}", visitsHandler.GetVisitsByPatientID)
		r.Get("/doctor/{doctorID}", visitsHandler.GetVisitsByDoctorID)
		r.Put("/{id}", visitsHandler.UpdateVisit)
		r.Delete("/{id}", visitsHandler.DeleteVisit)
	})

	router.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	router.Get("/swagger/*", httpSwagger.WrapHandler)

	log.Fatal(http.ListenAndServe(":3000", router))
}
