package main

import (
	"net/http"
	"os"
	"ticket/internals/ticket"

	"github.com/gorilla/mux"
)

const listenPort = "80"

func main() {

	route := mux.NewRouter()
	repo := ticket.NewRepository(nil)
	service := ticket.NewService(repo)

	ticket.RegisterHandlers(route, service)

	port := listenPort
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	err := http.ListenAndServe(":"+port, route)
	if err != nil {
		panic(err)
	}

}
