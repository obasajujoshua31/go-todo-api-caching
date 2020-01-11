package services

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"go-todo-api-caching/config"
)

type DB struct {
	*gorm.DB
}

func ConnectToDB(connString string, driver string) (DB, error) {
	db, err := gorm.Open(driver, connString)

	if err != nil {
		return DB{}, err
	}

	return DB{db}, nil
}

func GetConnectionString(appConfig config.AppConfig, driver string) string {
	return fmt.Sprintf("%s://%s:%s@%s/%s?sslmode=disable",
		driver, appConfig.DBUser, appConfig.DBPassword, appConfig.DBHost, appConfig.DBName)
}
