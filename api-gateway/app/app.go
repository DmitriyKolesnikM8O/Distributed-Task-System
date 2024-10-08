package app

import (
	"context"
	"log"
	"net/http"

	"github.com/KolesnikM8O/distributed-task-system/api-gateway/redis"
	"github.com/KolesnikM8O/distributed-task-system/api-gateway/repository"
	"github.com/KolesnikM8O/distributed-task-system/api-gateway/service/handlers"
	"github.com/gorilla/mux"
)

func Start(r *mux.Router) {
	log.Printf("Start app")

	log.Printf("Init redis")
	redis.InitRedis()

	db, err := repository.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close(context.Background())
	log.Printf("DB connected")

	r.HandleFunc("/task", func(w http.ResponseWriter, r *http.Request) {
		handlers.CreateTaskHandler(w, r, db)
	}).Methods("POST")

	r.HandleFunc("/task/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetTaskHandler(w, r, db)
	}).Methods("GET")

	r.HandleFunc("/task/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.UpdateTaskHandler(w, r, db)
	}).Methods("PUT")

	r.HandleFunc("/task/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.DeleteTaskHandler(w, r, db)
	}).Methods("DELETE")

	log.Printf("Listening on :8080")
	http.ListenAndServe(":8080", r)
}
