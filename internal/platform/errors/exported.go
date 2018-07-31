package errors

import (
	"github.com/pkg/errors"
)

// New returns an error with the supplied message.
// New also records the stack trace at the point it was called.
func New(message string) error {
	return errors.New(message)
}

// Errorf formats according to a format specifier and returns the string
// as a value that satisfies error.
// Errorf also records the stack trace at the point it was called.
func Errorf(format string, args ...interface{}) error {
	return errors.Errorf(format, args)
}

// Wrap wraps an error with a contextual message
func Wrap(err error, message string) error {
	return errors.Wrap(err, message)
}