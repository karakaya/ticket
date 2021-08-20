package models

import (
	"gorm.io/gorm"
	"ticket/config"


)
var db *gorm.DB
type Category struct {
	gorm.Model
	Title string
}

func init(){
	db = config.DB
	db.AutoMigrate(&Category{})
}

func (c *Category) CreateCategory() *Category{
	return nil
}

func GetCategory(ID int) *Category {
	return nil
}

func GetAllCategories() []Category{
	return nil
}

func DeleteCategory(ID int) Category{
	return Category{}
}
