package controllers

import (
	"fastquiz-api/internal/application/services"
	"fastquiz-api/internal/domain/constants"
	"fastquiz-api/internal/domain/entities"
	"fastquiz-api/internal/presentation/collections"
	"fastquiz-api/internal/presentation/resources"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type QuizController struct {
	QuizService *services.QuizService
}

func (c *QuizController) GetQuizBySlug(ctx echo.Context) error {
	slug := ctx.Param("slug")
	quiz, err := c.QuizService.GetQuizBySlug(slug)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{"error": "Quiz not found"})
	}

	return ctx.JSON(http.StatusOK, quiz)
}

func (c *QuizController) GetActiveQuizzes(ctx echo.Context) error {
	pageStr := ctx.QueryParam("page")
	pageSizeStr := ctx.QueryParam("pageSize")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 {
		page = 1
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize <= 0 {
		pageSize = 30
	}

	quizzes, err := c.QuizService.GetActiveQuizzes(page, pageSize)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{"error": "Quiz not found"})
	}

	quizCollectionResource := collections.NewQuizCollectionResource(quizzes)
	return ctx.JSON(http.StatusOK, quizCollectionResource.Quizzes)
}

func (c *QuizController) Create(ctx echo.Context) error {
	nameStr := ctx.QueryParam("name")
	difficultyStr := ctx.QueryParam("difficulty")

	if nameStr == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "Name parameter is required and cannot be empty",
		})
	}

	var difficulty int8

	if difficultyStr != "" {
		difficultyInt, err := strconv.Atoi(difficultyStr)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, map[string]string{
				"error": "Invalid difficulty parameter",
			})
		}
		if difficultyInt < -128 || difficultyInt > 127 {
			return ctx.JSON(http.StatusBadRequest, map[string]string{
				"error": "Difficulty parameter out of range for 1-10",
			})
		}
		difficulty = int8(difficultyInt)
	} else {
		difficulty = 5
	}

	quiz := entities.Quiz{
		TypeID:     1,
		Type:       entities.Type{ID: 1},
		Name:       nameStr,
		Status:     true,
		Language:   constants.Languages.Turkish,
		SourceType: "imdb",
		SourceID:   "tt0068646",
		Action:     entities.QuizCreated,
		Difficulty: difficulty,
	}

	err := c.QuizService.Create(&quiz)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error":   "Quiz could not be created!",
			"details": err.Error(),
		})
	} else {
		quizResource := resources.NewQuizResource(quiz)
		return ctx.JSON(http.StatusOK, quizResource)
	}
}
