package models

import (
	"fmt"
	"gorm.io/gorm"
	"log"
	"net/http"
	"ticket/config"
)

type User struct{
	gorm.Model
	Name string
	Email string
	Password string
	Tickets []Ticket `gorm:"many2many:user_tickets"`
	IsAdmin bool     `gorm:"default:false"`
}


func  CreateUser(r *http.Request) *User{
	user:=&User{
		Name: r.FormValue("name"),
		Email: r.FormValue("email"),
		Password: r.FormValue("password"),
	}
	config.DB.Save(user)
	return user
}


func GetUser(id int) *User {
	var user User
	result := config.DB.First(&user,id)
	if result.RowsAffected < 1{
		log.Println("record not found")
		return nil
	}
	return &user
}

func GetAllUsers() []User{
	var users []User
	config.DB.Find(&users)
	return users
}

func (u *User) UpdateUser(id int) *User{
	var currUser User
	config.DB.First(&User{},id).Scan(&currUser)
	fmt.Println(currUser)
	currUser.Name = u.Name
	currUser.Email= u.Email
	currUser.Password = u.Password
	config.DB.Save(&currUser)
	return u
}
func DeleteUser(id int) bool {
	check := config.DB.Delete(&User{},id)
	if check.RowsAffected == 0{
		log.Println("record not found")
		return false
	}
	log.Println("record found and deleted")
	return  true
}
