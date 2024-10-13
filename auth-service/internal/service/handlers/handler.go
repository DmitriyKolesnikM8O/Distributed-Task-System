package handlers

import (
	"net/http"

	"github.com/KolesnikM8O/distributed-task-system/auth-service/internal/service/url"
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

	r.HandleFunc(url.SignupURL, func(w http.ResponseWriter, r *http.Request) {
		s.SignupHandler(w, r)
	}).Methods("POST")

	r.HandleFunc(url.LoginURL, func(w http.ResponseWriter, r *http.Request) {
		s.LoginHandler(w, r)
	}).Methods("POST")
}
