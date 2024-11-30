package entities

type Answer struct {
	ID            uint   `gorm:"primaryKey" json:"id"`
	QuestionID    uint   `json:"question_id"`
	Answer        string `json:"answer"`
	CorrectAnswer bool   `json:"correct_answer"`
}
