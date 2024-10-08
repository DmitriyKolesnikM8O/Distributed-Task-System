package main

import (
	"github.com/KolesnikM8O/distributed-task-system/api-gateway/app"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	app.Start(r)

}
