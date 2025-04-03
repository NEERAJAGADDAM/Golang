package handlers

import (
	"cms/Internals/repository"
	"cms/Internals/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	repo       repository.DbRepository
	jwtService *services.JWTService
}

func NewHandler(repo repository.DbRepository, jwtService *services.JWTService) *Handler {
	return &Handler{
		repo:       repo,
		jwtService: jwtService,
	}
}

func (h *Handler) GetAllUsers(c *gin.Context) {
	users, err := h.repo.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get users"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"users": users})
}
