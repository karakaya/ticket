package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"ticket/config"
	"ticket/route"
)

func main() {
	config.ConnectDB()

	r := mux.NewRouter()
	r.StrictSlash(true)
	route.Ticket(r)
	route.User(r)
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}

