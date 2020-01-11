package main

import (
	"go-todo-api-caching/server"
	"log"
)

func main() {
	err := server.StartServer()

	if err != nil {
		log.Fatal(err)
	}

}
