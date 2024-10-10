package usecases

import (
	"context"

	"github.com/popeskul/quizer/internal/core/domain/quiz"
	"github.com/popeskul/quizer/internal/core/ports"
)

type SubmitAnswersUseCase struct {
	quizService ports.QuizService
}

func NewSubmitAnswersUseCase(quizService ports.QuizService) *SubmitAnswersUseCase {
	return &SubmitAnswersUseCase{quizService: quizService}
}

func (uc *SubmitAnswersUseCase) Execute(ctx context.Context, answers map[string]string) (quiz.QuizResult, error) {
	return uc.quizService.SubmitAnswers(ctx, answers)
}
