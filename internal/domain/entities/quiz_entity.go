package entities

import (
	"fastquiz-api/internal/domain/constants"
	"time"
)

type QuizAction int

const (
	QuizCreated QuizAction = iota
	QuizPending
	QuizProcessing
	QuizCompleted
	QuizError
)

type Quiz struct {
	ID            uint               `gorm:"primaryKey" json:"id"`
	TypeID        uint               `json:"type_id"`
	Type          Type               `gorm:"foreignKey:TypeID"`
	Name          string             `json:"name"`
	Status        bool               `json:"status"`
	Language      constants.Language `json:"language"`
	SourceType    string             `json:"source_type"`
	SourceID      string             `json:"source_id"`
	Slug          string             `json:"slug"`
	QuestionCount int8               `gorm:"default:10" json:"question_count"`
	Difficulty    int8               `gorm:"default:5" json:"difficulty"`
	Action        QuizAction         `gorm:"type:tinyint" json:"action"` // created, pending, processing, completed, error
	Questions     []Question         `gorm:"foreignKey:QuizID"`
	CreatedAt     time.Time          `json:"created_at"`
	UpdatedAt     time.Time          `json:"updated_at"`
}
