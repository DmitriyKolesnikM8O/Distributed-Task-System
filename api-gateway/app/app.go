package app

import (
	"context"
	"log"
	"net/http"

	"github.com/KolesnikM8O/distributed-task-system/api-gateway/redis"
	repository "github.com/KolesnikM8O/distributed-task-system/api-gateway/repository/postgreSQL"
	"github.com/KolesnikM8O/distributed-task-system/api-gateway/service/handlers"
	"github.com/gorilla/mux"
)

func Start(r *mux.Router) {
	log.Printf("Start app")

	log.Printf("Init redis")
	redis.InitRedis()

	repository := repository.New()
	db, err := repository.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close(context.Background())
	log.Printf("DB connected")

	service := handlers.New(db)
	service.Register(r)
	log.Printf("Listening on :8080")
	http.ListenAndServe(":8080", r)
}
