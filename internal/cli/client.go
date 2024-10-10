package cli

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/popeskul/quizer/internal/core/domain/quiz"
)

type Client struct {
	BaseURL string
	HTTP    *http.Client
}

func NewClient(baseURL string) *Client {
	return &Client{
		BaseURL: baseURL,
		HTTP:    &http.Client{},
	}
}

func (c *Client) GetQuiz() (*quiz.Quiz, error) {
	resp, err := c.HTTP.Get(c.BaseURL + "/quiz")
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, string(body))
	}

	var q quiz.Quiz
	if err := json.NewDecoder(resp.Body).Decode(&q); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return &q, nil
}

func (c *Client) SubmitAnswers(answers map[string]string) (*quiz.QuizResult, error) {
	data, err := json.Marshal(answers)
	if err != nil {
		return nil, fmt.Errorf("error marshaling answers: %w", err)
	}

	resp, err := c.HTTP.Post(c.BaseURL+"/submit", "application/json", bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, string(body))
	}

	var result quiz.QuizResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return &result, nil
}
