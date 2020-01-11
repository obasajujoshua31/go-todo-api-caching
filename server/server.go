package server

import (
	"fmt"
	"go-todo-api-caching/config"
	"go-todo-api-caching/server/services"
	"net/http"
)

const driver = "postgres"

func StartServer() error {
	appConfig, err := config.GetAppConfig()

	if err != nil {
		return err
	}

	connString := services.GetConnectionString(appConfig, driver)
	r := InitializeRoutes(connString)

	err = http.ListenAndServe(fmt.Sprintf(":%d", appConfig.Port), r)

	if err != nil {
		return err
	}

	return nil
}
