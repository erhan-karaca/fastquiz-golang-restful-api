package interfaces

import "fastquiz-api/internal/domain/entities"

type IAnswerRepository interface {
	FindAll() ([]entities.Answer, error)
	FindByID(id uint) (*entities.Answer, error)
	Create(answer *entities.Answer) error
	Update(answer *entities.Answer) error
	Delete(id uint) error
}
