package main

import (
	"log"
	stdhttp "net/http"
	"time"

	todohttp "ecom-tech/internal/http"
	"ecom-tech/internal/todo"
)

func main() {
	store := todo.NewStore()
	service := todo.NewService(store)
	handler := todohttp.NewHandler(service)

	loggedHandler := todohttp.LoggingMiddleware(handler.Router())

	srv := &stdhttp.Server{
		Addr:         ":8080",
		Handler:      loggedHandler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Println("Server started at :8080")
	log.Fatal(srv.ListenAndServe())
}
