package app

import (
	"context"
	"log"
	"net/http"

	repository "github.com/KolesnikM8O/distributed-task-system/auth-service/repository/postgreSQL"
	"github.com/KolesnikM8O/distributed-task-system/auth-service/service/handlers"
	"github.com/gorilla/mux"
)

func Start(r *mux.Router) {
	log.Printf("Start auth-service")

	log.Printf("Connecting to DB")

	repository := repository.New()
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
