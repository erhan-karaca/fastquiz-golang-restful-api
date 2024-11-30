package repositories

import (
	"fastquiz-api/internal/domain/entities"
	"fastquiz-api/internal/domain/interfaces"
	"fastquiz-api/internal/infrastructure/persistence/db"
)

type AnswerRepository struct{}

func NewAnswerRepository() interfaces.IAnswerRepository {
	return &AnswerRepository{}
}

func (r *AnswerRepository) FindAll() ([]entities.Answer, error) {
	var answers []entities.Answer
	err := db.DB.Find(&answers).Error
	if err != nil {
		return nil, err
	}
	return answers, nil
}

func (r *AnswerRepository) FindByID(id uint) (*entities.Answer, error) {
	var answer entities.Answer
	err := db.DB.First(&answer, id).Error
	if err != nil {
		return nil, err
	}
	return &answer, nil
}

func (r *AnswerRepository) Create(answer *entities.Answer) error {
	return db.DB.Create(answer).Error
}

func (r *AnswerRepository) Update(answer *entities.Answer) error {
	return db.DB.Save(answer).Error
}

func (r *AnswerRepository) Delete(id uint) error {
	return db.DB.Delete(&entities.Answer{}, id).Error
}
