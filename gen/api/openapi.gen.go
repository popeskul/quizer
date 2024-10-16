// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.3 DO NOT EDIT.
package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	strictnethttp "github.com/oapi-codegen/runtime/strictmiddleware/nethttp"
)

// Answer defines model for Answer.
type Answer struct {
	Id   *string `json:"id,omitempty"`
	Text *string `json:"text,omitempty"`
}

// Question defines model for Question.
type Question struct {
	Answers *[]Answer `json:"answers,omitempty"`
	Id      *string   `json:"id,omitempty"`
	Text    *string   `json:"text,omitempty"`
}

// Quiz defines model for Quiz.
type Quiz struct {
	Questions *[]Question `json:"questions,omitempty"`
}

// QuizAnswers defines model for QuizAnswers.
type QuizAnswers map[string]string

// QuizResult defines model for QuizResult.
type QuizResult struct {
	CorrectAnswers *int     `json:"correctAnswers,omitempty"`
	Percentile     *float32 `json:"percentile,omitempty"`
	TotalQuestions *int     `json:"totalQuestions,omitempty"`
}

// PostSubmitJSONRequestBody defines body for PostSubmit for application/json ContentType.
type PostSubmitJSONRequestBody = QuizAnswers

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Returns 500 Internal Server Error
	// (GET /internal)
	GetInternal(w http.ResponseWriter, r *http.Request)
	// Get a quiz
	// (GET /quiz)
	GetQuiz(w http.ResponseWriter, r *http.Request)
	// Returns random status
	// (POST /random)
	PostRandom(w http.ResponseWriter, r *http.Request)
	// Submit quiz answers
	// (POST /submit)
	PostSubmit(w http.ResponseWriter, r *http.Request)
	// Returns 200 OK
	// (GET /success)
	GetSuccess(w http.ResponseWriter, r *http.Request)
	// Swagger UI
	// (GET /swagger)
	GetSwagger(w http.ResponseWriter, r *http.Request)
}

// Unimplemented server implementation that returns http.StatusNotImplemented for each endpoint.

type Unimplemented struct{}

