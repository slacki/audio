package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

// Error represents a handler error. It provides methods for a HTTP status
// code and embeds the built-in error interface.
type Error interface {
	error
	Status() int
}

// StatusError represents an error with an associated HTTP status code.
type StatusError struct {
	Code int
	Err  error
}

// StatusErrorResponse represents error JSON response
type StatusErrorResponse struct {
	Success bool   `json:"success"`
	Status  int    `json:"status"`
	Message string `json:"message,omitempty"`
}

// Error allows StatusError to satisfy the error interface.
func (se StatusError) Error() string {
	return se.Err.Error()
}

// Status returns HTTP status code.
func (se StatusError) Status() int {
	return se.Code
}

// Env represents a (simple) application-wide configuration.
type Env struct {
	UploadPath string
	// DB *sqlx.DB
	// WS *ws.Hub
}

// The Handler struct that takes a configured Env and a function matching
// our useful signature.
type Handler struct {
	*Env
	H func(e *Env, w http.ResponseWriter, r *http.Request) error
}

// ServeHTTP allows our Handler type to satisfy http.Handler.
func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := h.H(h.Env, w, r)
	if err != nil {
		switch e := err.(type) {
		case Error:
			// We can retrieve the status here and write out a specific
			// HTTP status code.
			log.Printf("HTTP %d - %s", e.Status(), e)
			http.Error(w, e.Error(), e.Status())
		default:
			// Any error types we don't specifically look out for default
			// to serving a HTTP 500
			http.Error(w, http.StatusText(http.StatusInternalServerError),
				http.StatusInternalServerError)
		}
	}
}

// HTTPError prepares StatusError with custom message
func HTTPError(status int, message string, original error) StatusError {
	response := StatusErrorResponse{
		Status:  status,
		Message: message,
		Success: false,
	}

	r, err := json.Marshal(response)
	if err != nil {
		log.Printf("Error generating json, %v", err)
	}

	if original != nil {
		log.Printf("--- %s", original.Error())
	}

	return StatusError{
		Code: status,
		Err:  errors.New(string(r)),
	}
}
