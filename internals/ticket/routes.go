package ticket

import (
	"encoding/json"
	"net/http"
	"ticket/internals/middleware"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func RegisterHandlers(r *mux.Router, service Service) {
	res := resource{service}
	r.Use(middleware.ContentTypeApplicationJsonMiddleware)
	r.HandleFunc("/ticket", res.create).Methods("POST")
	r.HandleFunc("/ticket/{id}", res.find).Methods("GET")
}

type resource struct {
	service Service
}

func (res resource) create(w http.ResponseWriter, r *http.Request) {
	var ticketRequest CreateTicketRequest
	err := json.NewDecoder(r.Body).Decode(&ticketRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	uid, err := res.service.CreateTicket(ticketRequest)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	response, _ := json.Marshal(&CreateTicketResponse{ID: uid})

	w.Write(response)
	w.WriteHeader(http.StatusCreated)

}

func (res resource) find(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	uid, err := uuid.Parse(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	ticket, err := res.service.GetTicket(uid)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(ticket)
}
