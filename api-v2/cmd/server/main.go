package main

import (
	"log"

	"api-v2/internal/app"
	// "github.com/gin-gonic/gin"
)

// func test(c *gin.Context) {
// 	c.JSON(200, gin.H{
// 		"message": "Gin is running on port 5001",
// 	})
// }

func main() {
	router := app.NewApp()

	// router.GET("/ping", test)

	if err := router.Run(":5001"); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
