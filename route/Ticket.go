package route

import (
	"ticket/controller"
	"ticket/middleware"

	"github.com/gorilla/mux"
)

func Ticket(r *mux.Router) {
	t := r.PathPrefix("/ticket").Subrouter()
	t.HandleFunc("/", controller.CreateTicket).Methods("POST")
	t.HandleFunc("/", controller.GetAllTickets).Methods("GET")
	t.HandleFunc("/{id}", controller.ViewTicket).Methods("GET")
	t.HandleFunc("/{id}", controller.UpdateTicket).Methods("PATCH")
	t.HandleFunc("/{id}", middleware.IsAuth(controller.DeleteTicket)).Methods("DELETE")
}
