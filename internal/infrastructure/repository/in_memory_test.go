package repository_test

import (
	"context"
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/popeskul/quizer/internal/core/domain/quiz"
	"github.com/popeskul/quizer/internal/infrastructure/repository"
)

const testResultsFile = "test_quiz_results.json"

func setupTest(t *testing.T) func() {
	oldResultsFile := repository.ResultsFile
	repository.ResultsFile = testResultsFile

	err := os.WriteFile(testResultsFile, []byte("[]"), 0644)
	require.NoError(t, err)

	return func() {
		repository.ResultsFile = oldResultsFile
		os.Remove(testResultsFile)
	}
}

func TestInMemoryRepository_GetQuiz(t *testing.T) {
	cleanup := setupTest(t)
	defer cleanup()

	tests := []struct {
		name          string
		setupRepo     func(*repository.InMemoryRepository)
		expectedQuiz  *quiz.Quiz
		expectedError error
	}{
		{
			name: "Successful quiz retrieval",
			setupRepo: func(r *repository.InMemoryRepository) {
				q := &quiz.Quiz{
					Questions: []quiz.Question{
						{ID: "q1", Text: "Test question"},
					},
				}
				err := r.SetQuiz(context.Background(), q)
				require.NoError(t, err)
			},
			expectedQuiz: &quiz.Quiz{
				Questions: []quiz.Question{
					{ID: "q1", Text: "Test question"},
				},
			},
			expectedError: nil,
		},
		{
			name:          "Quiz not set",
			setupRepo:     func(r *repository.InMemoryRepository) {},
			expectedQuiz:  nil,
			expectedError: repository.ErrQuizNotSet,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := repository.NewInMemoryRepository()
			tt.setupRepo(repo.(*repository.InMemoryRepository))

			quiz, err := repo.QuizRepository().GetQuiz(context.Background())

			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError, err)
				assert.Nil(t, quiz)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, quiz)
				assert.Equal(t, tt.expectedQuiz, quiz)
			}
		})
	}
}

func TestInMemoryRepository_SaveResult(t *testing.T) {
	cleanup := setupTest(t)
	defer cleanup()

	tests := []struct {
		name           string
		result         quiz.QuizResult
		expectedError  error
		expectedResult quiz.QuizResult
	}{
		{
			name: "Save result successfully",
			result: quiz.QuizResult{
				CorrectAnswers: 5,
				TotalQuestions: 10,
				Percentile:     50.0,
			},
			expectedError: nil,
			expectedResult: quiz.QuizResult{
				CorrectAnswers: 5,
				TotalQuestions: 10,
				Percentile:     50.0,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := repository.NewInMemoryRepository()

			err := repo.QuizRepository().SaveResult(context.Background(), tt.result)

			assert.Equal(t, tt.expectedError, err)

			results, err := repo.QuizRepository().GetResults(context.Background())
			assert.NoError(t, err)
			assert.Len(t, results, 1)
			assert.Equal(t, tt.expectedResult, results[0])

			// Check if the result was saved to the file
			data, err := os.ReadFile(testResultsFile)
			assert.NoError(t, err)

			var savedResults []quiz.QuizResult
			err = json.Unmarshal(data, &savedResults)
			assert.NoError(t, err)
			assert.Len(t, savedResults, 1)
			assert.Equal(t, tt.expectedResult, savedResults[0])
		})
	}
}

func TestInMemoryRepository_GetResults(t *testing.T) {
	tests := []struct {
		name            string
		setupResults    []quiz.QuizResult
		expectedResults []quiz.QuizResult
		expectedError   error
	}{
		{
			name: "Get results successfully",
			setupResults: []quiz.QuizResult{
				{CorrectAnswers: 5, TotalQuestions: 10, Percentile: 50.0},
				{CorrectAnswers: 8, TotalQuestions: 10, Percentile: 80.0},
			},
			expectedResults: []quiz.QuizResult{
				{CorrectAnswers: 5, TotalQuestions: 10, Percentile: 50.0},
				{CorrectAnswers: 8, TotalQuestions: 10, Percentile: 80.0},
			},
			expectedError: nil,
		},
		{
			name:            "Get empty results",
			setupResults:    []quiz.QuizResult{},
			expectedResults: []quiz.QuizResult{},
			expectedError:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cleanup := setupTest(t)
			defer cleanup()

			repo := repository.NewInMemoryRepository()

			// Setup results
			for _, result := range tt.setupResults {
				err := repo.QuizRepository().SaveResult(context.Background(), result)
				require.NoError(t, err)
			}

			results, err := repo.QuizRepository().GetResults(context.Background())

			assert.Equal(t, tt.expectedError, err)
			assert.Equal(t, tt.expectedResults, results)
		})
	}
}
