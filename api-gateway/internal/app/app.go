package app

import (
	"context"
	"log"
	"net/http"

	"github.com/KolesnikM8O/distributed-task-system/api-gateway/internal/config/config"
	"github.com/KolesnikM8O/distributed-task-system/api-gateway/internal/redis"
	repository "github.com/KolesnikM8O/distributed-task-system/api-gateway/internal/repository/postgreSQL"
	"github.com/KolesnikM8O/distributed-task-system/api-gateway/internal/service/handlers"
	"github.com/gorilla/mux"
)

func Start(r *mux.Router) {
	log.Printf("Start app")

	cfg := config.GetConfig()

	log.Printf("Init redis")
	redis.InitRedis(&cfg.Redis)

	repository := repository.New()
	db, err := repository.InitDB(&cfg.Storage)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close(context.Background())
	log.Printf("DB connected")

	service := handlers.New(db)
	service.Register(r)

	port := cfg.Listen.Port
	log.Printf("Listening on :%s", port)
	http.ListenAndServe(":"+port, r)
}
