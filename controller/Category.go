package controller

import "net/http"


func ViewCategory(w http.ResponseWriter, r *http.Request){

	w.Write([]byte("here is your category"))
}

func CreateCategory(w http.ResponseWriter, r *http.Request){

	w.Write([]byte("create category\n"))
}

func UpdateCategory(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("update category"))
}

func DeleteCategory(w http.ResponseWriter, r *http.Request){

	w.Write([]byte("destroy category"))
}
