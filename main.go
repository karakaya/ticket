package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"ticket/config"
	"ticket/models"
	"ticket/route"
)

func main() {
	config.ConnectDB()
	config.DB.AutoMigrate(&models.User{},&models.Ticket{},&models.Category{})
	r := mux.NewRouter().StrictSlash(true)


	route.Ticket(r)
	route.User(r)
	route.Category(r)
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}

