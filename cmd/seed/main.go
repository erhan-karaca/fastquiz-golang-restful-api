package main

import (
	"fastquiz-api/internal/infrastructure/persistence/db"
	"fastquiz-api/internal/infrastructure/persistence/seeders"
	"fastquiz-api/pkg/config"
)

func main() {
	config.LoadConfig()
	db.ConnectDatabase()
	db.RunMigrations()
	seeders.RunAllSeeders()
}
