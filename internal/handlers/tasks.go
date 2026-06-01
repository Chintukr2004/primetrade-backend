package handlers

import (
	"encoding/json"
	"net/http"
	"primetrade-backend/internal/database"
	"primetrade-backend/internal/models"
)

// CreateTask adds a new task for the authenticated user
func CreateTask(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(int)

	var task models.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err := database.DB.QueryRow(
		"INSERT INTO tasks (user_id, title, status) VALUES ($1, $2, $3) RETURNING id, created_at",
		userID, task.Title, "pending",
	).Scan(&task.ID, &task.CreatedAt)

	if err != nil {
		http.Error(w, "Error creating task", http.StatusInternalServerError)
		return
	}

	task.UserID = userID
	task.Status = "pending"

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

// GetTasks retrieves all tasks for the authenticated user
func GetTasks(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(int)

	rows, err := database.DB.Query("SELECT id, title, status, created_at FROM tasks WHERE user_id = $1", userID)
	if err != nil {
		http.Error(w, "Error retrieving tasks", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var t models.Task
		if err := rows.Scan(&t.ID, &t.Title, &t.Status, &t.CreatedAt); err != nil {
			continue
		}
		t.UserID = userID
		tasks = append(tasks, t)
	}

	// If tasks is nil, return an empty array instead of null
	if tasks == nil {
		tasks = []models.Task{}
	}

	json.NewEncoder(w).Encode(tasks)
}

// UpdateTask modifies an existing task's status
func UpdateTask(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(int)
	taskID := r.PathValue("id") // Go 1.22 feature

	var task models.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	res, err := database.DB.Exec("UPDATE tasks SET status = $1 WHERE id = $2 AND user_id = $3", task.Status, taskID, userID)
	if err != nil {
		http.Error(w, "Error updating task", http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "Task not found or unauthorized", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Task updated successfully"})
}

// DeleteTask removes a task
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(int)
	taskID := r.PathValue("id")

	res, err := database.DB.Exec("DELETE FROM tasks WHERE id = $1 AND user_id = $2", taskID, userID)
	if err != nil {
		http.Error(w, "Error deleting task", http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "Task not found or unauthorized", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Task deleted successfully"})
}