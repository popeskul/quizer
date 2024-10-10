package ports

import (
	"context"

	"github.com/popeskul/quizer/internal/core/domain/quiz"
)

type QuizService interface {
	GetQuiz(ctx context.Context) (*quiz.Quiz, error)
	SubmitAnswers(ctx context.Context, answers map[string]string) (quiz.QuizResult, error)
}

type Services interface {
	QuizService() QuizService
}
