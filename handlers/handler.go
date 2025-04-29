package handlers

import (
	"github.com/matheusvcouto/crud-go/repository"
)

type Handler struct {
	userRepo *repository.UserRepository
}

func NewHandler(userRepo *repository.UserRepository) *Handler {
	return &Handler{
		userRepo: userRepo,
	}
}
