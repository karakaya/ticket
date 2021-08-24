package route

import (
	"github.com/gorilla/mux"
	"ticket/controller"
)

func Category(r *mux.Router){
	c:= r.PathPrefix("/category").Subrouter()
	c.HandleFunc("/",controller.CreateCategory).Methods("POST")
	c.HandleFunc("/",controller.GetAllCategories).Methods("GET")
	c.HandleFunc("/{id}",controller.ViewCategory).Methods("GET")
	c.HandleFunc("/{id}",controller.UpdateCategory).Methods("PATCH")
	c.HandleFunc("/{id}",controller.DeleteCategory).Methods("DELETE")


}


