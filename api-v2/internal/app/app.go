package app

import (
	"api-v2/internal/routes"

	"github.com/gin-gonic/gin"
)

func NewApp() *gin.Engine {
	router := gin.Default()

	routes.RegisterHealthRoutes(router)

	return router
}
