package entities

type Question struct {
	ID              uint     `gorm:"primaryKey" json:"id"`
	CorrectAnswerID uint     `json:"correct_answer_id"`
	Question        string   `json:"question"`
	Status          bool     `json:"status"`
	QuizID          uint     `json:"quiz_id"`
	Answers         []Answer `gorm:"foreignKey:QuestionID"`
}
