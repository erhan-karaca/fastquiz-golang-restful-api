package collections

import (
	"fastquiz-api/internal/domain/entities"
	"fastquiz-api/internal/presentation/resources"
)

type QuizCollectionResource struct {
	Quizzes []resources.QuizResource
}

func NewQuizCollectionResource(quizzes []entities.Quiz) QuizCollectionResource {
	quizResources := make([]resources.QuizResource, len(quizzes))
	for i, quiz := range quizzes {
		quizResources[i] = resources.NewQuizResource(quiz)
	}
	return QuizCollectionResource{
		Quizzes: quizResources,
	}

}
