package seeders

import (
	"fastquiz-api/internal/domain/entities"
	"fastquiz-api/internal/infrastructure/persistence/db"
	"log"
)

func SeedTypes() {
	types := []entities.Type{
		{Name: "movie", Status: true},
		{Name: "book", Status: true},
		{Name: "music", Status: true},
	}

	for _, t := range types {
		err := db.DB.Where("name = ?", t.Name).FirstOrCreate(&t).Error
		if err != nil {
			log.Printf("Failed to seed type: %s, error: %v", t.Name, err)
		} else {
			log.Printf("Seeded type: %s", t.Name)
		}
	}
}
