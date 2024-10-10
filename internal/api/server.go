package api

import (
	"encoding/json"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/popeskul/quizer/gen/api"
	"github.com/popeskul/quizer/internal/api/handlers"
)

type Server struct {
	handlers *handlers.Handlers
}

func NewServer(handlers *handlers.Handlers) *Server {
	return &Server{
		handlers: handlers,
	}
}

func (s *Server) GetSuccess(w http.ResponseWriter, r *http.Request) {
	s.handlers.Status.Success(w, r)
}

func (s *Server) GetInternal(w http.ResponseWriter, r *http.Request) {
	s.handlers.Status.Internal(w, r)
}

func (s *Server) PostRandom(w http.ResponseWriter, r *http.Request) {
	s.handlers.Status.Random(w, r)
}

func (s *Server) GetQuiz(w http.ResponseWriter, r *http.Request) {
	quiz, err := s.handlers.Quiz.GetQuiz(r.Context())
	if err != nil {
		http.Error(w, "Failed to get quiz", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(quiz)
}

func (s *Server) PostSubmit(w http.ResponseWriter, r *http.Request) {
	var answers map[string]string
	if err := json.NewDecoder(r.Body).Decode(&answers); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	result, err := s.handlers.Quiz.SubmitAnswers(r.Context(), answers)
	if err != nil {
		http.Error(w, "Failed to submit answers", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(result)
}

func (s *Server) GetSwagger(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "api/openapi/api.yaml")
}

func (s *Server) Handler() http.Handler {
	r := chi.NewRouter()

	r.Get("/swagger/*", s.swaggerHandler)

	api.HandlerFromMux(s, r)

	return r
}

func (s *Server) swaggerHandler(w http.ResponseWriter, r *http.Request) {
	swaggerDir := "./static/swagger-ui"

	if r.URL.Path == "/swagger" || r.URL.Path == "/swagger/" {
		http.ServeFile(w, r, filepath.Join(swaggerDir, "index.html"))
		return
	}

	fs := http.StripPrefix("/swagger/", http.FileServer(http.Dir(swaggerDir)))
	fs.ServeHTTP(w, r)
}
