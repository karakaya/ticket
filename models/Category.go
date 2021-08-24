package models

import (
	"gorm.io/gorm"
	"log"
	"ticket/config"
)
var db *gorm.DB
type Category struct {
	gorm.Model
	Title string
}

func (c *Category) CreateCategory() *Category{
	check := config.DB.Create(&Category{Title: c.Title})
	 if check.RowsAffected >= 1{
	 	log.Println("category saved")
	 }
	return c
}

func GetCategory(ID string) *Category {
	var category Category
	config.DB.Find(&category,ID)
	return &category
}

func GetAllCategories() []Category{
	var categories []Category
	config.DB.Find(&categories)
	return categories
}

func (c *Category) UpdateCategory(id int) *Category{
	var currCategory Category
	config.DB.First(&currCategory,id)
	currCategory.Title = c.Title
	config.DB.Save(&currCategory)
	return c
}

func DeleteCategory(ID int) bool{
	check := config.DB.Delete(&Category{},ID)
	if check.RowsAffected == 0{
		log.Println("record not found")
		return false
	}
	return true
}
