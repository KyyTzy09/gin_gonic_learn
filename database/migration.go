package database

import (
	"gin-01/app/entity"
	"gin-01/config"
	"log"
)

func RunMigration() {
	err := config.DB.AutoMigrate(
		&entity.UserEntity{},
	)

	if err != nil {
		log.Fatal("Failed to run migration:", err)
	}

	log.Println("Database migration completed successfully")
}
