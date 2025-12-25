package main

import (
	"ecom-tech/api"
	todohttp "ecom-tech/internal/http"
	"ecom-tech/internal/todo"
	"log"
	"net/http"
	"time"
)

func main() {
	store := todo.NewStore()
	service := todo.NewService(store)
	router := api.SetupRouter(service)
	routerWithLogging := todohttp.LoggingMiddleware(router)

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      routerWithLogging,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Println("Server started at :8080")
	log.Fatal(srv.ListenAndServe())
}
