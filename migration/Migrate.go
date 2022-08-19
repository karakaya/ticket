package migration

import (
	"log"
	"ticket/config"
	"ticket/internals/ticket"
)

func Migrate() {
	log.Println("migrating")
	// config.DB.AutoMigrate(&ticket.User{}, &models.User{}, &models.Reply{}, &models.Ticket{})
}

func InitAdmin() {
	name := "admin"
	password := "password"
	user := &ticket.User{
		Name: name,
		// Password: models.HashPassword(password),
		Password: password,
	}

	config.DB.Create(&user)

}
