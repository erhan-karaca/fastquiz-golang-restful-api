package seeders

import (
	"fastquiz-api/internal/domain/entities"
	"fastquiz-api/internal/infrastructure/persistence/db"
	"log"
)

func SeedQuestionsAndAnswers() {
	// Örnek bir quiz_id alın (daha önce oluşturulmuş bir quiz var olduğunu varsayıyoruz)
	var quiz entities.Quiz
	if err := db.DB.First(&quiz).Error; err != nil {
		log.Println("No quiz found to seed questions and answers.")
		return
	}

	// Sorular ve cevaplar için örnek veriler
	questions := []entities.Question{
		{Question: "Who directed the Harry Potter movies?", Status: true, QuizID: quiz.ID},
		{Question: "What is the name of Harry's pet owl?", Status: true, QuizID: quiz.ID},
	}

	for _, q := range questions {
		// Soruyu veritabanına ekleyin
		if err := db.DB.Create(&q).Error; err != nil {
			log.Printf("Failed to seed question: %s, error: %v", q.Question, err)
			continue
		}

		// Cevapları oluştur
		answers := []entities.Answer{
			{QuestionID: q.ID, Answer: "Chris Columbus"},
			{QuestionID: q.ID, Answer: "Steven Spielberg"},
			{QuestionID: q.ID, Answer: "Alfonso Cuarón"},
			{QuestionID: q.ID, Answer: "David Yates", CorrectAnswer: true}, // Doğru cevap
		}

		// Cevapları veritabanına ekle
		for _, a := range answers {
			if err := db.DB.Create(&a).Error; err != nil {
				log.Printf("Failed to seed answer: %s, error: %v", a.Answer, err)
			}
		}

		// Doğru cevabı işaretle
		q.CorrectAnswerID = answers[3].ID
		if err := db.DB.Save(&q).Error; err != nil {
			log.Printf("Failed to update question with correct answer: %s, error: %v", q.Question, err)
		}
	}
}
