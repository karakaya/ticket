package models

import (
	"ticket/config"

	"gorm.io/gorm"
)

type Ticket struct {
	gorm.Model
	Subject     string
	Description string
	Priority    string
	//CategoryID  Category `gorm:"foreignKey:ID;References:ID;"`
	CategoryID int
	UserID     int `gorm:"foreignKey:id"`
	Status     int
	Note       string
}

func ViewTicket(ID int) *Ticket {
	var ticket Ticket
	config.DB.First(Ticket{}, ID).Scan(&ticket)
	return &ticket
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
		UserID:      t.UserID,
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
