package controller

import "net/http"


func ViewTicket(w http.ResponseWriter, r *http.Request){

	w.Write([]byte("here is your ticket"))
}

func CreateTicket(w http.ResponseWriter, r *http.Request){

	w.Write([]byte("create ticket\n"))
}

func UpdateTicket(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("update ticket"))
}

func DeleteTicket(w http.ResponseWriter, r *http.Request){

	w.Write([]byte("destroy ticket"))
}
