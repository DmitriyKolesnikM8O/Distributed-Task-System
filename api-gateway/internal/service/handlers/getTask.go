package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Service) GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	val, err := s.rdb.Get(s.rdb.Context(), id).Result()
	if err != nil {
		if err.Error() == "redis: nil" {
			var status string
			err := s.db.QueryRow(context.Background(), "SELECT status FROM tasks WHERE id = $1", id).Scan(&status)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			err = s.rdb.Set(s.rdb.Context(), id, status, 0).Err()
			if err != nil {
				log.Printf("Ошибка сохранения записи в Redis: %s", err)
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			log.Printf("Got task %s", id)
			w.Write([]byte(status))
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			log.Printf("Ошибка работы с Redis: %s", err)
		}
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		log.Printf("Got task from Redis %s", id)
		w.Write([]byte(val))
	}

}
