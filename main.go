package main

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"net/http"
	"ticket/config"
	"ticket/route"
)

func main() {
	config.ConnectDB()
	r := mux.NewRouter().StrictSlash(true)
	config.DB.Session(&gorm.Session{SkipDefaultTransaction: true})
	route.Ticket(r)
	route.User(r)
	route.Category(r)
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}

