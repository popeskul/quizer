package repository

import "github.com/popeskul/quizer/internal/core/domain/quiz"

func CreateDefaultQuiz() *quiz.Quiz {
	return &quiz.Quiz{
		Questions: []quiz.Question{
			{
				ID:   "q1",
				Text: "What is the capital of France?",
				Answers: []quiz.Answer{
					{ID: "1", Text: "London"},
					{ID: "2", Text: "Berlin"},
					{ID: "3", Text: "Paris"},
					{ID: "4", Text: "Madrid"},
				},
				CorrectAnswerID: "3",
			},
			{
				ID:   "q2",
				Text: "Which planet is known as the Red Planet?",
				Answers: []quiz.Answer{
					{ID: "1", Text: "Venus"},
					{ID: "2", Text: "Mars"},
					{ID: "3", Text: "Jupiter"},
					{ID: "4", Text: "Saturn"},
				},
				CorrectAnswerID: "2",
			},
		},
	}
}
