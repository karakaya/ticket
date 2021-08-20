package route

import (
	"github.com/gorilla/mux"
	"ticket/controller"
)

func Category(r *mux.Router){
	t:= r.PathPrefix("/ticket").Subrouter()
	t.HandleFunc("/",controller.CreateCategory).Methods("POST")
	t.HandleFunc("/{id}",controller.ViewCategory).Methods("GET")
	t.HandleFunc("/{id}",controller.UpdateCategory).Methods("PATCH")
	t.HandleFunc("/{id}",controller.DeleteCategory).Methods("DELETE")
}


