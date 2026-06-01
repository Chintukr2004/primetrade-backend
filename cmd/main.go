package main

import (
	"log"
	"net/http"
	"primetrade-backend/internal/database"
	"primetrade-backend/internal/handlers"
	"primetrade-backend/internal/middleware"
)

func main() {
	database.ConnectDB()

	mux := http.NewServeMux()

	mux.HandleFunc("POST /api/v1/register", handlers.Register)
	mux.HandleFunc("POST /api/v1/login", handlers.Login)

	mux.HandleFunc("POST /api/v1/tasks", middleware.Authenticate(handlers.CreateTask))
	mux.HandleFunc("GET /api/v1/tasks", middleware.Authenticate(handlers.GetTasks))
	mux.HandleFunc("PUT /api/v1/tasks/{id}", middleware.Authenticate(handlers.UpdateTask))
	mux.HandleFunc("DELETE /api/v1/tasks/{id}", middleware.Authenticate(handlers.DeleteTask))

	log.Println("Server starting on port 8080...")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal("server failed to start: %v", err)
	}
}
