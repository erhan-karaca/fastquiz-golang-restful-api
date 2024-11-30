package db

import (
	"fastquiz-api/internal/domain/entities"
	"fastquiz-api/pkg/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.AppConfig.DBUser,
		config.AppConfig.DBPass,
		config.AppConfig.DBHost,
		config.AppConfig.DBPort,
		config.AppConfig.DBName,
	)

	var database *gorm.DB
	var err error

	for {
		database, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Printf("Database connection failed: %v. Retrying in 5 seconds...", err)
			time.Sleep(10 * time.Second) // 10 saniye bekle ve tekrar dene
			continue
		}
		break
	}

	DB = database
	log.Println("Database connected successfully!")
}

func RunMigrations() {
	err := DB.AutoMigrate(&entities.Type{}, &entities.Quiz{}, &entities.Question{}, &entities.Answer{})
	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	log.Println("Migrations ran successfully!")
}
