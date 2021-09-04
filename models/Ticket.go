package models

import (
	"fmt"
	"ticket/config"

	"gorm.io/gorm"
)

type Ticket struct {
	gorm.Model
	Subject     string
	Description string
	Priority    string
	CategoryID  int
	UserID      uint
	User        User
	Status      int
	Note        string
	Reply       []Reply
}
type Reply struct {
	gorm.Model
	TicketID uint
	Ticket   Ticket

	UserID  uint
	User    User
	Message string
}

func ViewTicket(ID int) Ticket {
	var ticket Ticket
	err := config.DB.Preload("User").Preload("Reply").Find(&ticket, ID).Error
	if err != nil {
		fmt.Println(err)
	}

	return ticket
}

func (t *Ticket) UpdateTicket(newTicket *Ticket) *Ticket {

	t.Subject = newTicket.Subject
	t.Description = newTicket.Description
	t.CategoryID = newTicket.CategoryID
	config.DB.Save(&t)
	return t
}
func (t *Ticket) CreateTicket() *Ticket {
	ticket := &Ticket{
		Subject:     t.Subject,
		Description: t.Description,
		CategoryID:  t.CategoryID,
		User:        t.User,
	}
	config.DB.Create(ticket)
	return ticket
}

func GetAllTickets() []Ticket {
	var ticket []Ticket
	config.DB.Find(&ticket)
	return ticket
}
func DeleteTicket() Ticket {
	return Ticket{}
}
