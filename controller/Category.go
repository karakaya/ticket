package controller

import (
	"encoding/json"
	"net/http"
	"ticket/config"
	"ticket/models"

	"github.com/gorilla/mux"
)

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	if len(r.FormValue("title")) == 0 {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("title field is empty"))
	}
	category := models.Category{
		Title: r.FormValue("title"),
	}
	res, _ := json.Marshal(category.CreateCategory())
	w.Write(res)
}
func ViewCategory(w http.ResponseWriter, r *http.Request) {
	id := ConvertInt(mux.Vars(r)["id"])
	res, _ := json.Marshal(models.GetCategory(id))
	w.Write(res)
}

func GetAllCategories(w http.ResponseWriter, r *http.Request) {
	allCategories, _ := json.Marshal(models.GetAllCategories())
	w.Write(allCategories)
}

func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	var category models.Category
	id := mux.Vars(r)["id"]
	if id == "" {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("category field is empty"))
		return
	}
	config.DB.Find(&category, id)
	newCategory := &models.Category{Title: r.FormValue("title")}
	res, _ := json.Marshal(category.UpdateCategory(newCategory))

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
