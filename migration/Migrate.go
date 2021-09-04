package migration

import (
	"log"
	"ticket/config"
	"ticket/models"
)

func Migrate() {
	log.Println("migrating")
	config.DB.AutoMigrate(&models.Category{}, &models.User{}, &models.Reply{}, &models.Ticket{})
}

func InitAdmin() {
	name := "admin"
	password := "password"
	user := &models.User{
		Name:     name,
		Password: models.HashPassword(password),
		IsAdmin:  true,
	}

	config.DB.Create(&user)

}
