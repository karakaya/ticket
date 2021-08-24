package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"ticket/models"
)
func ViewUser(w http.ResponseWriter, r *http.Request){
	id,ok := mux.Vars(r)["id"]
	if !ok || len(id) < 1{
		log.Println("id parameter is missing")
		return
	}
	i := ConvertInt(id)
	user := models.GetUser(i)
	w.WriteHeader(http.StatusOK)
	res, _ := json.Marshal(user)
	w.Write(res)
}

func CreateUser(w http.ResponseWriter, r *http.Request){
	user := models.CreateUser(r)
	res, _ := json.Marshal(user)
	w.Write(res)
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users := models.GetAllUsers()
	allUsers,_ := json.Marshal(users)
	w.Write(allUsers)
}

func UpdateUser(w http.ResponseWriter, r *http.Request){
	id := ConvertInt(mux.Vars(r)["id"])
	user := models.User{
		Name: r.FormValue("name"),
		Email: r.FormValue("email"),
		Password: r.FormValue("password"),
	}
	userRes := user.UpdateUser(id)
	res,_:=json.Marshal(userRes)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteUser(w http.ResponseWriter, r *http.Request){
	id := mux.Vars(r)["id"]
	models.DeleteUser(ConvertInt(id))
	return
}

func ConvertInt(number string) int{
	i, err := strconv.Atoi(number); if err!=nil{
		log.Println("invalid format")
		return 0
	}
	return i
}