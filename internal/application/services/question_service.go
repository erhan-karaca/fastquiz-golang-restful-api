package services

import (
	"fastquiz-api/internal/domain/entities"
	"fastquiz-api/internal/domain/interfaces"
	"fastquiz-api/internal/unitofworks"
)

type QuestionService struct {
	repo interfaces.IQuestionRepository
}

func NewQuestionService(repo *unitofworks.UnitOfWork) *QuestionService {
	return &QuestionService{repo: repo.QuestionRepository()}
}

func (s *QuestionService) FindAll() ([]entities.Question, error) {
	return s.repo.FindAll()
}

func (s *QuestionService) FindByID(id uint) (*entities.Question, error) {
	return s.repo.FindByID(id)
}

func (s *QuestionService) Create(question *entities.Question) error {
	return s.repo.Create(question)
}

func (s *QuestionService) Update(question *entities.Question) error {
	return s.repo.Update(question)
}

func (s *QuestionService) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s *QuestionService) GetPendingQuestions() ([]entities.Question, error) {
	return s.repo.GetPendingQuestions()
}
