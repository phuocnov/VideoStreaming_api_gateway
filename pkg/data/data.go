package data

import "APIGateway/pkg/dto"

var Todos []dto.Todo

func init() {
	Todos = []dto.Todo{
		{ID: 1, Title: "Create a new API", Completed: false},
		{ID: 2, Title: "Update the API", Completed: false},
		{ID: 3, Title: "Delete the API", Completed: false},
	}
}
