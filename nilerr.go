package nilerr

import (
	"errors"
	"fmt"
)

// E is an instance of a nil error that can be used in errors.Is() checks.
var E = nilErr{}

// Message defines the unformatted string used as the error message from nilerr.
var Message = "%v is nil"

// New creates a new nilerr error.
func New(name string) error {
	return nilErr{Name: name}
}

// Is checks if err is a nilerr type ignore the message content.
func Is(err error) bool {
	return errors.Is(err, E)
}

type nilErr struct {
	Name string
}

func (e nilErr) Error() string {
	return fmt.Sprintf("%v is nil", e.Name)
}

func (nilErr) Is(err error) bool {
	_, ok := err.(nilErr)
	return ok
}
