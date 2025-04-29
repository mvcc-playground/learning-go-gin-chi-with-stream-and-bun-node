package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/matheusvcouto/crud-go/models"
	"github.com/matheusvcouto/crud-go/utils"
)

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser models.User
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, "Dados inválidos Json", http.StatusBadRequest)
		return
	}

	user, err := h.userRepo.AddUser(newUser)
	if err != nil {
		utils.SendJson(w, http.StatusInternalServerError, "Erro ao criar usuario")
		return
	}

	w.Header().Set("Content-Type", "aplication/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func (h *Handler) CreateUserGin(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos Json"})
		return
	}

	user, err := h.userRepo.AddUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar usuario"})
		return
	}

	c.JSON(http.StatusCreated, user)
}
