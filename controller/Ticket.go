package controller

import (
	"encoding/json"
	"net/http"
	"ticket/config"
	"ticket/models"

	"github.com/gorilla/mux"
)

func ViewTicket(w http.ResponseWriter, r *http.Request) {
	id := ConvertInt(mux.Vars(r)["id"])
	ticket := models.ViewTicket(id)
	res, _ := json.Marshal(ticket)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func CreateTicket(w http.ResponseWriter, r *http.Request) {
	ticket := &models.Ticket{
		Subject:     r.FormValue("subject"),
		Description: r.FormValue("description"),
		CategoryID:  ConvertInt(r.FormValue("category")),
		UserID:      ConvertInt(r.FormValue("user")),
	}
	ticRes := ticket.CreateTicket()
	res, _ := json.Marshal(ticRes)
	w.Write(res)
}

func GetAllTickets(w http.ResponseWriter, r *http.Request) {
	tickets := models.GetAllTickets()
	res, _ := json.Marshal(tickets)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateTicket(w http.ResponseWriter, r *http.Request) {
	ticket := &models.Ticket{
		Subject:     r.FormValue("subject"),
		Description: r.FormValue("description"),
		CategoryID:  ConvertInt(r.FormValue("category")),
		UserID:      ConvertInt(r.FormValue("user")),
		Status:      ConvertInt(r.FormValue("status")),
		Note:        r.FormValue("note"),
	}
	ticket.UpdateTicket(ConvertInt(mux.Vars(r)["id"]))
	w.Write([]byte("update ticket"))
}

func DeleteTicket(w http.ResponseWriter, r *http.Request) {
	id := ConvertInt(mux.Vars(r)["id"])
	ticket := config.DB.Delete(&models.Ticket{}, id)
	if ticket.RowsAffected == 0 {
		w.Write([]byte("record not found"))
		return
	}
	w.Write([]byte("record deleted"))
}
