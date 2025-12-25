package api

import (
	todohttp "ecom-tech/internal/http"
	"ecom-tech/internal/todo"
	"net/http"
)

func SetupRouter(service *todo.Service) http.Handler {
	mux := http.NewServeMux()

	handlers := todohttp.NewTodoHandlers(service)

	//все маршруты API
	mux.HandleFunc("GET /todos", handlers.GetAll)
	mux.HandleFunc("POST /todos", handlers.Create)
	mux.HandleFunc("GET /todos/{id}", handlers.GetByID)
	mux.HandleFunc("PUT /todos/{id}", handlers.Update)
	mux.HandleFunc("DELETE /todos/{id}", handlers.Delete)

	// Возвращение роутера для main.go
	return mux
}
