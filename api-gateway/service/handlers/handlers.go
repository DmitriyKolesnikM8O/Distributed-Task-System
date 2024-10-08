package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/KolesnikM8O/distributed-task-system/api-gateway/redis"
	"github.com/KolesnikM8O/distributed-task-system/api-gateway/service/model"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5"
)

func CreateTaskHandler(w http.ResponseWriter, r *http.Request, db *pgx.Conn) {
	task := model.Task{
		Status: r.URL.Query().Get("status"),
	}
	if task.Status == "" {
		task.Status = "pending"
	}

	task.ID = uuid.New().String()
	_, err := db.Exec(context.Background(), "INSERT INTO tasks (id, status) VALUES ($1, $2)", task.ID, task.Status)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("Created task %s", task.ID)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func GetTaskHandler(w http.ResponseWriter, r *http.Request, db *pgx.Conn) {
	id := mux.Vars(r)["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	val, err := redis.RDB.Get(redis.Ctx, id).Result()
	if err != nil {
		if err.Error() == "redis: nil" {
			var status string
			err := db.QueryRow(context.Background(), "SELECT status FROM tasks WHERE id = $1", id).Scan(&status)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			err = redis.RDB.Set(redis.Ctx, id, status, 0).Err()
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

func UpdateTaskHandler(w http.ResponseWriter, r *http.Request, db *pgx.Conn) {
	id := mux.Vars(r)["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var status string
	err := db.QueryRow(context.Background(), "SELECT status FROM tasks WHERE id = $1", id).Scan(&status)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	status = r.URL.Query().Get("status")
	if status == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = db.Exec(context.Background(), "UPDATE tasks SET status = $1 WHERE id = $2", status, id)
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

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request, db *pgx.Conn) {
	id := mux.Vars(r)["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err := db.Exec(context.Background(), "DELETE FROM tasks WHERE id = $1", id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("Deleted task %s", id)
	w.WriteHeader(http.StatusOK)
}
