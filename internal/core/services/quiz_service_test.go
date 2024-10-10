package services_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/popeskul/quizer/internal/core/domain/quiz"
	"github.com/popeskul/quizer/internal/core/ports"
	"github.com/popeskul/quizer/internal/core/services"
)

func TestQuizService_GetQuiz(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := ports.NewMockQuizRepository(ctrl)
	quizService := services.NewQuizService(mockRepo)

	tests := []struct {
		name          string
		setupMocks    func()
		expectedQuiz  *quiz.Quiz
		expectedError error
	}{
		{
			name: "Successful quiz retrieval",
			setupMocks: func() {
				mockRepo.EXPECT().GetQuiz(gomock.Any()).Return(&quiz.Quiz{
					Questions: []quiz.Question{{ID: "1", Text: "Test Question"}},
				}, nil)
			},
			expectedQuiz: &quiz.Quiz{
				Questions: []quiz.Question{{ID: "1", Text: "Test Question"}},
			},
			expectedError: nil,
		},
		{
			name: "Error retrieving quiz",
			setupMocks: func() {
				mockRepo.EXPECT().GetQuiz(gomock.Any()).Return(nil, errors.New("database error"))
			},
			expectedQuiz:  nil,
			expectedError: errors.New("database error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMocks()

			result, err := quizService.GetQuiz(context.Background())

			assert.Equal(t, tt.expectedQuiz, result)
			assert.Equal(t, tt.expectedError, err)
		})
	}
}

func TestQuizService_SubmitAnswers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := ports.NewMockQuizRepository(ctrl)
	quizService := services.NewQuizService(mockRepo)

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
				mockRepo.EXPECT().GetQuiz(gomock.Any()).Return(&quiz.Quiz{
					Questions: []quiz.Question{
						{ID: "1", CorrectAnswerID: "A"},
						{ID: "2", CorrectAnswerID: "B"},
					},
				}, nil)
				mockRepo.EXPECT().GetResults(gomock.Any()).Return([]quiz.QuizResult{}, nil)
				mockRepo.EXPECT().SaveResult(gomock.Any(), gomock.Any()).Return(nil)
			},
			expectedResult: quiz.QuizResult{
				CorrectAnswers: 2,
				TotalQuestions: 2,
				Percentile:     100,
			},
			expectedError: nil,
		},
		{
			name:    "Error getting quiz",
			answers: map[string]string{"1": "A"},
			setupMocks: func() {
				mockRepo.EXPECT().GetQuiz(gomock.Any()).Return(nil, errors.New("database error"))
			},
			expectedResult: quiz.QuizResult{},
			expectedError:  errors.New("database error"),
		},
		{
			name:    "Error getting results",
			answers: map[string]string{"1": "A"},
			setupMocks: func() {
				mockRepo.EXPECT().GetQuiz(gomock.Any()).Return(&quiz.Quiz{
					Questions: []quiz.Question{{ID: "1", CorrectAnswerID: "A"}},
				}, nil)
				mockRepo.EXPECT().GetResults(gomock.Any()).Return(nil, errors.New("database error"))
			},
			expectedResult: quiz.QuizResult{},
			expectedError:  errors.New("database error"),
		},
		{
			name:    "Error saving result",
			answers: map[string]string{"1": "A"},
			setupMocks: func() {
				mockRepo.EXPECT().GetQuiz(gomock.Any()).Return(&quiz.Quiz{
					Questions: []quiz.Question{{ID: "1", CorrectAnswerID: "A"}},
				}, nil)
				mockRepo.EXPECT().GetResults(gomock.Any()).Return([]quiz.QuizResult{}, nil)
				mockRepo.EXPECT().SaveResult(gomock.Any(), gomock.Any()).Return(errors.New("database error"))
			},
			expectedResult: quiz.QuizResult{},
			expectedError:  errors.New("database error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMocks()

			result, err := quizService.SubmitAnswers(context.Background(), tt.answers)

			assert.Equal(t, tt.expectedResult, result)
			assert.Equal(t, tt.expectedError, err)
		})
	}
}
