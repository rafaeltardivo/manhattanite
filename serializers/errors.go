package serializers

import "fmt"

// Error wrapper for file  errors.
func ErrSerialize(msg string) error {
	return fmt.Errorf("Error serializing %s", msg)
}
