package quiz

import "sync"

type Question struct {
	ID              string
	Text            string
	Answers         []Answer
	CorrectAnswerID string
}

type Answer struct {
	ID   string
	Text string
}

type Quiz struct {
	Questions []Question
	mu        sync.RWMutex
}

func (q *Quiz) GetQuestions() []Question {
	q.mu.RLock()
	defer q.mu.RUnlock()
	return q.Questions
}

func (q *Quiz) SetQuestions(questions []Question) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.Questions = questions
}

type QuizResult struct {
	CorrectAnswers int
	TotalQuestions int
	Percentile     float64
}
