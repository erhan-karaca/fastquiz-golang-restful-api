package services

import (
	"fastquiz-api/internal/domain/entities"
	"fastquiz-api/internal/domain/interfaces"
	"fastquiz-api/internal/unitofworks"
	"fastquiz-api/pkg/utils"
	"fmt"
	"os/exec"
)

type QuizService struct {
	repo interfaces.IQuizRepository
}

func NewQuizService(repo *unitofworks.UnitOfWork) *QuizService {
	return &QuizService{repo: repo.QuizRepository()}
}

func (s *QuizService) FindAll() ([]entities.Quiz, error) {
	return s.repo.FindAll()
}

func (s *QuizService) FindByID(id uint) (*entities.Quiz, error) {
	return s.repo.FindByID(id)
}

func (s *QuizService) Create(quiz *entities.Quiz) error {
	err := s.repo.Create(quiz)
	go RunWorker()
	return err
}

func (s *QuizService) Update(quiz *entities.Quiz) error {
	return s.repo.Update(quiz)
}

func (s *QuizService) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s *QuizService) GetQuizBySlug(slug string) (*entities.Quiz, error) {
	return s.repo.GetQuizBySlug(slug)
}

func (s *QuizService) UpdateQuizStatus(quizID uint, status string) error {
	return s.repo.UpdateQuizStatus(quizID, status)
}

func (s *QuizService) GetCreatedQuizzes() ([]entities.Quiz, error) {
	return s.repo.GetCreatedQuizzes()
}

func (s *QuizService) GetActiveQuizzes(page, pageSize int) ([]entities.Quiz, error) {
	return s.repo.GetActiveQuizzes(page, pageSize)
}

func RunWorker() {
	cmd := exec.Command(utils.GetRootPath() + "/worker")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Worker çalıştırılamadı: %s\n", err)
	} else {
		fmt.Printf("Worker çıktı:\n%s\n", string(output))
	}
}
