package routes

import (
	"api-v2/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterHealthRoutes(router *gin.Engine) {
	router.GET("/health", handlers.Health)
}
