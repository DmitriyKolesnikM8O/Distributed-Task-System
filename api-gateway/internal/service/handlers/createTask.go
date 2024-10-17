package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/KolesnikM8O/distributed-task-system/api-gateway/internal/service/model"
	"github.com/google/uuid"
)

func (s *Service) CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	task := model.Task{
		Status: r.URL.Query().Get("status"),
	}
	if task.Status == "" {
		task.Status = "pending"
	}

	task.ID = uuid.New().String()
	_, err := s.db.Exec(context.Background(), "INSERT INTO tasks (id, status) VALUES ($1, $2)", task.ID, task.Status)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("Created task %s", task.ID)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}
