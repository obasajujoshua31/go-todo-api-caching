package config

import (
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"os"
	"strconv"
)

type AppConfig struct {
	DBName     string
	DBPassword string
	DBUser     string
	DBHost     string
	Port       int
	DBPort     int
}

func GetAppConfig() (AppConfig, error) {
	port := os.Getenv("PORT")

	portInt, err := strconv.Atoi(port)
	if err != nil {
		return AppConfig{}, err
	}

	dbPort := os.Getenv("DBPORT")

	dbPortInt, err := strconv.Atoi(dbPort)
	if err != nil {
		return AppConfig{}, err
	}

	err = godotenv.Load()
	if err != nil {
		return AppConfig{}, nil
	}

	appConfig := AppConfig{
		DBName:     os.Getenv("DBNAME"),
		DBPassword: os.Getenv("DBPASSWORD"),
		DBUser:     os.Getenv("DBUSER"),
		DBHost:     os.Getenv("DBHOST"),
		Port:       portInt,
		DBPort:     dbPortInt,
	}

	return appConfig, nil
}
