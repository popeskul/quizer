package ports

import (
	"context"

	"github.com/popeskul/quizer/internal/core/domain/quiz"
)

type QuizRepository interface {
	GetQuiz(ctx context.Context) (*quiz.Quiz, error)
	SetQuiz(ctx context.Context, q *quiz.Quiz) error
	SaveResult(ctx context.Context, result quiz.QuizResult) error
	GetResults(ctx context.Context) ([]quiz.QuizResult, error)
}

type Repository interface {
	QuizRepository() QuizRepository
}
