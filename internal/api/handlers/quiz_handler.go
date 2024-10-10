package handlers

import (
	"context"

	"github.com/popeskul/quizer/internal/core/domain/quiz"
	"github.com/popeskul/quizer/internal/core/ports"
)

type QuizHandler struct {
	useCases ports.QuizUseCase
}

func NewQuizHandler(useCases ports.QuizUseCase) *QuizHandler {
	return &QuizHandler{
		useCases: useCases,
	}
}

func (h *QuizHandler) GetQuiz(ctx context.Context) (*quiz.Quiz, error) {
	return h.useCases.GetQuiz(ctx)
}

func (h *QuizHandler) SubmitAnswers(ctx context.Context, answers map[string]string) (quiz.QuizResult, error) {
	return h.useCases.SubmitAnswers(ctx, answers)
}
