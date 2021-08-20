package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"ticket/database"
	"ticket/models"
	"ticket/route"
)

func main() {
	database.ConnectDB()
	db := database.DB
	migrate := os.Getenv("migrate")
	if migrate == "true"{
		db.AutoMigrate(&models.Category{})
		db.AutoMigrate(&models.User{})
		db.AutoMigrate(&models.Ticket{})
	}
	r := mux.NewRouter()
	r.StrictSlash(true)
	route.Ticket(r)
	route.User(r)
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}

