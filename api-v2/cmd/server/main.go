package main

import (
	"log"

	"api-v2/internal/app"
	"api-v2/internal/config"
)

func main() {
	cfg := config.Load()
	server := app.NewApp()

	log.Printf("Server running on :%s", cfg.Port)

	if err := server.Run(":" + cfg.Port); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
