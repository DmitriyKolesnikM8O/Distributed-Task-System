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
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Start(r *mux.Router) {
	log.Printf("Start app")

	http.Handle("/metrics", promhttp.Handler())
	log.Println("Metrics server on :9090...")

	cfg := config.GetConfig()

	log.Printf("Init redis")
	rdb, err := redis.InitRedis(&cfg.Redis)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Redis connected")

	repository := repository.New()
	db, err := repository.InitDB(&cfg.Storage)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close(context.Background())
	log.Printf("DB connected")

	service := handlers.New(db, rdb)
	service.Register(r, cfg)

	port := cfg.Listen.Port
	log.Printf("Listening on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
