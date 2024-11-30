package main

import (
	"fastquiz-api/internal/application/services"
	"fastquiz-api/internal/infrastructure/persistence/db"
	"fastquiz-api/internal/infrastructure/persistence/repositories"
	"fastquiz-api/internal/presentation/controllers"
	"fastquiz-api/internal/presentation/middleware"
	"fastquiz-api/internal/unitofworks"
	"fastquiz-api/pkg/config"
	"github.com/labstack/echo/v4"
	middlewareEcho "github.com/labstack/echo/v4/middleware"
	"net/http"
)

func main() {
	config.LoadConfig()
	db.ConnectDatabase()

	quizRepo := repositories.NewQuizRepository()
	questionRepo := repositories.NewQuestionRepository()
	answerRepo := repositories.NewAnswerRepository()

	uow := unitofworks.NewUnitOfWork(quizRepo, questionRepo, answerRepo)

	quizService := services.NewQuizService(uow)

	quizController := &controllers.QuizController{QuizService: quizService}

	e := echo.New()

	// Middleware
	e.Use(middlewareEcho.Logger())
	e.Use(middlewareEcho.Recover())

	// Enable CORS
	e.Use(middlewareEcho.CORSWithConfig(middlewareEcho.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3000", "http://fastquiz-frontend:3000", "http://fastquiz-frontend", "http://fastquiz.link"}, // Frontend URL'niz
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowCredentials: true,
	}))

	quizGroup := e.Group("/api/quizzes")
	quizGroup.GET("", quizController.GetActiveQuizzes, middleware.AuthMiddleware)
	quizGroup.GET("/create", quizController.Create, middleware.AuthMiddleware)
	quizGroup.GET("/:slug", quizController.GetQuizBySlug, middleware.AuthMiddleware)

	e.Logger.Fatal(e.Start(":8080"))
}
