package coll

import "strings"

// Errors to handle multiple error
type Errors []error

// Append error
func (e *Errors) Append(errs ...error) Errors {
	*e = append(*e, errs...)
	return *e
}

func (e Errors) Error() string {
	var builder strings.Builder
	for i, err := range e {
		if i > 0 {
			builder.WriteString("; ")
		}
		builder.WriteString(err.Error())
	}
	return builder.String()
}

// ToError convert errors to error type
func (e Errors) ToError() error {
	if len(e) > 0 {
		return e
	}
	return nil
}
