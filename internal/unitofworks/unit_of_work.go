package unitofworks

import (
	"fastquiz-api/internal/domain/interfaces"
)

type UnitOfWork struct {
	quizRepo     interfaces.IQuizRepository
	questionRepo interfaces.IQuestionRepository
	answerRepo   interfaces.IAnswerRepository
}

func NewUnitOfWork(
	quizRepo interfaces.IQuizRepository,
	questionRepo interfaces.IQuestionRepository,
	answerRepo interfaces.IAnswerRepository,
) *UnitOfWork {
	return &UnitOfWork{
		quizRepo:     quizRepo,
		questionRepo: questionRepo,
		answerRepo:   answerRepo,
	}
}

func (uow *UnitOfWork) QuizRepository() interfaces.IQuizRepository {
	return uow.quizRepo
}

func (uow *UnitOfWork) QuestionRepository() interfaces.IQuestionRepository {
	return uow.questionRepo
}

func (uow *UnitOfWork) AnswerRepository() interfaces.IAnswerRepository {
	return uow.answerRepo
}
