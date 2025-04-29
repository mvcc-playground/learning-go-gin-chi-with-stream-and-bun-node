package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-chi/chi/v5"
	"github.com/matheusvcouto/crud-go/utils"
)

func (h *Handler) HelloWord(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Wello word"))
}

func (h *Handler) HelloName(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	if name == "" {
		utils.SendJson(w, http.StatusBadRequest, "Name is required")
		return
	}

	w.Write([]byte("Hello " + name))
}

func (h *Handler) HelloWordGin(c *gin.Context) {
	c.String(http.StatusOK, "hello word")
}
func (h *Handler) HelloNameGin(c *gin.Context) {
	name := c.Param("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name is required"})
		return
	}

	c.String(http.StatusOK, "Hello %s", name)
}

func (h *Handler) HelloNameQuery(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		utils.SendJson(w, http.StatusBadRequest, "Name is required")
		return
	}

	w.Write([]byte("Hello " + name))
}
