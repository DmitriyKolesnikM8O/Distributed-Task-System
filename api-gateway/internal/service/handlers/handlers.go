package handlers

import (
	"net/http"

	"github.com/KolesnikM8O/distributed-task-system/api-gateway/internal/middleware"
	"github.com/KolesnikM8O/distributed-task-system/api-gateway/internal/service/url"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5"
)

type service struct {
	db *pgx.Conn
}

func New(db *pgx.Conn) *service {
	return &service{
		db: db,
	}
}

func (s *service) Register(r *mux.Router) {

	r.HandleFunc(url.TaskURL, middleware.JWTMiddleware(func(w http.ResponseWriter, r *http.Request) {
		s.CreateTaskHandler(w, r)
	})).Methods("POST")

	r.HandleFunc(url.TaskIDURL, middleware.JWTMiddleware(func(w http.ResponseWriter, r *http.Request) {
		s.GetTaskHandler(w, r)
	})).Methods("GET")

	r.HandleFunc(url.TaskIDURL, middleware.JWTMiddleware(func(w http.ResponseWriter, r *http.Request) {
		s.UpdateTaskHandler(w, r)
	})).Methods("PUT")

	r.HandleFunc(url.TaskIDURL, middleware.JWTMiddleware(func(w http.ResponseWriter, r *http.Request) {
		s.DeleteTaskHandler(w, r)
	})).Methods("DELETE")
}
