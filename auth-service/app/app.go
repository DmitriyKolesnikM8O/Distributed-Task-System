package app

import (
	"context"
	"log"
	"net/http"

	"github.com/KolesnikM8O/distributed-task-system/auth-service/repository"
	"github.com/KolesnikM8O/distributed-task-system/auth-service/service/handlers"
	"github.com/gorilla/mux"
)

func Start(r *mux.Router) {
	log.Printf("Start auth-service")

	log.Printf("Connecting to DB")
	db, err := repository.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close(context.Background())

	service := handlers.New(db)
	service.Register(r)

	log.Printf("Listening on :8081")
	log.Fatal(http.ListenAndServe(":8081", r))
}
