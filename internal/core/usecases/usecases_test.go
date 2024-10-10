package usecases_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/popeskul/quizer/internal/core/domain/quiz"
	"github.com/popeskul/quizer/internal/core/ports"
	"github.com/popeskul/quizer/internal/core/usecases"
)

func TestGetQuizUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockQuizService := ports.NewMockQuizService(ctrl)
	useCase := usecases.NewGetQuizUseCase(mockQuizService)

	tests := []struct {
		name          string
		setupMocks    func()
		expectedQuiz  *quiz.Quiz
		expectedError error
	}{
		{
			name: "Successful quiz retrieval",
			setupMocks: func() {
				mockQuizService.EXPECT().GetQuiz(gomock.Any()).Return(
					&quiz.Quiz{Questions: []quiz.Question{{ID: "1", Text: "Test Question"}}},
					nil,
				)
			},
			expectedQuiz:  &quiz.Quiz{Questions: []quiz.Question{{ID: "1", Text: "Test Question"}}},
			expectedError: nil,
		},
		{
			name: "Error retrieving quiz",
			setupMocks: func() {
				mockQuizService.EXPECT().GetQuiz(gomock.Any()).Return(nil, errors.New("service error"))
			},
			expectedQuiz:  nil,
			expectedError: errors.New("service error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMocks()

			result, err := useCase.Execute(context.Background())

			assert.Equal(t, tt.expectedQuiz, result)
			assert.Equal(t, tt.expectedError, err)
		})
	}
}

func TestSubmitAnswersUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockQuizService := ports.NewMockQuizService(ctrl)
	useCase := usecases.NewSubmitAnswersUseCase(mockQuizService)

	tests := []struct {
		name           string
		answers        map[string]string
		setupMocks     func()
		expectedResult quiz.QuizResult
		expectedError  error
	}{
		{
			name:    "Successful submission",
			answers: map[string]string{"1": "A", "2": "B"},
			setupMocks: func() {
				mockQuizService.EXPECT().SubmitAnswers(gomock.Any(), map[string]string{"1": "A", "2": "B"}).Return(
					quiz.QuizResult{CorrectAnswers: 2, TotalQuestions: 2, Percentile: 100},
					nil,
				)
			},
			expectedResult: quiz.QuizResult{CorrectAnswers: 2, TotalQuestions: 2, Percentile: 100},
			expectedError:  nil,
		},
		{
			name:    "Error submitting answers",
			answers: map[string]string{"1": "A"},
			setupMocks: func() {
				mockQuizService.EXPECT().SubmitAnswers(gomock.Any(), map[string]string{"1": "A"}).Return(
					quiz.QuizResult{},
					errors.New("service error"),
				)
			},
			expectedResult: quiz.QuizResult{},
			expectedError:  errors.New("service error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMocks()

			result, err := useCase.Execute(context.Background(), tt.answers)

			assert.Equal(t, tt.expectedResult, result)
			assert.Equal(t, tt.expectedError, err)
		})
	}
}

func TestUseCases(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockQuizService := ports.NewMockQuizService(ctrl)
	useCases := usecases.NewUseCases(mockQuizService)

	t.Run("QuizUseCases", func(t *testing.T) {
		quizUseCases := useCases.QuizUseCases()
		assert.NotNil(t, quizUseCases)

		mockQuizService.EXPECT().GetQuiz(gomock.Any()).Return(&quiz.Quiz{}, nil)
		_, err := quizUseCases.GetQuiz(context.Background())
		assert.NoError(t, err)

		mockQuizService.EXPECT().SubmitAnswers(gomock.Any(), gomock.Any()).Return(quiz.QuizResult{}, nil)
		_, err = quizUseCases.SubmitAnswers(context.Background(), map[string]string{})
		assert.NoError(t, err)
	})
}
