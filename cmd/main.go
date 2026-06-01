package main

import (
	"log"
	"net/http"
	"primetrade-backend/internal/database"
	"primetrade-backend/internal/handlers"
	"primetrade-backend/internal/middleware"
)

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Browsers send an "OPTIONS" request first to check permissions (Preflight)
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

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
	err := http.ListenAndServe(":8080", enableCORS(mux))
	if err != nil {
		log.Fatal("server failed to start: %v", err)
	}
}
