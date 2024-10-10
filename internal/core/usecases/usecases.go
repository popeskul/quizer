package usecases

import (
	"context"

	"github.com/popeskul/quizer/internal/core/domain/quiz"
	"github.com/popeskul/quizer/internal/core/ports"
)

type UseCases struct {
	quizUseCases ports.QuizUseCase
}

func NewUseCases(quizService ports.QuizService) ports.UseCases {
	return &UseCases{
		quizUseCases: NewQuizUseCases(quizService),
	}
}

func (uc *UseCases) QuizUseCases() ports.QuizUseCase {
	return uc.quizUseCases
}

type QuizUseCases struct {
	getQuizUseCase       *GetQuizUseCase
	submitAnswersUseCase *SubmitAnswersUseCase
}

func NewQuizUseCases(quizService ports.QuizService) ports.QuizUseCase {
	return &QuizUseCases{
		getQuizUseCase:       NewGetQuizUseCase(quizService),
		submitAnswersUseCase: NewSubmitAnswersUseCase(quizService),
	}
}

func (quc *QuizUseCases) GetQuiz(ctx context.Context) (*quiz.Quiz, error) {
	return quc.getQuizUseCase.Execute(ctx)
}

func (quc *QuizUseCases) SubmitAnswers(ctx context.Context, answers map[string]string) (quiz.QuizResult, error) {
	return quc.submitAnswersUseCase.Execute(ctx, answers)
}
