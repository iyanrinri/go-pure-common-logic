package routes

import (
	"go-common-logic/internal/handlers"
	"go-common-logic/internal/services"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	router := gin.Default()

	twoSumService := services.NewTwoSumService()

	twoSumHandler := handlers.NewTwoSumHandler(twoSumService)

	v1 := router.Group("/api/v1")
	{
		v1.POST("/two-sum", twoSumHandler.Handle)
	}

	return router
}
