package config

import (
	"log"
)

func Migrate(){
	log.Println("migrating")
	/*DB.AutoMigrate(&models.Category{})
	DB.AutoMigrate(&models.Ticket{})
	DB.AutoMigrate(&models.User{})

	 */
}
