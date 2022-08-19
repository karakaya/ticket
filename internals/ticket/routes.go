package ticket

import (
	"encoding/json"
	"net/http"
	"ticket/internals/middleware"

	"github.com/gorilla/mux"
)

func RegisterHandlers(r *mux.Router, service Service) {
	res := resource{service}
	r.Use(middleware.ContentTypeApplicationJsonMiddleware)
	r.HandleFunc("/ticket", res.create).Methods("POST")
}

type resource struct {
	service Service
}

func (res resource) create(w http.ResponseWriter, r *http.Request) {
	var ticket Ticket
	err := json.NewDecoder(r.Body).Decode(&ticket)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	uid, err := res.service.CreateTicket(ticket)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	response, _ := json.Marshal(&CreateTicketResponse{ID: uid})

	w.Write(response)
	w.WriteHeader(http.StatusCreated)

}
