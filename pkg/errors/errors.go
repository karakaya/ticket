package errors

import (
	"encoding/json"
	"errors"
	"net/http"
)

// from https://www.joeshaw.org/error-handling-in-go-http-applications/
// good article btw
type ticketAPIError struct {
	status int
	msg    string
}

var (
	ErrAuth       = &ticketAPIError{status: http.StatusUnauthorized, msg: "invalid token"}
	ErrNotFound   = &ticketAPIError{status: http.StatusNotFound, msg: "not found"}
	ErrUUDInvalid = &ticketAPIError{status: http.StatusBadRequest, msg: "invalid UUID"}
	ErrDuplicate  = &ticketAPIError{status: http.StatusBadRequest, msg: "duplicate"}
)

func (t ticketAPIError) Error() string {
	return t.msg
}

func (t ticketAPIError) APIError() (int, string) {
	return t.status, t.msg
}

type APIError interface {
	// APIError returns an HTTP status code and an API-safe error message.
	APIError() (int, string)
}

func JSONHandleError(w http.ResponseWriter, err error) {
	var apiErr APIError

	if errors.As(err, &apiErr) {
		status, msg := apiErr.APIError()
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(map[string]string{"error": msg})

	} else {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "internal server error"})
	}
}
