package main

import (
	"net/http"
	"ticket/config"
	"ticket/route"

	"github.com/gorilla/mux"
)

func main() {
	config.ConnectDB()
	//	config.DB.AutoMigrate(&models.User{}, &models.Ticket{}, &models.Category{})
	r := mux.NewRouter().StrictSlash(true)
	route.Authroutes(r)
	route.Ticket(r)
	route.User(r)
	route.Category(r)
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}
