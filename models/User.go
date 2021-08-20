package models

import "gorm.io/gorm"

type User struct{
	gorm.Model
	Name string
	Email string
	Password string
	Tickets []Ticket `gorm:"many2many:user_tickets"`
	IsAdmin bool     `gorm:"default:false"`
}
