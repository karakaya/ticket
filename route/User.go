package route

import (
	"ticket/controller"
	"ticket/middleware"

	"github.com/gorilla/mux"
)

func User(r *mux.Router) {
	u := r.PathPrefix("/user").Subrouter()
	u.HandleFunc("/", controller.CreateUser).Methods("POST")
	u.HandleFunc("/", controller.GetAllUsers).Methods("GET")
	u.HandleFunc("/{id}", controller.ViewUser).Methods("GET")
	u.HandleFunc("/{id}", middleware.IsAuth(controller.UpdateUser)).Methods("PUT")
	u.HandleFunc("/{id}", middleware.IsAuth(controller.DeleteUser)).Methods("DELETE")

	r.HandleFunc("/login", controller.Login).Methods("POST")
}
