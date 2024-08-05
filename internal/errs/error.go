package errs

import "fmt"

// NewErrIndexOutOfRange creates a new ErrIndexOutOfRange error
func NewErrIndexOutOfRange(length int, index int) error {
	return fmt.Errorf("index out of range: length %d, index %d", length, index)
}
