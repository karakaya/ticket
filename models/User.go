package models

import (
	"log"
	"ticket/config"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	IsAdmin  bool `gorm:"default:false"`
}

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.Password = HashPassword(u.Password)
	return

}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	u.Password = HashPassword(u.Password)
	return
}

func (u *User) CreateUser() *User {
	config.DB.Create(&u)
	return u
}

func GetUser(id int) *User {
	var user User
	result := config.DB.Omit("Password").Find(&user, id)
	if result.RowsAffected < 1 {
		log.Println("record not found")
		return nil
	}
	return &user
}

func GetAllUsers() []User {
	var users []User
	config.DB.Find(&users)
	return users
}
func (u *User) UpdateUser(user *User) *User {
	u.Name = user.Name
	u.Email = user.Email
	u.Password = user.Password
	config.DB.Save(&u)

	return u
}

func (u *User) DeleteUser() *User {
	config.DB.Delete(&u)
	return u
}
