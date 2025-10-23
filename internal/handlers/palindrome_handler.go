package handlers

import (
	"net/http"

	"go-common-logic/internal/models"
	"go-common-logic/internal/services"

	"github.com/gin-gonic/gin"
)

type PalindromeHandler struct {
	service *services.PalindromeService
}

func NewPalindromeHandler(service *services.PalindromeService) *PalindromeHandler {
	return &PalindromeHandler{
		service: service,
	}
}

func (h *PalindromeHandler) Check(c *gin.Context) {
	var request models.PalindromeRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "invalid request format: " + err.Error(),
		})
		return
	}

	response := h.service.Check(&request)
	c.JSON(http.StatusOK, response)
}
