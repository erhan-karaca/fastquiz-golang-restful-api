package repositories

import (
	"fastquiz-api/internal/domain/entities"
	"fastquiz-api/internal/domain/interfaces"
	"fastquiz-api/internal/infrastructure/persistence/db"
)

type QuizRepository struct{}

func NewQuizRepository() interfaces.IQuizRepository {
	return &QuizRepository{}
}

func (r *QuizRepository) FindAll() ([]entities.Quiz, error) {
	var quizzes []entities.Quiz
	err := db.DB.Find(&quizzes).Error
	if err != nil {
		return nil, err
	}
	return quizzes, nil
}

func (r *QuizRepository) FindByID(id uint) (*entities.Quiz, error) {
	var quiz entities.Quiz
	err := db.DB.First(&quiz, id).Error
	if err != nil {
		return nil, err
	}
	return &quiz, nil
}

func (r *QuizRepository) Create(quiz *entities.Quiz) error {
	return db.DB.Create(quiz).Error
}

func (r *QuizRepository) Update(quiz *entities.Quiz) error {
	return db.DB.Save(quiz).Error
}

func (r *QuizRepository) Delete(id uint) error {
	return db.DB.Delete(&entities.Quiz{}, id).Error
}

func (r *QuizRepository) GetQuizBySlug(slug string) (*entities.Quiz, error) {
	var quiz entities.Quiz
	err := db.DB.Preload("Type").Preload("Questions", "status = ?", true).Preload("Questions.Answers").Where("slug = ?", slug).First(&quiz).Error
	if err != nil {
		return nil, err
	}
	return &quiz, err
}

func (r *QuizRepository) UpdateQuizStatus(quizID uint, status string) error {
	return db.DB.Model(&entities.Quiz{}).Where("id = ?", quizID).Update("action", status).Error
}

func (r *QuizRepository) GetCreatedQuizzes() ([]entities.Quiz, error) {
	var quizzes []entities.Quiz
	err := db.DB.Where("action = ? AND status = ?", entities.QuizCreated, 1).Find(&quizzes).Error
	if err != nil {
		return nil, err
	}
	return quizzes, nil
}

func (r *QuizRepository) GetActiveQuizzes(page, pageSize int) ([]entities.Quiz, error) {
	var quizzes []entities.Quiz
	offset := (page - 1) * pageSize
	err := db.DB.Preload("Type").Where("action = ? AND status = ?", entities.QuizCompleted, 1).Order("id DESC").Limit(pageSize).Offset(offset).Find(&quizzes).Error
	if err != nil {
		return nil, err
	}
	return quizzes, nil
}
