package server

import (
	"encoding/json"
	"github.com/go-redis/redis/v7"
	"github.com/gorilla/mux"
	"go-todo-api-caching/extra"
	"go-todo-api-caching/server/services"
	"io/ioutil"
	"net/http"
)

func GetAllTodosHandler(connString string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		key := r.RequestURI

		redisClient := services.NewClient()

		res, err := redisClient.GetRecord(key)

		if res != nil {
			var todos []services.Todo

			err := json.Unmarshal(res, &todos)

			if err != nil {
				WriteErrorResponse(w, extra.UnableToUnmarshalError)
				return
			}
			WriteResponse(w, todos)
			return
		}

		if err == redis.Nil {
			db, err := services.ConnectToDB(connString, driver)

			if err != nil {
				WriteErrorResponse(w, extra.UnableToConnectToDBError)
				return
			}
			defer db.Close()

			todos := db.GetAllTodos()

			b, err := json.Marshal(todos)

			if err != nil {
				WriteErrorResponse(w, extra.UnableToMarshalResponse)
				return
			}

			go redisClient.SetNewRecord(key, b)

			w.WriteHeader(http.StatusOK)
			WriteResponse(w, todos)
			return
		}
		WriteErrorResponse(w, extra.UnableToConnectToRedis)
		return
	}
}

func GetAllCompletedTodosHandler(connString string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		key := r.RequestURI

		rc := services.NewClient()

		res, err := rc.GetRecord(key)

		if res != nil {
			var todos []services.Todo

			err := json.Unmarshal(res, &todos)

			if err != nil {
				WriteErrorResponse(w, extra.UnableToUnmarshalError)
				return
			}
			WriteResponse(w, todos)
			return
		}

		if err == redis.Nil {
			db, err := services.ConnectToDB(connString, driver)

			if err != nil {
				WriteErrorResponse(w, extra.UnableToConnectToDBError)
				return
			}
			defer db.Close()

			todos := db.GetAllCompletedTodos()

			b, err := json.Marshal(todos)

			if err != nil {
				WriteErrorResponse(w, extra.UnableToMarshalResponse)
				return
			}

			go rc.SetNewRecord(key, b)

			w.WriteHeader(http.StatusOK)
			WriteResponse(w, todos)
			return
		}

		WriteErrorResponse(w, extra.UnableToConnectToRedis)
		return

	}
}

func GetOneTodoHandler(connString string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		key := r.RequestURI

		rc := services.NewClient()

		res, err := rc.GetRecord(key)

		if res != nil {
			var todo services.Todo

			err := json.Unmarshal(res, &todo)
			if err != nil {
				WriteErrorResponse(w, extra.UnableToMarshalResponse)
				return
			}

			WriteResponse(w, todo)
			return
		}

		if err == redis.Nil {
			db, err := services.ConnectToDB(connString, driver)

			if err != nil {
				WriteErrorResponse(w, extra.UnableToConnectToDBError)
				return
			}
			defer db.Close()

			vars := mux.Vars(r)

			id := vars["id"]

			todo := db.GetTodoByID(id)

			if todo.Title == "" {
				http.NotFound(w, r)
				return
			}

			b, err := json.Marshal(todo)
			if err != nil {
				WriteErrorResponse(w, extra.UnableToMarshalResponse)
				return
			}

			go rc.SetNewRecord(key, b)

			w.WriteHeader(http.StatusOK)
			WriteResponse(w, todo)
			return
		}

		WriteErrorResponse(w, extra.UnableToConnectToRedis)
		return
	}
}

func GetAllNotCompletedTodosHandler(connString string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key := r.RequestURI

		rc := services.NewClient()

		res, err := rc.GetRecord(key)

		if res != nil {
			var todos []services.Todo

			err := json.Unmarshal(res, &todos)
			if err != nil {
				WriteErrorResponse(w, extra.UnableToUnmarshalError)
				return
			}

			WriteResponse(w, todos)
			return
		}

		if err == redis.Nil {
			db, err := services.ConnectToDB(connString, driver)

			if err != nil {
				WriteErrorResponse(w, extra.UnableToConnectToDBError)
				return
			}
			defer db.Close()

			todos := db.GetAllNotCompletedTodos()

			b, err := json.Marshal(todos)

			if err != nil {
				WriteErrorResponse(w, extra.UnableToMarshalResponse)
				return
			}

			go rc.SetNewRecord(key, b)

			w.WriteHeader(http.StatusOK)
			WriteResponse(w, todos)
			return
		}

		WriteErrorResponse(w, extra.UnableToConnectToRedis)
		return
	}
}

func CreateNewTodoHandler(connString string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		var todo services.Todo

		if err != nil {
			WriteErrorResponse(w, extra.UnableToReadBodyError)
			return
		}

		err = json.Unmarshal(body, &todo)

		if err != nil {
			WriteErrorResponse(w, extra.UnableToUnmarshalError)
			return
		}

		if todo.Title == "" {
			w.WriteHeader(http.StatusBadRequest)

			WriteResponse(w, map[string]string{
				"error": "Validation error",
			})
			return
		}

		db, err := services.ConnectToDB(connString, driver)

		if err != nil {
			WriteErrorResponse(w, extra.UnableToConnectToDBError)
			return
		}
		defer db.Close()

		todo = db.CreateTodo(todo)

		client := services.NewClient()

		go client.ClearCache("/todos")
		w.WriteHeader(http.StatusCreated)
		WriteResponse(w, todo)
	}
}

func CompleteTodoHandler(connString string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		db, err := services.ConnectToDB(connString, driver)

		if err != nil {
			WriteErrorResponse(w, extra.UnableToConnectToDBError)
			return
		}
		defer db.Close()

		vars := mux.Vars(r)

		id := vars["id"]

		todo := db.GetTodoByID(id)

		if todo.Title == "" {
			http.NotFound(w, r)
			return
		}

		todo = db.CompleteTodo(todo)

		rc := services.NewClient()

		go rc.ClearCache("/todos")
		go rc.ClearCache(r.RequestURI)
		WriteResponse(w, todo)
	}
}

func DeleteTodoHandler(connString string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		db, err := services.ConnectToDB(connString, driver)

		if err != nil {
			WriteErrorResponse(w, extra.UnableToConnectToDBError)
			return
		}
		defer db.Close()

		vars := mux.Vars(r)

		id := vars["id"]

		todo := db.GetTodoByID(id)

		if todo.Title == "" {
			http.NotFound(w, r)
			return
		}

		rc := services.NewClient()

		db.DeleteTodo(todo)
		go rc.ClearCache("/todos")
		go rc.ClearCache(r.RequestURI)
		w.WriteHeader(http.StatusNoContent)
	}
}

func WriteResponse(w http.ResponseWriter, data interface{}) {
	b, err := json.MarshalIndent(data, "", " ")

	if err != nil {
		WriteErrorResponse(w, extra.UnableToMarshalResponse)
		return
	}

	_, err = w.Write(b)

	if err != nil {
		WriteErrorResponse(w, extra.UnableToMarshalResponse)
		return
	}
}

func WriteErrorResponse(w http.ResponseWriter, error extra.ServerError) {
	b, _ := json.Marshal(error)

	w.Write(b)
}
