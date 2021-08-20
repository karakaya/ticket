package models

import "gorm.io/gorm"

type Ticket struct{
	gorm.Model
	Subject string
	Description string
	Priority string
	Category Category `gorm:"foreignKey:ID;References:ID;"`
	Status int
	Note string
}

func init(){
	db.AutoMigrate(&Ticket{})
}
func (t *Ticket) CreateTicket() *Ticket{
	return nil
}

func GetTicket(ID int) *Ticket{
	return nil
}

func GetAllTickets() []Ticket{
	return nil
}
func DeleteTicket(ID int) Ticket{
	return Ticket{}
}
