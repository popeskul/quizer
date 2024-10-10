package services

import (
	"github.com/popeskul/quizer/internal/core/ports"
)

type Services struct {
	quizService ports.QuizService
}

func NewServices(repo ports.QuizRepository) ports.Services {
	return &Services{
		quizService: NewQuizService(repo),
	}
}

func (s *Services) QuizService() ports.QuizService {
	return s.quizService
}
