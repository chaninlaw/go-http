package main

import (
	"log"
	"net/http"

	"github.com/chaninlaw/auth/internal/api"
	"github.com/chaninlaw/auth/internal/repository"
	"github.com/chaninlaw/auth/internal/service"
	"github.com/chaninlaw/auth/pkg"
	"github.com/gorilla/mux"
)

func main() {
	pkg.LoadEnv(".env")

	dbPool, err := pkg.DBConnectPool()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer dbPool.Close()

	// Initialize the repository
	userRepo := repository.NewUserRepository(dbPool)

	// Initialize the service
	userService := service.NewUserService(userRepo)

	// Initialize the HTTP handler
	handler := api.NewHandler(userService)

	// Setup the HTTP server
	r := mux.NewRouter()
	r.Use(api.ApiKeyMiddleware)

	r.HandleFunc("/users", handler.GetUsers).Methods(http.MethodGet)
	r.HandleFunc("/user/{id}", handler.GetUser).Methods(http.MethodGet)
	r.HandleFunc("/user", handler.CreateUser).Methods(http.MethodPost)

	// Start the server
	httpAddr := ":8080"
	log.Printf("Server starting on %s", httpAddr)
	if err := http.ListenAndServe(httpAddr, r); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
