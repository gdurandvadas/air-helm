package main

import (
	"net/http"

	"github.com/gdurandvadas/air-helm/example/api/internal/api/handlers"
	"github.com/gdurandvadas/air-helm/example/api/internal/api/middlewares"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Setup subpath /api/v1
	api := r.PathPrefix("/api/v1").Subrouter()

	// Message handlers
	api.HandleFunc("/message", handlers.CreateMessage).Methods("POST")
	api.HandleFunc("/message/{id}", handlers.GetMessage).Methods("GET")

	// Middlewares
	r.Use(middlewares.AccessLog)

	// Start the server
	http.ListenAndServe(":8080", r)
}
