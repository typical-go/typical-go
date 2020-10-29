package errkit

import (
	"errors"
	"strings"
)

type (
	// Errors to handle multiple error
	Errors []error
)

// Append error
func (e *Errors) Append(errs ...error) *Errors {
	*e = append(*e, errs...)
	return e
}

// Join list of item to string
func (e *Errors) Join(sep string) string {
	var msgs []string
	for _, err := range *e {
		if err != nil {
			msgs = append(msgs, err.Error())
		}
	}
	return strings.Join(msgs, sep)
}

// Unwrap to error type https://blog.golang.org/go1.13-errors#TOC_3.1.
func (e *Errors) Unwrap() error {
	msg := e.Join("; ")
	if msg != "" {
		return errors.New(msg)
	}
	return nil
}
