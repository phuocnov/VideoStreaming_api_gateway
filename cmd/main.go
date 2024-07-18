package main

import (
	"APIGateway/pkg/database"
	"APIGateway/pkg/handler"
	"APIGateway/redis"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	redis.Init()
	database.Init()
	r := mux.NewRouter()
	r.HandleFunc("/todos", handler.GetAllTodo).Methods("GET")
	r.HandleFunc("/todos/{id}", handler.GetTodoById).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}
