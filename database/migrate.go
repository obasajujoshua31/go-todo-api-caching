package database

import (
	"github.com/jinzhu/gorm"
	"go-todo-api-caching/server/services"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&services.Todo{})
}
