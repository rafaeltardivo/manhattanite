package serializers

import "fmt"

// ErrSerialize wrapper for file  errors.
func ErrSerialize(msg string) error {
	return fmt.Errorf("error serializing data: %s", msg)
}
