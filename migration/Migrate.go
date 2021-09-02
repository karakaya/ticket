package migration

import (
	"log"
	"ticket/config"
	"ticket/models"
)

func Migrate() {
	log.Println("migrating")
	config.DB.AutoMigrate(&models.User{}, &models.Ticket{}, &models.Category{})
}

func InitAdmin() {
	name := "admin"
	password := "password"
	user := &models.User{
		Name:     name,
		Password: models.HashPassword(password),
		IsAdmin:  true,
	}
	log.Printf("init user; name: %s, password: %s \n", name, password)

	config.DB.Create(&user)

}
