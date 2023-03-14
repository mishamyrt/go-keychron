package hid

import (
	"errors"
	"fmt"
)

// ErrPayloadOverflow returns when provided payload is too large
var ErrPayloadOverflow = errors.New("payload is too large. Max length is PacketLength (64)")

// ErrWrongUsage is returned when the device has opened, but the Usage values do not match the expected ones
var ErrWrongUsage = errors.New("wrong device usage")

// NewErrCountMismatch creates a byte count mismatch error.
func NewErrCountMismatch(expected, actual int) error {
	message := fmt.Sprintf(
		"the transmitted number of bytes is not the same as expected. Transmitted %d when it should be %d",
		actual,
		expected,
	)
	return errors.New(message)
}
