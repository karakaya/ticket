package route

import (
	"github.com/gorilla/mux"
	"ticket/controller"
)

func User(r *mux.Router){
	t:= r.PathPrefix("/user").Subrouter()
	t.HandleFunc("/",controller.CreateUser).Methods("POST")
	t.HandleFunc("/{id}",controller.ViewUser).Methods("GET")
	t.HandleFunc("/{id}",controller.UpdateUser).Methods("PATCH")
	t.HandleFunc("/{id}",controller.DeleteUser).Methods("DELETE")
}


