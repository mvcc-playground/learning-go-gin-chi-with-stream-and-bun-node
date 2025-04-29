package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userlist := h.userRepo.GetAllUsers()
	json.NewEncoder(w).Encode(userlist)
}

func (h *Handler) GetUsersGin(c *gin.Context) {
	userlist := h.userRepo.GetAllUsers()
	c.JSON(http.StatusOK, userlist)
}
