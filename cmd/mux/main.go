package main

import (
	"net/http"

	"github.com/matheusvcouto/crud-go/handlers"
	"github.com/matheusvcouto/crud-go/repository"
)

func main() {
	usersRepository := repository.NewUserRepository()

	mux := http.NewServeMux()
	h := handlers.NewHandler(usersRepository)

	mux.HandleFunc("GET /users", h.GetUsers)
	mux.HandleFunc("POST /users", h.CreateUser)

	http.ListenAndServe(":8000", mux)
}
