package interfaces

import "fastquiz-api/internal/domain/entities"

type IQuestionRepository interface {
	FindAll() ([]entities.Question, error)
	FindByID(id uint) (*entities.Question, error)
	Create(quiz *entities.Question) error
	Update(quiz *entities.Question) error
	Delete(id uint) error
	GetPendingQuestions() ([]entities.Question, error)
}
