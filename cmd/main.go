package main

import (
	"net/http"
	"os"

	"github.com/karakaya/ticket/internals/rabbit"
	"github.com/karakaya/ticket/internals/ticket"

	"github.com/karakaya/ticket/internals/db"

	"github.com/gorilla/mux"
)

const listenPort = "80"

func main() {

	route := mux.NewRouter()

	dbc, err := db.GetMongoClient()
	if err != nil {
		panic(err)
	}

	rabbitmq := new(rabbit.RabbitMQ)
	rabbitmq.Connect()

	repo := ticket.NewRepository(dbc)
	service := ticket.NewService(repo, rabbitmq)

	ticket.RegisterHandlers(route, service)

	port := listenPort
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	if err := http.ListenAndServe(":"+port, route); err != nil {
		panic(err)
	}

}
