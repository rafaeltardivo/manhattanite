package utils

import "fmt"

// ErrNotFound wrapper for not found errors.
func ErrNotFound(msg string, name string) error {
	return fmt.Errorf("%s not found: %s", msg, name)
}

// ErrFile wrapper for file  errors.
func ErrFile(msg string) error {
	return fmt.Errorf("error processing file: %s", msg)
}

// ErrInvalidQuery for invalid query errors.
func ErrInvalidQuery(msg string) error {
	return fmt.Errorf("error parsing query string: %s", msg)
}

// ErrInvalidMethod wrapper for invalid request method errors.
func ErrInvalidMethod(msg string) error {
	return fmt.Errorf("error invalid method: %s", msg)
}

// ErrorResponse response payload.
type ErrorResponse struct {
	Msg string `json:"error"`
}
