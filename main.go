package main

import (
	"net/http"
	"ticket/config"
	"ticket/migration"
	"ticket/route"

	"github.com/gorilla/mux"
)

func main() {
	config.ConnectDB()

	//migration.Migrate()
	migration.InitAdmin()
	r := mux.NewRouter().StrictSlash(true)
	route.Routes(r)

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}
