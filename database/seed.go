package database

import (
	"github.com/jinzhu/gorm"
	"go-todo-api-caching/server/services"
)

var todos = []services.Todo{
	{
		Title:       "I want to eat",
		IsCompleted: false,
	},
	{
		Title:       "I want to wash",
		IsCompleted: true,
	},
	{
		Title: "I want to pray",
	},
}

func SeedData(db *gorm.DB) {
	for _, todo := range todos {
		db.Create(&todo)
	}
}
