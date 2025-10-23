package handlers

import (
	"net/http"

	"go-common-logic/internal/models"
	"go-common-logic/internal/services"

	"github.com/gin-gonic/gin"
)

type TwoSumHandler struct {
	service *services.TwoSumService
}

func NewTwoSumHandler(service *services.TwoSumService) *TwoSumHandler {
	return &TwoSumHandler{
		service: service,
	}
}

func (h *TwoSumHandler) Handle(c *gin.Context) {
	var request models.TwoSumRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "invalid request format: " + err.Error(),
		})
		return
	}

	response, err := h.service.Solve(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}
