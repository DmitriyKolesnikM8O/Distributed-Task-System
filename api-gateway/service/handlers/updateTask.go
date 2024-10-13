package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/KolesnikM8O/distributed-task-system/api-gateway/redis"
	"github.com/gorilla/mux"
)

func (s *service) UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var status string
	err := s.db.QueryRow(context.Background(), "SELECT status FROM tasks WHERE id = $1", id).Scan(&status)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	status = r.URL.Query().Get("status")
	if status == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = s.db.Exec(context.Background(), "UPDATE tasks SET status = $1 WHERE id = $2", status, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = redis.RDB.Set(redis.Ctx, id, status, 0).Err()
	if err != nil {
		log.Printf("Error saving task to Redis: %s", err)
	}

	log.Printf("Updated task %s", id)
	w.WriteHeader(http.StatusOK)
}
