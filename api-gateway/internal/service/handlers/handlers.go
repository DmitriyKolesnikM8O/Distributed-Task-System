package handlers

import (
	"net/http"

	"github.com/KolesnikM8O/distributed-task-system/api-gateway/internal/config/config"
	"github.com/KolesnikM8O/distributed-task-system/api-gateway/internal/middleware"
	"github.com/KolesnikM8O/distributed-task-system/api-gateway/internal/service/url"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5"
)

type Service struct {
	db  *pgx.Conn
	rdb *redis.Client
}

func New(db *pgx.Conn, rdb *redis.Client) *Service {
	return &Service{
		db:  db,
		rdb: rdb,
	}
}

func (s *Service) Register(r *mux.Router, cfg *config.Config) {

	r.HandleFunc(url.TaskURL, middleware.JWTMiddleware(cfg, func(w http.ResponseWriter, r *http.Request) {
		s.CreateTaskHandler(w, r)
	})).Methods("POST")

	r.HandleFunc(url.TaskIDURL, middleware.JWTMiddleware(cfg, func(w http.ResponseWriter, r *http.Request) {
		s.GetTaskHandler(w, r)
	})).Methods("GET")

	r.HandleFunc(url.TaskIDURL, middleware.JWTMiddleware(cfg, func(w http.ResponseWriter, r *http.Request) {
		s.UpdateTaskHandler(w, r)
	})).Methods("PUT")

	r.HandleFunc(url.TaskIDURL, middleware.JWTMiddleware(cfg, func(w http.ResponseWriter, r *http.Request) {
		s.DeleteTaskHandler(w, r)
	})).Methods("DELETE")
}
