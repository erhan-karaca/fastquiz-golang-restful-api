package interfaces

import "fastquiz-api/internal/domain/entities"

type IQuizRepository interface {
	FindAll() ([]entities.Quiz, error)
	FindByID(id uint) (*entities.Quiz, error)
	Create(quiz *entities.Quiz) error
	Update(quiz *entities.Quiz) error
	Delete(id uint) error
	GetQuizBySlug(slug string) (*entities.Quiz, error)
	UpdateQuizStatus(id uint, status string) error
	GetCreatedQuizzes() ([]entities.Quiz, error)
	GetActiveQuizzes(page, pageSize int) ([]entities.Quiz, error)
}
