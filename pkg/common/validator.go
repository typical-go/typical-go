package common

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
