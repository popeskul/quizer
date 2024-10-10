package handlers

import (
	"math/rand"
	"net/http"
)

type StatusHandler struct{}

func NewStatusHandler() *StatusHandler {
	return &StatusHandler{}
}

func (h *StatusHandler) Success(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (h *StatusHandler) Internal(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
}

func (h *StatusHandler) Random(w http.ResponseWriter, r *http.Request) {
	randomNum := rand.Float32()
	if randomNum < 0.2 {
		panic("Random panic")
	} else if randomNum < 0.6 {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
