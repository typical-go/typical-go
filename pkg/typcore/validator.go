package typcore

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

// IsValidator return true is object
func IsValidator(obj interface{}) bool {
	_, ok := obj.(Validator)
	return ok
}
