package main

import (
	"encoding/json"
	"fastquiz-api/internal/application/services"
	"fastquiz-api/internal/domain/entities"
	"fastquiz-api/internal/infrastructure/external/chatgpt"
	"fastquiz-api/internal/infrastructure/persistence/db"
	"fastquiz-api/internal/infrastructure/persistence/repositories"
	"fastquiz-api/internal/unitofworks"
	"fastquiz-api/pkg/config"
	"fastquiz-api/pkg/utils"
	"fmt"
	"github.com/xeipuuv/gojsonschema"
	"log"
	"os"
	"strconv"
	"time"
)

type Question struct {
	Question  string   `json:"question"`
	Options   []string `json:"options"`
	OptionKey int      `json:"option_key"`
}

type Quiz struct {
	QuizTitle string     `json:"quiz_title"`
	Questions []Question `json:"questions"`
}

func main() {

	log.Println("Quiz Worker Started")

	config.LoadConfig()
	db.ConnectDatabase()

	quizRepo := repositories.NewQuizRepository()
	questionRepo := repositories.NewQuestionRepository()
	answerRepo := repositories.NewAnswerRepository()

	uow := unitofworks.NewUnitOfWork(quizRepo, questionRepo, answerRepo)

	ProcessCreatedQuiz(uow)

	log.Println("Quiz Worker Finished")
}

func ProcessCreatedQuiz(uow *unitofworks.UnitOfWork) {

	log.Println("Processing created quiz...")

	quizService := services.NewQuizService(uow)

	quizzes, err := quizService.GetCreatedQuizzes()
	if err != nil {
		log.Printf("Error fetching created quizzes: %v", err)
	}

	chatgptClient := chatgpt.NewChatGptClient()

	promptFile, err := os.ReadFile(utils.GetRootPath() + "/cmd/worker/prompt.txt")
	if err != nil {
		log.Fatalf("Dosya okunurken hata oluştu: %v", err)
	}
	defaultPrompt := string(promptFile)

	for index, quiz := range quizzes {
		startTime := time.Now()

		movieName := quiz.Name
		language := quiz.Language
		questionCount := strconv.Itoa(int(quiz.QuestionCount))
		quizDifficulty := strconv.Itoa(int(quiz.Difficulty))
		questionType := quiz.Type.Name
		prompt := fmt.Sprintf(defaultPrompt, questionCount, questionType, movieName, language, quizDifficulty, questionType, movieName)

		log.Printf("Model: %s Prompt: %s", config.AppConfig.ChatGptModel, prompt)

		response, err := chatgptClient.GenerateResponse(prompt)

		if err != nil {
			log.Printf("[Quiz %d/%d] Error generating quiz ID %d. Duration: %s, Error: %v", index+1, len(quizzes), quiz.ID, time.Since(startTime), err)
			log.Printf("Error generating quiz %d: %v", quiz.ID, err)
			quiz.Action = entities.QuizError
			_ = quizService.Update(&quiz)
			continue
		}

		jsonResult, _ := validateJsonSchema(response.Content)
		if jsonResult {
			var tmpQuiz Quiz
			err = json.Unmarshal([]byte(response.Content), &tmpQuiz)
			if err != nil {
				log.Printf("[Quiz %d/%d] JSON parsing error for quiz ID %d. Duration: %s, Error: %v", index+1, len(quizzes), quiz.ID, time.Since(startTime), err)
				log.Printf("JSON verisini çözme hatası:  %d: %v", quiz.ID, err)
				quiz.Action = entities.QuizError
				_ = quizService.Update(&quiz)
				continue
			}

			MapQuiz(uow, quiz, tmpQuiz)

			quiz.Action = entities.QuizCompleted
			quiz.Slug = utils.FormatSlug(quiz.Name, "quiz")
			_ = quizService.Update(&quiz)

			log.Printf("[Quiz %d/%d] Quiz created successfully. Duration: %s", index+1, len(quizzes), time.Since(startTime))
		} else {
			log.Printf("[Quiz %d/%d] Invalid JSON schema for quiz ID %d. Duration: %s", index+1, len(quizzes), quiz.ID, time.Since(startTime))
			quiz.Action = entities.QuizError
			_ = quizService.Update(&quiz)
		}
	}

}

func MapQuiz(uow *unitofworks.UnitOfWork, quiz entities.Quiz, newQuiz Quiz) entities.Quiz {

	questionService := services.NewQuestionService(uow)
	answerService := services.NewAnswerService(uow)

	for _, q := range newQuiz.Questions {
		question := entities.Question{
			CorrectAnswerID: 0,
			Question:        q.Question,
			Status:          true,
			QuizID:          quiz.ID,
		}

		_ = questionService.Create(&question)

		for i, option := range q.Options {
			answer := entities.Answer{
				QuestionID:    question.ID,
				Answer:        option,
				CorrectAnswer: i == q.OptionKey,
			}
			_ = answerService.Create(&answer)

			question.Answers = append(question.Answers, answer)

			if answer.CorrectAnswer {
				question.CorrectAnswerID = answer.ID
				_ = questionService.Update(&question)
			}
		}

		quiz.Questions = append(quiz.Questions, question)
	}
	return quiz
}

func validateJsonSchema(json string) (bool, error) {
	promptFile, err := os.ReadFile(utils.GetRootPath() + "/cmd/worker/expected_schema.json")
	if err != nil {
		log.Fatalf("Dosya okunurken hata oluştu: %v", err)
	}
	schemaLoader := gojsonschema.NewStringLoader(string(promptFile))
	jsonLoader := gojsonschema.NewStringLoader(json)

	result, err := gojsonschema.Validate(schemaLoader, jsonLoader)
	if err != nil {
		fmt.Println("Error:", err)
		return false, err
	}

	if !result.Valid() {
		fmt.Println("JSON does not comply with the schema. Errors:")
		for _, desc := range result.Errors() {
			fmt.Printf("- %s\n", desc)
		}
		return false, nil
	}
	return true, nil
}
