package effect

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("mode is not found")

func NewErrNotSupported(text string) error {
	return fmt.Errorf("request feature is not supported: %v", text)
}

func NewErrOutOfRange[T int | uint8 | int64](val, min, max T) error {
	return fmt.Errorf("value %v is out of range %v-%v", val, min, max)
}
