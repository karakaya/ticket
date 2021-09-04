package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	godotenv.Load(".env")
	dbhost := goDotEnvVariable("DB_HOST")
	dbport := goDotEnvVariable("DB_PORT")
	dbname := goDotEnvVariable("DB_NAME")
	dbuser := goDotEnvVariable("DB_USER")
	dbpassword := goDotEnvVariable("DB_PASSWORD")
	fmt.Println(dbhost)
	db, err := gorm.Open(mysql.Open(dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}
	DB = db
}

func goDotEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("err loading .env file")
	}
	return os.Getenv(key)
}
