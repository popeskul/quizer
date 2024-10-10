package handlers_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/popeskul/quizer/internal/api/handlers"
	"github.com/popeskul/quizer/internal/core/domain/quiz"
	"github.com/popeskul/quizer/internal/core/ports"
)

func TestQuizHandler_GetQuiz(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockQuizUseCase := ports.NewMockQuizUseCase(ctrl)
	handler := handlers.NewQuizHandler(mockQuizUseCase)

	t.Run("Successful quiz retrieval", func(t *testing.T) {
		expectedQuiz := &quiz.Quiz{Questions: []quiz.Question{{ID: "1", Text: "Test Question"}}}
		mockQuizUseCase.EXPECT().GetQuiz(gomock.Any()).Return(expectedQuiz, nil)

		quiz, err := handler.GetQuiz(context.Background())

		assert.NoError(t, err)
		assert.Equal(t, expectedQuiz, quiz)
	})

	t.Run("Error getting quiz", func(t *testing.T) {
		mockQuizUseCase.EXPECT().GetQuiz(gomock.Any()).Return(nil, errors.New("use case error"))

		quiz, err := handler.GetQuiz(context.Background())

		assert.Error(t, err)
		assert.Nil(t, quiz)
	})
}

func TestQuizHandler_SubmitAnswers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockQuizUseCase := ports.NewMockQuizUseCase(ctrl)
	handler := handlers.NewQuizHandler(mockQuizUseCase)

	t.Run("Successful submit", func(t *testing.T) {
		answers := map[string]string{"1": "A", "2": "B"}
		expectedResult := quiz.QuizResult{CorrectAnswers: 2, TotalQuestions: 2, Percentile: 100}
		mockQuizUseCase.EXPECT().SubmitAnswers(gomock.Any(), answers).Return(expectedResult, nil)

		result, err := handler.SubmitAnswers(context.Background(), answers)

		assert.NoError(t, err)
		assert.Equal(t, expectedResult, result)
	})

	t.Run("Error submitting answers", func(t *testing.T) {
		answers := map[string]string{"1": "A", "2": "B"}
		mockQuizUseCase.EXPECT().SubmitAnswers(gomock.Any(), answers).Return(quiz.QuizResult{}, errors.New("use case error"))

		result, err := handler.SubmitAnswers(context.Background(), answers)

		assert.Error(t, err)
		assert.Equal(t, quiz.QuizResult{}, result)
	})
}

func TestStatusHandler_Success(t *testing.T) {
	handler := handlers.NewStatusHandler()
	req := httptest.NewRequest(http.MethodGet, "/success", nil)
	w := httptest.NewRecorder()

	handler.Success(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestStatusHandler_Internal(t *testing.T) {
	handler := handlers.NewStatusHandler()
	req := httptest.NewRequest(http.MethodGet, "/internal", nil)
	w := httptest.NewRecorder()

	handler.Internal(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestStatusHandler_Random(t *testing.T) {
	handler := handlers.NewStatusHandler()
	req := httptest.NewRequest(http.MethodPost, "/random", nil)

	for i := 0; i < 100; i++ {
		w := httptest.NewRecorder()

		func() {
			defer func() {
				if r := recover(); r != nil {
					assert.Equal(t, "Random panic", r)
				}
			}()
			handler.Random(w, req)
		}()

		if w.Code != 0 {
			assert.Contains(t, []int{http.StatusOK, http.StatusInternalServerError}, w.Code)
		}
	}
}
