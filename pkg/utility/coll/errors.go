package coll

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
func (e Errors) Join(sep string) string {
	var b strings.Builder
	for i, err := range e {
		if i > 0 {
			b.WriteString(sep)
		}
		b.WriteString(err.Error())
	}
	return b.String()
}

// ToError convert errors to error type
func (e Errors) ToError() error {
	if len(e) < 1 {
		return nil
	}
	return errors.New(e.Join("; "))
}