// Returns 500 Internal Server Error
// (GET /internal)
func (_ Unimplemented) GetInternal(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Get a quiz
// (GET /quiz)
func (_ Unimplemented) GetQuiz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Returns random status
// (POST /random)
func (_ Unimplemented) PostRandom(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Submit quiz answers
// (POST /submit)
func (_ Unimplemented) PostSubmit(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Returns 200 OK
// (GET /success)
func (_ Unimplemented) GetSuccess(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Swagger UI
// (GET /swagger)
func (_ Unimplemented) GetSwagger(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// GetInternal operation middleware
func (siw *ServerInterfaceWrapper) GetInternal(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetInternal(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetQuiz operation middleware
func (siw *ServerInterfaceWrapper) GetQuiz(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetQuiz(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// PostRandom operation middleware
func (siw *ServerInterfaceWrapper) PostRandom(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostRandom(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// PostSubmit operation middleware
func (siw *ServerInterfaceWrapper) PostSubmit(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostSubmit(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetSuccess operation middleware
func (siw *ServerInterfaceWrapper) GetSuccess(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetSuccess(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetSwagger operation middleware
func (siw *ServerInterfaceWrapper) GetSwagger(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetSwagger(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshalingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshalingParamError) Error() string {
	return fmt.Sprintf("Error unmarshaling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshalingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{})
}

type ChiServerOptions struct {
	BaseURL          string
	BaseRouter       chi.Router
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r chi.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options ChiServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = chi.NewRouter()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/internal", wrapper.GetInternal)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/quiz", wrapper.GetQuiz)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/random", wrapper.PostRandom)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/submit", wrapper.PostSubmit)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/success", wrapper.GetSuccess)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/swagger", wrapper.GetSwagger)
	})

	return r
}

type GetInternalRequestObject struct {
}

type GetInternalResponseObject interface {
	VisitGetInternalResponse(w http.ResponseWriter) error
}

type GetInternal500Response struct {
}

func (response GetInternal500Response) VisitGetInternalResponse(w http.ResponseWriter) error {
	w.WriteHeader(500)
	return nil
}

type GetQuizRequestObject struct {
}

type GetQuizResponseObject interface {
	VisitGetQuizResponse(w http.ResponseWriter) error
}

type GetQuiz200JSONResponse Quiz

func (response GetQuiz200JSONResponse) VisitGetQuizResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PostRandomRequestObject struct {
}

type PostRandomResponseObject interface {
	VisitPostRandomResponse(w http.ResponseWriter) error
}

type PostRandom200Response struct {
}

func (response PostRandom200Response) VisitPostRandomResponse(w http.ResponseWriter) error {
	w.WriteHeader(200)
	return nil
}

type PostRandom500Response struct {
}

func (response PostRandom500Response) VisitPostRandomResponse(w http.ResponseWriter) error {
	w.WriteHeader(500)
	return nil
}

type PostSubmitRequestObject struct {
	Body *PostSubmitJSONRequestBody
}

type PostSubmitResponseObject interface {
	VisitPostSubmitResponse(w http.ResponseWriter) error
}

type PostSubmit200JSONResponse QuizResult

func (response PostSubmit200JSONResponse) VisitPostSubmitResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetSuccessRequestObject struct {
}

type GetSuccessResponseObject interface {
	VisitGetSuccessResponse(w http.ResponseWriter) error
}

type GetSuccess200Response struct {
}

func (response GetSuccess200Response) VisitGetSuccessResponse(w http.ResponseWriter) error {
	w.WriteHeader(200)
	return nil
}

type GetSwaggerRequestObject struct {
}

type GetSwaggerResponseObject interface {
	VisitGetSwaggerResponse(w http.ResponseWriter) error
}

type GetSwagger200Response struct {
}

func (response GetSwagger200Response) VisitGetSwaggerResponse(w http.ResponseWriter) error {
	w.WriteHeader(200)
	return nil
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {
	// Returns 500 Internal Server Error
	// (GET /internal)
	GetInternal(ctx context.Context, request GetInternalRequestObject) (GetInternalResponseObject, error)
	// Get a quiz
	// (GET /quiz)
	GetQuiz(ctx context.Context, request GetQuizRequestObject) (GetQuizResponseObject, error)
	// Returns random status
	// (POST /random)
	PostRandom(ctx context.Context, request PostRandomRequestObject) (PostRandomResponseObject, error)
	// Submit quiz answers
	// (POST /submit)
	PostSubmit(ctx context.Context, request PostSubmitRequestObject) (PostSubmitResponseObject, error)
	// Returns 200 OK
	// (GET /success)
	GetSuccess(ctx context.Context, request GetSuccessRequestObject) (GetSuccessResponseObject, error)
	// Swagger UI
	// (GET /swagger)
	GetSwagger(ctx context.Context, request GetSwaggerRequestObject) (GetSwaggerResponseObject, error)
}

type StrictHandlerFunc = strictnethttp.StrictHttpHandlerFunc
type StrictMiddlewareFunc = strictnethttp.StrictHttpMiddlewareFunc

type StrictHTTPServerOptions struct {
	RequestErrorHandlerFunc  func(w http.ResponseWriter, r *http.Request, err error)
	ResponseErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares, options: StrictHTTPServerOptions{
		RequestErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		},
		ResponseErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		},
	}}
}

func NewStrictHandlerWithOptions(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc, options StrictHTTPServerOptions) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares, options: options}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
	options     StrictHTTPServerOptions
}

// GetInternal operation middleware
func (sh *strictHandler) GetInternal(w http.ResponseWriter, r *http.Request) {
	var request GetInternalRequestObject

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.GetInternal(ctx, request.(GetInternalRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetInternal")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(GetInternalResponseObject); ok {
		if err := validResponse.VisitGetInternalResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// GetQuiz operation middleware
func (sh *strictHandler) GetQuiz(w http.ResponseWriter, r *http.Request) {
	var request GetQuizRequestObject

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.GetQuiz(ctx, request.(GetQuizRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetQuiz")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(GetQuizResponseObject); ok {
		if err := validResponse.VisitGetQuizResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// PostRandom operation middleware
func (sh *strictHandler) PostRandom(w http.ResponseWriter, r *http.Request) {
	var request PostRandomRequestObject

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.PostRandom(ctx, request.(PostRandomRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostRandom")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(PostRandomResponseObject); ok {
		if err := validResponse.VisitPostRandomResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// PostSubmit operation middleware
func (sh *strictHandler) PostSubmit(w http.ResponseWriter, r *http.Request) {
	var request PostSubmitRequestObject

	var body PostSubmitJSONRequestBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		sh.options.RequestErrorHandlerFunc(w, r, fmt.Errorf("can't decode JSON body: %w", err))
		return
	}
	request.Body = &body

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.PostSubmit(ctx, request.(PostSubmitRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostSubmit")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(PostSubmitResponseObject); ok {
		if err := validResponse.VisitPostSubmitResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// GetSuccess operation middleware
func (sh *strictHandler) GetSuccess(w http.ResponseWriter, r *http.Request) {
	var request GetSuccessRequestObject

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.GetSuccess(ctx, request.(GetSuccessRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetSuccess")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(GetSuccessResponseObject); ok {
		if err := validResponse.VisitGetSuccessResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// GetSwagger operation middleware
func (sh *strictHandler) GetSwagger(w http.ResponseWriter, r *http.Request) {
	var request GetSwaggerRequestObject

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.GetSwagger(ctx, request.(GetSwaggerRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetSwagger")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(GetSwaggerResponseObject); ok {
		if err := validResponse.VisitGetSwaggerResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}
