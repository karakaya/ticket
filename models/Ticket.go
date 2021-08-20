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
