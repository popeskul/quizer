package ports

//go:generate mockgen -destination=repositories_mock.go -package=ports github.com/popeskul/quizer/internal/core/ports Repository,QuizRepository
//go:generate mockgen -destination=services_mock.go -package=ports github.com/popeskul/quizer/internal/core/ports Services,QuizService
//go:generate mockgen -destination=usecases_mock.go -package=ports github.com/popeskul/quizer/internal/core/ports UseCases,QuizUseCase
