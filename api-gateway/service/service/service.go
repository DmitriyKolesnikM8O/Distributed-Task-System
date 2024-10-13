package service

import "github.com/gorilla/mux"

type Service interface {
	Register(r *mux.Router)
}
