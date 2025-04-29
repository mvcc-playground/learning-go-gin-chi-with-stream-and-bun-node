package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/matheusvcouto/crud-go/handlers"
	mymidewares "github.com/matheusvcouto/crud-go/middleware"
	"github.com/matheusvcouto/crud-go/repository"
)

func main() {
	usersRepository := repository.NewUserRepository()

	r := chi.NewRouter()
	h := handlers.NewHandler(usersRepository)

	r.Use(middleware.Logger, middleware.Recoverer, middleware.Recoverer, mymidewares.LoggingIPMiddleware)

	r.Get("/", h.HelloWord)
	r.Get("/hello/{name}", h.HelloName)

	r.Route("/users", func(r chi.Router) {
		r.Get("/", h.GetUsers)
		r.Post("/", h.CreateUser)
	})

	http.ListenAndServe(":8000", r)
}
