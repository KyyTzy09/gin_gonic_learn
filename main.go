package main

import (
	"fmt"
	"gin-01/config"
	"gin-01/router"
	"log"
)

func main() {
	config.LoadConfig()

  config.ConnectDatabase()
  
  config.RunMigration()

	r := router.SetupRouter()
  
	// Start server
	port := config.AppConfig.App.Port
	log.Printf("Starting %s server on port %s", config.AppConfig.App.Name, port)

	if err := r.Run(":" + port); err != nil {
		log.Fatal(fmt.Sprintf("Failed to start server: %v", err))
	}
}
