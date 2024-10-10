package main

import (
	"bufio"
	"fmt"
	"github.com/popeskul/quizer/internal/core/domain/quiz"
	"os"
	"strings"

	"github.com/popeskul/quizer/internal/cli"
	"github.com/spf13/cobra"
)

var (
	client *cli.Client
	apiURL string
)

var rootCmd = &cobra.Command{
	Use:   "quizer",
	Short: "A CLI for interacting with the Quizer API",
	Long:  `Quizer CLI allows you to get quizzes and submit answers using the command line.`,
}

var getQuizCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a quiz",
	Run: func(cmd *cobra.Command, args []string) {
		quiz, err := client.GetQuiz()
		if err != nil {
			fmt.Printf("Error getting quiz: %v\n", err)
			return
		}
		displayQuiz(quiz)
	},
}

var submitAnswersCmd = &cobra.Command{
	Use:   "submit",
	Short: "Submit answers to a quiz",
	Run: func(cmd *cobra.Command, args []string) {
		quiz, err := client.GetQuiz()
		if err != nil {
			fmt.Printf("Error getting quiz: %v\n", err)
			return
		}

		answers := getAnswersFromUser(quiz)
		result, err := client.SubmitAnswers(answers)
		if err != nil {
			fmt.Printf("Error submitting answers: %v\n", err)
			return
		}
		displayResult(result)
	},
}

var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset all quiz results",
	Run: func(cmd *cobra.Command, args []string) {
		err := os.WriteFile("quiz_results.json", []byte("[]"), 0644)
		if err != nil {
			fmt.Printf("Error resetting results: %v\n", err)
			return
		}
		fmt.Println("Quiz results have been reset.")
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&apiURL, "api-url", "http://localhost:8080", "API base URL")
	rootCmd.AddCommand(getQuizCmd)
	rootCmd.AddCommand(submitAnswersCmd)
	rootCmd.AddCommand(resetCmd)

	cobra.OnInitialize(initClient)
}

func initClient() {
	client = cli.NewClient(apiURL)
}

func displayQuiz(q *quiz.Quiz) {
	fmt.Println("Quiz Questions:")
	for i, question := range q.Questions {
		fmt.Printf("%d. %s\n", i+1, question.Text)
		for _, answer := range question.Answers {
			fmt.Printf("   %s) %s\n", answer.ID, answer.Text)
		}
		fmt.Println()
	}
}

func getAnswersFromUser(q *quiz.Quiz) map[string]string {
	answers := make(map[string]string)
	reader := bufio.NewReader(os.Stdin)

	for _, question := range q.Questions {
		fmt.Printf("%s\n", question.Text)
		for _, answer := range question.Answers {
			fmt.Printf("%s) %s\n", answer.ID, answer.Text)
		}
		fmt.Print("Your answer: ")
		input, _ := reader.ReadString('\n')
		answers[question.ID] = strings.TrimSpace(input)
	}

	return answers
}

func displayResult(result *quiz.QuizResult) {
	fmt.Printf("Correct Answers: %d\n", result.CorrectAnswers)
	fmt.Printf("Total Questions: %d\n", result.TotalQuestions)
	fmt.Printf("Percentile: %.2f%%\n", result.Percentile)
}
