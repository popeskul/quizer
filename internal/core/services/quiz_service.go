package services

import (
	"context"
	"log"
	"sync"

	"github.com/popeskul/quizer/internal/core/domain/quiz"
	"github.com/popeskul/quizer/internal/core/ports"
)

type quizService struct {
	repo ports.QuizRepository
	mu   sync.Mutex
}

func NewQuizService(repo ports.QuizRepository) ports.QuizService {
	return &quizService{repo: repo}
}

func (s *quizService) SubmitAnswers(ctx context.Context, answers map[string]string) (quiz.QuizResult, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	q, err := s.repo.GetQuiz(ctx)
	if err != nil {
		log.Printf("Error getting quiz: %v", err)
		return quiz.QuizResult{}, err
	}

	correctAnswers := 0
	for _, question := range q.GetQuestions() {
		if answers[question.ID] == question.CorrectAnswerID {
			correctAnswers++
		}
	}

	totalQuestions := len(q.GetQuestions())

	log.Printf("Correct answers: %d, Total questions: %d", correctAnswers, totalQuestions)

	result := quiz.QuizResult{
		CorrectAnswers: correctAnswers,
		TotalQuestions: totalQuestions,
	}

	allResults, err := s.repo.GetResults(ctx)
	if err != nil {
		log.Printf("Error getting results: %v", err)
		return quiz.QuizResult{}, err
	}

	log.Printf("Number of previous results: %d", len(allResults))

	if len(allResults) > 0 {
		worseThanCurrent := 0
		for _, r := range allResults {
			if r.CorrectAnswers < correctAnswers {
				worseThanCurrent++
			}
		}
		result.Percentile = float64(worseThanCurrent) / float64(len(allResults)+1) * 100
		log.Printf("Results worse than current: %d, Total results (including current): %d", worseThanCurrent, len(allResults)+1)
	} else {
		result.Percentile = 100
	}

	log.Printf("Calculated percentile: %.2f", result.Percentile)

	err = s.repo.SaveResult(ctx, result)
	if err != nil {
		log.Printf("Error saving result: %v", err)
		return quiz.QuizResult{}, err
	}

	return result, nil
}

func (s *quizService) GetQuiz(ctx context.Context) (*quiz.Quiz, error) {
	return s.repo.GetQuiz(ctx)
}
