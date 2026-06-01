package main

import (
	"log"
	"net/http"
	"primetrade-backend/internal/database"
	"primetrade-backend/internal/handlers"
)

func main() {
	database.ConnectDB()

	mux := http.NewServeMux()

	mux.HandleFunc("POST /api/v1/register", handlers.Register)
	mux.HandleFunc("POST /api/v1/login", handlers.Login)

	log.Println("Server starting on port 8080...")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal("server failed to start: %v", err)
	}
}
