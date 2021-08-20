package route

import (
	"github.com/gorilla/mux"
	"ticket/controller"
)

func Ticket(r *mux.Router){
	t:= r.PathPrefix("/ticket").Subrouter()
	t.HandleFunc("/",controller.CreateTicket).Methods("POST")
	t.HandleFunc("/",controller.GetAllTickets).Methods("GET")
	t.HandleFunc("/{id}",controller.ViewTicket).Methods("GET")
	t.HandleFunc("/{id}",controller.UpdateTicket).Methods("PATCH")
	t.HandleFunc("/{id}",controller.DeleteTicket).Methods("DELETE")
}


