package handlers

import "github.com/popeskul/quizer/internal/core/ports"

type Handlers struct {
	Quiz   *QuizHandler
	Status *StatusHandler
}

func NewHandlers(useCases ports.UseCases) *Handlers {
	return &Handlers{
		Quiz:   NewQuizHandler(useCases.QuizUseCases()),
		Status: NewStatusHandler(),
	}
}
