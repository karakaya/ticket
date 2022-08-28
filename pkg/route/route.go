package route

import (
	"encoding/json"
	"net/http"

	"github.com/karakaya/ticket/pkg/errors"
	"github.com/karakaya/ticket/pkg/middleware"
	"github.com/karakaya/ticket/pkg/request"
	"github.com/karakaya/ticket/pkg/service"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func RegisterHandlers(r *mux.Router, service service.Service) {
	res := resource{service}
	r.Use(middleware.ContentTypeApplicationJsonMiddleware)
	r.HandleFunc("/ticket", res.create).Methods("POST")
	r.HandleFunc("/ticket", res.getAll).Methods("GET")
	r.HandleFunc("/ticket/{id}", res.find).Methods("GET")
	r.HandleFunc("/ticket/{id}", res.delete).Methods("DELETE")
}

type resource struct {
	service service.Service
}

func (res resource) create(w http.ResponseWriter, r *http.Request) {
	var ticketRequest request.CreateTicketRequest
	err := json.NewDecoder(r.Body).Decode(&ticketRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	uid, err := res.service.CreateTicket(ticketRequest)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	response, _ := json.Marshal(&request.CreateTicketResponse{ID: uid})

	w.Write(response)
	w.WriteHeader(http.StatusCreated)

}

func (res resource) find(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	uid, err := uuid.Parse(id)
	if err != nil {
		errors.JSONHandleError(w, errors.ErrUUDInvalid)
	}

	ticket, err := res.service.GetTicket(uid)
	if err != nil {
		errors.JSONHandleError(w, err)
		return
	}

	json.NewEncoder(w).Encode(ticket)
}

func (res resource) getAll(w http.ResponseWriter, r *http.Request) {
	tickets, err := res.service.GetAllTickets()
	if err != nil {
		errors.JSONHandleError(w, err)
		return
	}
	json.NewEncoder(w).Encode(tickets)
}

func (res resource) delete(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	uid, err := uuid.Parse(id)
	if err != nil {
		errors.JSONHandleError(w, errors.ErrUUDInvalid)
	}

	err = res.service.DeleteTicket(uid)
	if err != nil {
		errors.JSONHandleError(w, err)
	}

}
