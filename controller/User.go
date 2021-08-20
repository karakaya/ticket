package controller

import (
	"net/http"
)



func ViewUser(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("here is your user"))
}

func CreateUser(w http.ResponseWriter, r *http.Request){

	w.Write([]byte("create user\n"))
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("all users\n"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request){

	w.Write([]byte("update user"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request){

	w.Write([]byte("destroy user"))
}
