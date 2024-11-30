package repositories

import (
	"fastquiz-api/internal/domain/entities"
	"fastquiz-api/internal/domain/interfaces"
	"fastquiz-api/internal/infrastructure/persistence/db"
)

type QuestionRepository struct{}

func NewQuestionRepository() interfaces.IQuestionRepository {
	return &QuestionRepository{}
}

func (r *QuestionRepository) FindAll() ([]entities.Question, error) {
	var questions []entities.Question
	err := db.DB.Find(&questions).Error
	if err != nil {
		return nil, err
	}
	return questions, nil
}

func (r *QuestionRepository) FindByID(id uint) (*entities.Question, error) {
	var question entities.Question
	err := db.DB.First(&question, id).Error
	if err != nil {
		return nil, err
	}
	return &question, nil
}

func (r *QuestionRepository) Create(question *entities.Question) error {
	return db.DB.Create(question).Error
}

func (r *QuestionRepository) Update(question *entities.Question) error {
	return db.DB.Save(question).Error
}

func (r *QuestionRepository) Delete(id uint) error {
	return db.DB.Delete(&entities.Question{}, id).Error
}

func (r *QuestionRepository) GetPendingQuestions() ([]entities.Question, error) {
	var questions []entities.Question
	err := db.DB.Where("status = ?", "created").Find(&questions).Error
	return questions, err
}
