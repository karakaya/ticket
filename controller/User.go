package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"ticket/config"

	"ticket/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

func ViewUser(w http.ResponseWriter, r *http.Request) {
	id := ConvertInt(mux.Vars(r)["id"])
	if id < 1 {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("missing id parameter"))
		return
	}

	user := models.GetUser(id)
	w.WriteHeader(http.StatusOK)
	res, _ := json.Marshal(user)
	w.Write(res)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {

	if len(r.FormValue("password")) == 0 || len(r.FormValue("name")) == 0 || len(r.FormValue("email")) == 0 {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("some field are empty"))
		return
	}

	user := &models.User{
		Name:     r.FormValue("name"),
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}
	res, _ := json.Marshal(user.CreateUser())
	w.Write(res)
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users := models.GetAllUsers()
	allUsers, _ := json.Marshal(users)
	w.Write(allUsers)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	username := GetAuthUsername(r)
	var curUser models.User

	config.DB.Find(&curUser, "name = ?", username)

	if len(r.FormValue("password")) == 0 || len(r.FormValue("name")) == 0 || len(r.FormValue("email")) == 0 {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("some field are empty"))
		return
	}

	newUser := &models.User{
		Name:     r.FormValue("name"),
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}
	res, _ := json.Marshal(curUser.UpdateUser(newUser))
	w.Write(res)

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	username := GetAuthUsername(r)
	var user models.User
	config.DB.First(&user, "name = ?", username)
	user.DeleteUser()
	w.WriteHeader(http.StatusOK)
}

func ConvertInt(number string) int {
	i, err := strconv.Atoi(number)
	if err != nil {
		log.Println("invalid format")
		return 0
	}
	return i
}

func GetAuthUsername(r *http.Request) string {
	tokenString := r.Header.Get("Authorization")
	claims, _ := VerifyToken(tokenString)
	return claims.(jwt.MapClaims)["username"].(string)

}
