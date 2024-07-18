package handler

import (
	"APIGateway/pkg/database"
	"APIGateway/pkg/dto"
	"APIGateway/redis"
	"fmt"
	"net/http"
	"strconv"

	"github.com/goccy/go-json"
	"github.com/gorilla/mux"
)

// Get all todos
func GetAllTodo(writer http.ResponseWriter, request *http.Request) {
	const cache_key = "todos"

	var todos []dto.Todo

	if err := redis.GetTodos(cache_key, &todos); err != nil {
		err := database.DB.Select(&todos, "SELECT id, title, completed FROM todos")
		if err != nil {
			responseWithJson(writer, http.StatusInternalServerError, map[string]string{"message": "Database error"})
			return
		}
		redis.SetTodos(cache_key, todos)
		fmt.Println("Get todos from database")
	}

	responseWithJson(writer, http.StatusOK, todos)
}

// Get todo by id
func GetTodoById(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		responseWithJson(writer, http.StatusBadRequest, map[string]string{"message": "Invalid todo id"})
		return
	}

	var todo dto.Todo
	err = database.DB.Get(&todo, "SELECT id, title, completed FROM todos WHERE id=$1", id)
	if err != nil {
		responseWithJson(writer, http.StatusNotFound, map[string]string{"message": "Todo not found"})
		return
	}

	responseWithJson(writer, http.StatusOK, todo)
}

func responseWithJson(writer http.ResponseWriter, code int, payload interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(code)
	json.NewEncoder(writer).Encode(payload)
}
