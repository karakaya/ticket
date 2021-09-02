package models

import (
	"log"
	"ticket/config"

	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Title string
}

func (c *Category) CreateCategory() *Category {
	config.DB.Create(&Category{Title: c.Title})
	return c
}

func GetCategory(ID int) *Category {
	var category Category
	config.DB.Find(&category, ID)
	return &category
}

func GetAllCategories() []Category {
	var categories []Category
	config.DB.Find(&categories)
	return categories
}

func (c *Category) UpdateCategory(category *Category) *Category {
	c.Title = category.Title
	config.DB.Save(&c)
	return c
}

func DeleteCategory(ID int) bool {
	check := config.DB.Delete(&Category{}, ID)
	if check.RowsAffected == 0 {
		log.Println("record not found")
		return false
	}
	return true
}
