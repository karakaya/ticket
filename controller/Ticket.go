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
	res, _ := json.Marshal(models.ViewTicket(id))
	w.Write(res)
}

func CreateTicket(w http.ResponseWriter, r *http.Request) {
	if len(r.FormValue("subject")) == 0 || len(r.FormValue("description")) == 0 || len(r.FormValue("category")) == 0 {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("some fields are empty"))
		return
	}
	var user models.User
	config.DB.Find(&user, "name = ?", GetAuthUsername(r))
	category := models.GetCategory(ConvertInt(r.FormValue("category")))
	if category.Title == "" {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("invalid category"))
		return
	}

	ticket := &models.Ticket{
		Subject:     r.FormValue("subject"),
		Description: r.FormValue("description"),
		CategoryID:  int(category.ID),
		UserID:      int(user.ID),
	}

	res, _ := json.Marshal(ticket.CreateTicket())
	w.Write(res)
}

func GetAllTickets(w http.ResponseWriter, r *http.Request) {
	res, _ := json.Marshal(models.GetAllTickets())
	w.Write(res)
}

func UpdateTicket(w http.ResponseWriter, r *http.Request) {
	if len(r.FormValue("subject")) == 0 || len(r.FormValue("description")) == 0 || len(r.FormValue("category")) == 0 {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("some fields are empty"))
		return
	}

	var user models.User
	config.DB.Find(&user, "name = ?", GetAuthUsername(r))
	category := models.GetCategory(ConvertInt(r.FormValue("category")))
	if category.Title == "" {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("invalid category"))
		return
	}

	id := ConvertInt(mux.Vars(r)["id"])
	var curTicket models.Ticket
	config.DB.Find(&curTicket, id)
	if uint(curTicket.UserID) != user.ID {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	newTicket := &models.Ticket{
		Subject:     r.FormValue("subject"),
		Description: r.FormValue("description"),
		CategoryID:  int(category.ID),
		UserID:      int(user.ID),
	}

	res, _ := json.Marshal(curTicket.UpdateTicket(newTicket))
	w.Write(res)
}

func DeleteTicket(w http.ResponseWriter, r *http.Request) {

	var ticket models.Ticket
	id := ConvertInt(mux.Vars(r)["id"])
	config.DB.Find(&ticket, id)

	var user models.User
	config.DB.Find(&user, "name = ?", GetAuthUsername(r))
	if ticket.Subject == "" {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("ticket not found"))
		return
	}
	if ticket.UserID != int(user.ID) {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	config.DB.Delete(&ticket)
	w.Write([]byte("ticket deleted"))
}
