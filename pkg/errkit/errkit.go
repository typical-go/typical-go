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
	for _, err := range errs {
		if err != nil {
			*e = append(*e, err)
		}
	}
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
