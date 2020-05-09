package common

import "errors"

// Validator responsible to validate object
type Validator interface {
	Validate() error
}

// Validate obj
func Validate(obj interface{}) error {
	if obj == nil {
		return errors.New("nil")
	}
	if validator, ok := obj.(Validator); ok {
		return validator.Validate()
	}
	return nil
}
