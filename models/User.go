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
func init(){
	db.AutoMigrate(&User{})
}
func (u *User) CreateUser() *User{
	return nil
}

func GetUser(ID int) *User{
	return nil
}

func GetAllUsers() []User{
	return nil
}

func DeleteUser(ID int) User {
	return User{}
}
