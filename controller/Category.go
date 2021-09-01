package controller

import (
	"encoding/json"
	"net/http"
	"ticket/models"

	"github.com/gorilla/mux"
)

func ViewCategory(w http.ResponseWriter, r *http.Request) {
	id, err := mux.Vars(r)["id"]
	if !err {
		panic(err)
	}
	category := models.GetCategory(id)
	res, _ := json.Marshal(category)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	category := models.Category{
		Title: r.FormValue("title"),
	}
	category.CreateCategory()

}

func GetAllCategories(w http.ResponseWriter, r *http.Request) {
	categories := models.GetAllCategories()
	allCategories, _ := json.Marshal(categories)
	w.WriteHeader(200)
	w.Write(allCategories)
}
func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	category := models.Category{Title: r.FormValue("title")}
	id := mux.Vars(r)["id"]
	update := category.UpdateCategory(ConvertInt(id))
	res, _ := json.Marshal(update)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	delete := models.DeleteCategory(ConvertInt(mux.Vars(r)["id"]))
	if !delete {
		w.Write([]byte("record not found or failed to delete"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("record deleted successfully"))
}
