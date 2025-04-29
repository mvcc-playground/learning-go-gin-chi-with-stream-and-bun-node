package main

import (
	"encoding/json"
	"net/http"
	"sync"
)

type User struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
type Users = []User

var (
	users = make(map[string]User)
	mu    sync.Mutex
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("GET /users", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		mu.Lock()
		var userlist Users
		for _, u := range users {
			userlist = append(userlist, u)
		}
		mu.Unlock()

		json.NewEncoder(w).Encode(userlist)
	})

	mux.HandleFunc("POST /users", func(w http.ResponseWriter, r *http.Request) {
		var newUser User
		if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
			http.Error(w, "Dados invalidos Json", http.StatusBadRequest)
		}

		mu.Lock()
		users[newUser.Id] = newUser
		mu.Unlock()

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(newUser)
	})

	http.ListenAndServe(":8000", mux)
}
