package services

import (
	"fastquiz-api/internal/domain/entities"
	"fastquiz-api/internal/domain/interfaces"
	"fastquiz-api/internal/unitofworks"
)

type AnswerService struct {
	repo interfaces.IAnswerRepository
}

func NewAnswerService(repo *unitofworks.UnitOfWork) *AnswerService {
	return &AnswerService{repo: repo.AnswerRepository()}
}

func (s *AnswerService) FindAll() ([]entities.Answer, error) {
	return s.repo.FindAll()
}

func (s *AnswerService) FindByID(id uint) (*entities.Answer, error) {
	return s.repo.FindByID(id)
}

func (s *AnswerService) Create(answer *entities.Answer) error {
	return s.repo.Create(answer)
}

func (s *AnswerService) Update(answer *entities.Answer) error {
	return s.repo.Update(answer)
}

func (s *AnswerService) Delete(id uint) error {
	return s.repo.Delete(id)
}
