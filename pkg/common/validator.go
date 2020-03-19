package common

import "errors"

// Validator responsible to validate object
type Validator interface {
	Validate() error
}

// Validate obj
func Validate(obj interface{}) error {
	if validator, ok := obj.(Validator); ok {
		return validator.Validate()
	}
	return nil
}

type dummyValidator struct {
	errMsg string
}

// DummyValidator return new instance of validator for test purpose
func DummyValidator(errMsg string) Validator {
	return &dummyValidator{errMsg: errMsg}
}

// Validate return error
func (v *dummyValidator) Validate() error {
	if v.errMsg == "" {
		return nil
	}
	return errors.New(v.errMsg)
}
