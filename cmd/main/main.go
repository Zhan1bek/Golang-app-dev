package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"tsis/internal/handlers"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/players", handlers.GetPlayers).Methods("GET")
	r.HandleFunc("/players/{id}", handlers.GetPlayerByID).Methods("GET")
	r.HandleFunc("players/{name}", handlers.GetPlayerByName).Methods("GET")
	r.HandleFunc("/health", handlers.HealthCheck).Methods("GET")

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}