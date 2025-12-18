package config

import (
	"gin-01/entity"
	"log"
)

func RunMigration() {
	err := DB.AutoMigrate(
		&entity.UserEntity{},
	)

	if err != nil {
		log.Fatal("Failed to run migration:", err)
	}

	log.Println("Database migration completed successfully")
}
