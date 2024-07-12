package handler

import (
	"APIGateway/pkg/data"
	"APIGateway/pkg/dto"
	"APIGateway/redis"
	"fmt"
	"net/http"
	"strconv"

	"github.com/goccy/go-json"
	"github.com/gorilla/mux"
)

func GetAllTodo(writer http.ResponseWriter, request *http.Request) {
	const cache_key = "todos"

	var todos []dto.Todo

	if err := redis.GetTodos(cache_key, &todos); err != nil {
		todos = data.Todos
		redis.SetTodos(cache_key, todos)
		fmt.Println("Get todos from data")
	}

	responseWithJson(writer, http.StatusOK, todos)
}

func GetTodoById(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		responseWithJson(writer, http.StatusBadRequest, map[string]string{"message": "Invalid todo id"})
		return
	}

	for _, todo := range data.Todos {
		if todo.ID == id {
			responseWithJson(writer, http.StatusOK, todo)
			return
		}
	}

	responseWithJson(writer, http.StatusNotFound, map[string]string{"message": "Todo not found"})
}

func responseWithJson(writer http.ResponseWriter, code int, payload interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(code)
	json.NewEncoder(writer).Encode(payload)
}
