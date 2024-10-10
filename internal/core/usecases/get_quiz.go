package usecases

import (
	"context"

	"github.com/popeskul/quizer/internal/core/domain/quiz"
	"github.com/popeskul/quizer/internal/core/ports"
)

type GetQuizUseCase struct {
	quizService ports.QuizService
}

func NewGetQuizUseCase(quizService ports.QuizService) *GetQuizUseCase {
	return &GetQuizUseCase{quizService: quizService}
}

func (uc *GetQuizUseCase) Execute(ctx context.Context) (*quiz.Quiz, error) {
	return uc.quizService.GetQuiz(ctx)
}
