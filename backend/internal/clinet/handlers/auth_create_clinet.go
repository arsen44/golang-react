package handlers

import (
	"backend/internal/clinet/repo"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ClientHandler struct {
	repo repo.ClientRepoInterface
}

func NewClientHandler(repo repo.ClientRepoInterface) ClinetHandlersInterface {
	return &ClientHandler{repo: repo}
}

func (h *ClientHandler) CreateClient(c *gin.Context) {
	var request struct {
		PhoneNumber string `json:"phone_number"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	accessToken, refreshToken, err := h.repo.CreateClient(request.PhoneNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}
