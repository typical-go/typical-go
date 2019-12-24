package common

import (
	"errors"
	"strings"
)

// Errors to handle multiple error
type Errors []error

// Append error
func (e *Errors) Append(errs ...error) *Errors {
	*e = append(*e, errs...)
	return e
}

// Join list of item to string
func (e *Errors) Join(sep string) string {
	var b strings.Builder
	for i, err := range *e {
		if i > 0 {
			b.WriteString(sep)
		}
		b.WriteString(err.Error())
	}
	return b.String()
}

// Unwrap to error type https://blog.golang.org/go1.13-errors#TOC_3.1.
func (e *Errors) Unwrap() error {
	if len(*e) < 1 {
		return nil
	}
	return errors.New(e.Join("; "))
}

// Slice of error
func (e *Errors) Slice() []error {
	return *e
}
