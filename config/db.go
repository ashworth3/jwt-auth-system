package config

import (
	"log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"jwt-auth-system/models"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("auth.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	// Auto migrate the User model
	db.AutoMigrate(&models.User{})

	return db
}
