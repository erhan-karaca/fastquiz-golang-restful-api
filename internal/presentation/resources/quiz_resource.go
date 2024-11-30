package resources

import "fastquiz-api/internal/domain/entities"

type QuizResource struct {
	ID         uint
	Title      string
	Slug       string
	Difficulty int8
	Type       string
}

func NewQuizResource(quiz entities.Quiz) QuizResource {
	return QuizResource{
		ID:         quiz.ID,
		Title:      quiz.Name,
		Slug:       quiz.Slug,
		Difficulty: quiz.Difficulty,
		Type:       quiz.Type.Name,
	}
}
