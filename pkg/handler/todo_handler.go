package handler

import (
	"APIGateway/pkg/data"
	"net/http"
	"strconv"

	"github.com/goccy/go-json"
	"github.com/gorilla/mux"
)

func GetAllTodo(writer http.ResponseWriter, request *http.Request) {
	todos := data.Todos
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
