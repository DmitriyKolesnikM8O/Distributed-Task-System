package app

import (
	"context"
	"log"
	"net/http"

	"github.com/KolesnikM8O/distributed-task-system/auth-service/internal/config/config"
	repository "github.com/KolesnikM8O/distributed-task-system/auth-service/internal/repository/postgreSQL"
	"github.com/KolesnikM8O/distributed-task-system/auth-service/internal/service/handlers"
	"github.com/gorilla/mux"
)

func Start(r *mux.Router) {
	log.Printf("Start auth-service")

	log.Printf("Connecting to DB")

	cfg := config.GetConfig()

	repository := repository.New()
	db, err := repository.InitDB(&cfg.Storage)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close(context.Background())

	service := handlers.New(db)
	service.Register(r)

	port := cfg.Listen.Port
	log.Printf("Listening on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
