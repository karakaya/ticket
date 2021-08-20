package route

import (
	"github.com/gorilla/mux"
	"ticket/controller"
)

func User(r *mux.Router){
	u:= r.PathPrefix("/user").Subrouter()
	u.HandleFunc("/",controller.CreateUser).Methods("POST")
	u.HandleFunc("/",controller.GetAllUsers).Methods("GET")
	u.HandleFunc("/{id}",controller.ViewUser).Methods("GET")
	u.HandleFunc("/{id}",controller.UpdateUser).Methods("PATCH")
	u.HandleFunc("/{id}",controller.DeleteUser).Methods("DELETE")
}


