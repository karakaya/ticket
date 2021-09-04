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

	migration.Migrate()
	//migration.InitAdmin()
	r := mux.NewRouter().StrictSlash(true)
	route.Routes(r)
	r.
		PathPrefix("/static/").
		Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("."+"/static/"))))

	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}

}
