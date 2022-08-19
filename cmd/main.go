package main

import (
	"net/http"
	"os"
	"ticket/internals/ticket"

	"ticket/internals/db"

	"github.com/gorilla/mux"
)

const listenPort = "80"

func main() {

	route := mux.NewRouter()

	dbc, err := db.GetMongoClient()
	if err != nil {
		panic(err)
	}
	repo := ticket.NewRepository(dbc)
	service := ticket.NewService(repo)

	ticket.RegisterHandlers(route, service)

	port := listenPort
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	if err := http.ListenAndServe(":"+port, route); err != nil {
		panic(err)
	}

}
