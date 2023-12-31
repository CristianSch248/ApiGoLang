package main

import (
	"apiGo/configs"
	"apiGo/handlers"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	err := configs.Load()
	if err != nil {
		panic(err)
	}

	route := chi.NewRouter()

	route.Post("/", handlers.Create)
	route.Put("/{id}", handlers.Update)
	route.Delete("/{id}", handlers.Delete)
	route.Get("/", handlers.List)
	route.Get("/{id}", handlers.Get)

	serverAddr := fmt.Sprintf(":%s", configs.GetServerPort())
	fmt.Printf("Servidor escutando em http://localhost%s\n", serverAddr)

	err = http.ListenAndServe(serverAddr, route)
	if err != nil {
		panic(err)
	}
}
