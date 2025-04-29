package repository

import (
	"sync"

	"github.com/google/uuid"

	"github.com/matheusvcouto/crud-go/models"
)

type UserRepository struct {
	mu    sync.Mutex
	users map[string]models.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		users: make(map[string]models.User),
	}
}

func (r *UserRepository) GetAllUsers() []models.User {
	r.mu.Lock()
	defer r.mu.Unlock()

	var userlist []models.User
	for _, u := range r.users {
		userlist = append(userlist, u)
	}
	return userlist
}

func (r *UserRepository) AddUser(user models.User) (models.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	uuid := uuid.New().String()
	user.Id = uuid
	r.users[user.Id] = user

	return user, nil
}
