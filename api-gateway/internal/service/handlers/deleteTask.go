package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Service) DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err := s.db.Exec(context.Background(), "DELETE FROM tasks WHERE id = $1", id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("Deleted task %s", id)
	w.WriteHeader(http.StatusOK)
}
