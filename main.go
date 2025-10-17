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

	db := db.NewDoctorDB(pool)

	router := chi.NewRouter()

	//doctor routes
	router.Route("/v1/doctor", func(r chi.Router) {
		r.Post("/auth/register", handlers.RegisterDoctorHandler)
		// r.Post("/auth/login", handlers.LoginDoctorHandler)
		// r.Post("/auth/logout", handlers.LogoutDoctorHandler)
	})

	log.Fatal(http.ListenAndServe(":8080", router))
}
