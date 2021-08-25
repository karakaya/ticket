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
	r := mux.NewRouter().StrictSlash(true)



	config.DB.AutoMigrate(&models.User{},&models.Ticket{},&models.Category{})


	route.Ticket(r)
	route.User(r)
	route.Category(r)
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}

