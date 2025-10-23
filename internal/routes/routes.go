package routes

import (
	"go-common-logic/internal/handlers"
	"go-common-logic/internal/services"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	router := gin.Default()

	// Initialize services
	twoSumService := services.NewTwoSumService()
	palindromeService := services.NewPalindromeService()

	// Initialize handlers
	twoSumHandler := handlers.NewTwoSumHandler(twoSumService)
	palindromeHandler := handlers.NewPalindromeHandler(palindromeService)

	// API v1 routes
	v1 := router.Group("/api/v1")
	{
		v1.POST("/two-sum", twoSumHandler.Handle)
		v1.POST("/palindrome", palindromeHandler.Check)
	}

	return router
}
