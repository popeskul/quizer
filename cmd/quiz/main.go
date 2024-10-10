package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/popeskul/quizer/internal/api"
	"github.com/popeskul/quizer/internal/api/handlers"
	"github.com/popeskul/quizer/internal/api/middleware"
	"github.com/popeskul/quizer/internal/config"
	"github.com/popeskul/quizer/internal/core/services"
	"github.com/popeskul/quizer/internal/core/usecases"
	"github.com/popeskul/quizer/internal/infrastructure/repository"
)

func main() {
	cfg, err := config.LoadConfig("./configs/config.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	repo := repository.NewInMemoryRepository()

	// Set default quiz
	ctx := context.Background()
	if err := repo.QuizRepository().SetQuiz(ctx, repository.CreateDefaultQuiz()); err != nil {
		log.Fatalf("Failed to set default quiz: %v", err)
	}

	service := services.NewServices(repo.QuizRepository())
	useCases := usecases.NewUseCases(service.QuizService())
	handler := handlers.NewHandlers(useCases)

	server := api.NewServer(handler)
	router := server.Handler()

	srv := &http.Server{
		Addr:         ":" + cfg.Server.Port,
		Handler:      middleware.ErrorHandler(router),
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
		IdleTimeout:  cfg.Server.IdleTimeout,
	}

	log.Printf("Server initialized. Listening on port %s", cfg.Server.Port)

	errChan := make(chan error, 1)

	go func() {
		log.Printf("Starting server on :%s", cfg.Server.Port)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errChan <- err
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	select {
	case <-stop:
		log.Println("Shutting down server...")
	case err := <-errChan:
		log.Printf("Error starting server: %s\n", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("Error shutting down server: %s\n", err)
	}

	log.Println("Server stopped")
}
