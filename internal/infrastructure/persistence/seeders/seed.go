package seeders

import (
	"fastquiz-api/internal/infrastructure/persistence/db"
	"log"
)

type SeederFunc func()

var seeders = []SeederFunc{
	SeedTypes,
	SeedQuizzes,
	SeedQuestionsAndAnswers,
}

func RunAllSeeders() {
	ClearTables()

	for _, seeder := range seeders {
		seeder()
		log.Println("Seed executed successfully!")
	}
}

func ClearTables() {
	tables := []string{"answers", "questions", "quizzes", "types"}

	db.DB.Exec("SET FOREIGN_KEY_CHECKS = 0")
	for _, table := range tables {
		err := db.DB.Exec("TRUNCATE TABLE " + table).Error
		if err != nil {
			log.Printf("Failed to clear table: %s, error: %v", table, err)
		} else {
			log.Printf("Cleared table: %s", table)
		}
	}
	db.DB.Exec("SET FOREIGN_KEY_CHECKS = 1")
}
