package api_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/popeskul/quizer/internal/api"
	"github.com/popeskul/quizer/internal/api/handlers"
	"github.com/popeskul/quizer/internal/core/domain/quiz"
	"github.com/popeskul/quizer/internal/core/ports"
)

func TestServer_GetQuiz(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUseCases := ports.NewMockUseCases(ctrl)
	mockQuizUseCase := ports.NewMockQuizUseCase(ctrl)
	mockUseCases.EXPECT().QuizUseCases().Return(mockQuizUseCase).AnyTimes()

	server := api.NewServer(handlers.NewHandlers(mockUseCases))

	t.Run("Successful quiz retrieval", func(t *testing.T) {
		expectedQuiz := &quiz.Quiz{Questions: []quiz.Question{{ID: "1", Text: "Test Question"}}}
		mockQuizUseCase.EXPECT().GetQuiz(gomock.Any()).Return(expectedQuiz, nil)

		req := httptest.NewRequest(http.MethodGet, "/quiz", nil)
		w := httptest.NewRecorder()

		server.GetQuiz(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var returnedQuiz quiz.Quiz
		err := json.Unmarshal(w.Body.Bytes(), &returnedQuiz)
		require.NoError(t, err)
		assert.Equal(t, expectedQuiz, &returnedQuiz)
	})

	t.Run("Error getting quiz", func(t *testing.T) {
		mockQuizUseCase.EXPECT().GetQuiz(gomock.Any()).Return(nil, assert.AnError)

		req := httptest.NewRequest(http.MethodGet, "/quiz", nil)
		w := httptest.NewRecorder()

		server.GetQuiz(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}

func TestServer_PostSubmit(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUseCases := ports.NewMockUseCases(ctrl)
	mockQuizUseCase := ports.NewMockQuizUseCase(ctrl)
	mockUseCases.EXPECT().QuizUseCases().Return(mockQuizUseCase).AnyTimes()

	server := api.NewServer(handlers.NewHandlers(mockUseCases))

	t.Run("Successful submit", func(t *testing.T) {
		answers := map[string]string{"1": "A", "2": "B"}
		expectedResult := quiz.QuizResult{CorrectAnswers: 2, TotalQuestions: 2, Percentile: 100}

		mockQuizUseCase.EXPECT().SubmitAnswers(gomock.Any(), answers).Return(expectedResult, nil)

		body, _ := json.Marshal(answers)
		req := httptest.NewRequest(http.MethodPost, "/submit", strings.NewReader(string(body)))
		w := httptest.NewRecorder()

		server.PostSubmit(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var returnedResult quiz.QuizResult
		err := json.Unmarshal(w.Body.Bytes(), &returnedResult)
		require.NoError(t, err)
		assert.Equal(t, expectedResult, returnedResult)
	})

	t.Run("Invalid request body", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/submit", strings.NewReader("invalid json"))
		w := httptest.NewRecorder()

		server.PostSubmit(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Error submitting answers", func(t *testing.T) {
		answers := map[string]string{"1": "A", "2": "B"}
		mockQuizUseCase.EXPECT().SubmitAnswers(gomock.Any(), answers).Return(quiz.QuizResult{}, assert.AnError)

		body, _ := json.Marshal(answers)
		req := httptest.NewRequest(http.MethodPost, "/submit", strings.NewReader(string(body)))
		w := httptest.NewRecorder()

		server.PostSubmit(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}
