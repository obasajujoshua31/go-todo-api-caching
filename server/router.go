package server

import (
	"github.com/gorilla/mux"
	"net/http"
)

func InitializeRoutes(connString string) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/todos", GetAllTodosHandler(connString)).Methods(http.MethodGet)
	r.HandleFunc("/completed-todos", GetAllCompletedTodosHandler(connString)).Methods(http.MethodGet)
	r.HandleFunc("/not-completed-todos", GetAllNotCompletedTodosHandler(connString)).Methods(http.MethodGet)
	r.HandleFunc("/todos/{id}", GetOneTodoHandler(connString)).Methods(http.MethodGet)
	r.HandleFunc("/todos/{id}", CompleteTodoHandler(connString)).Methods(http.MethodPut)
	r.HandleFunc("/todos/{id}", DeleteTodoHandler(connString)).Methods(http.MethodDelete)
	r.HandleFunc("/todos", CreateNewTodoHandler(connString)).Methods(http.MethodPost)
	return r
}
