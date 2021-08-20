package config

import "gorm.io/gorm"
import "gorm.io/driver/mysql"
var DB *gorm.DB
func ConnectDB(){
	db,err := gorm.Open(mysql.Open("root:password@tcp(127.0.0.1)/ticket?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil{
		panic(err)
	}
	DB = db
}
