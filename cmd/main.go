package main

import (
	"fmt"
	"log"
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

	err = rabbit.ConnRabbitMQ()
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

	fmt.Printf("Starting server at port :%s\n", port)
	if err := http.ListenAndServe(":"+port, route); err != nil {
		log.Fatal(err)
	}

}
