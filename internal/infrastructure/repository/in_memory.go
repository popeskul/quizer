package repository

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"os"
	"sync"

	"github.com/popeskul/quizer/internal/core/domain/quiz"
	"github.com/popeskul/quizer/internal/core/ports"
)

var ResultsFile = "quiz_results.json"

type InMemoryRepository struct {
	quiz    *quiz.Quiz
	results []quiz.QuizResult
	mu      sync.RWMutex
}

func NewInMemoryRepository() ports.Repository {
	repo := &InMemoryRepository{}
	repo.loadResults()
	return repo
}

func (r *InMemoryRepository) QuizRepository() ports.QuizRepository {
	return r
}

func (r *InMemoryRepository) GetQuiz(ctx context.Context) (*quiz.Quiz, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		if r.quiz == nil {
			return nil, ErrQuizNotSet
		}
		return r.quiz, nil
	}
}

func (r *InMemoryRepository) SetQuiz(ctx context.Context, q *quiz.Quiz) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		r.quiz = q
		return nil
	}
}

func (r *InMemoryRepository) SaveResult(ctx context.Context, result quiz.QuizResult) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		r.results = append(r.results, result)
		return r.saveResults()
	}
}

func (r *InMemoryRepository) GetResults(ctx context.Context) ([]quiz.QuizResult, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		results := make([]quiz.QuizResult, len(r.results))
		copy(results, r.results)
		return results, nil
	}
}

func (r *InMemoryRepository) loadResults() {
	data, err := os.ReadFile(ResultsFile)
	if err != nil {
		if !os.IsNotExist(err) {
			log.Printf("Error reading results file: %v", err)
		}
		return
	}

	err = json.Unmarshal(data, &r.results)
	if err != nil {
		log.Printf("Error unmarshaling results: %v", err)
	}
	log.Printf("Loaded %d results from file", len(r.results))
}

func (r *InMemoryRepository) saveResults() error {
	data, err := json.Marshal(r.results)
	if err != nil {
		log.Printf("Error marshaling results: %v", err)
		return err
	}

	err = os.WriteFile(ResultsFile, data, 0644)
	if err != nil {
		log.Printf("Error writing results file: %v", err)
		return err
	}

	log.Printf("Saved %d results to file", len(r.results))
	return nil
}

var ErrQuizNotSet = errors.New("quiz not set")
