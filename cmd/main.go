package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/karakaya/ticket/pkg/db"
	"github.com/karakaya/ticket/pkg/rabbit"
	"github.com/karakaya/ticket/pkg/repository"
	"github.com/karakaya/ticket/pkg/route"
	"github.com/karakaya/ticket/pkg/service"

	"github.com/gorilla/mux"
)

const listenPort = "80"

func main() {

	muxRouter := mux.NewRouter()

	dbc, err := db.GetMongoClient()
	if err != nil {
		panic(err)
	}

	err = rabbit.ConnRabbitMQ()
	if err != nil {
		panic(err)
	}

	repo := repository.NewRepository(dbc)
	service := service.NewService(repo)

	route.RegisterHandlers(muxRouter, service)

	port := listenPort
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	fmt.Printf("Starting server at port :%s\n", port)
	if err := http.ListenAndServe(":"+port, muxRouter); err != nil {
		log.Fatal(err)
	}

}
