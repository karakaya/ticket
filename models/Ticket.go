package models

import (
	"fmt"
	"gorm.io/gorm"
	"ticket/config"
)

type Ticket struct{
	gorm.Model
	Subject string
	Description string
	Priority string
	//CategoryID Category `gorm:"foreignKey:ID;References:ID;"`
	CategoryID int
	UserID int `gorm:"foreignKey:id"`
	Status int
	Note string
}
func ViewTicket(ID int) *Ticket{
	var ticket Ticket
	config.DB.First(Ticket{},ID).Scan(&ticket)
	return &ticket
}

func (t *Ticket) UpdateTicket(id int) *Ticket{
	var currTicket Ticket
	config.DB.First(&currTicket,id)
	currTicket.Subject = t.Subject
	currTicket.Description = t.Description
	currTicket.CategoryID = t.CategoryID
	currTicket.Note = t.Note
	currTicket.Status = t.Status
	config.DB.Save(&currTicket)
	fmt.Println(currTicket)
	return nil
}
func (t *Ticket) CreateTicket() *Ticket{
	ticket :=&Ticket{
		Subject: t.Subject,
		Description: t.Description,
		CategoryID: t.CategoryID,
		UserID: t.UserID,

	}
	config.DB.Create(ticket)
	return ticket
}

func GetAllTickets() []Ticket{
	var ticket []Ticket
	config.DB.Find(&ticket)
	return ticket
}
func DeleteTicket() Ticket{
	return Ticket{}
}
