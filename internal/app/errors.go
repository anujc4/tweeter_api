package app

import (
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

type Error struct {
	code    int
	message string
	err     error
}

// Create a new Error struct
func NewError(err error) *Error {
	return &Error{err: errors.WithStack(err)}
}

// Set the code for the error
// HTTP status codes as registered with IANA.
// See: https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml
func (err *Error) SetCode(code int) *Error {
	err.code = code
	return err
}

// Set the message for the error
func (err *Error) SetMessage(message string) *Error {
	err.message = message
	return err
}

// Getter method for the message
func (err *Error) Message() string {
	return err.message
}

// Getter method for the status code
func (err *Error) Code() int {
	return err.code
}

// Pre-defined errors specific to usecase
func NewValidationError(err error) *Error {
	errJson, _ := json.Marshal(err)
	return &Error{
		code:    http.StatusBadRequest,
		message: string(errJson),
	}
}

func NewParseFormError(err error) *Error {
	errJson, _ := json.Marshal(err)
	return &Error{
		code:    http.StatusBadRequest,
		message: string(errJson),
	}
}
