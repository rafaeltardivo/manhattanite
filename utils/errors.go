package utils

import "fmt"

// Error wrapper for not found errors.
func ErrNotFound(msg string, name string) error {
	return fmt.Errorf("%s not found: %s", msg, name)
}

// Error wrapper for file  errors.
func ErrFile(msg string) error {
	return fmt.Errorf("Error processing file: %s", msg)
}

// Error wrapper for invalid query errors.
func ErrInvalidQuery(msg string) error {
	return fmt.Errorf("Error parsing query string: %s", msg)
}

// Error wrapper for invalid request method errors.
func ErrInvalidMethod(msg string) error {
	return fmt.Errorf("Error invalid method: %s", msg)
}

// Error response payload.
type ErrorResponse struct {
	Msg string `json:"error"`
}
