package ports

import (
	"context"

	"github.com/popeskul/quizer/internal/core/domain/quiz"
)

type QuizUseCase interface {
	GetQuiz(ctx context.Context) (*quiz.Quiz, error)
	SubmitAnswers(ctx context.Context, answers map[string]string) (quiz.QuizResult, error)
}

type UseCases interface {
	QuizUseCases() QuizUseCase
}
