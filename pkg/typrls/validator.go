package typrls

import (
	"errors"
	"fmt"
)

type (
	// Validator responsible for validation before release`
	Validator interface {
		Validate(*Context) error
	}
	// Validators for composite validation
	Validators []Validator
	// ValidateFn validate function
	ValidateFn    func(*Context) error
	validatorImpl struct {
		fn ValidateFn
	}
	// NoGitChangeValidation check if no git change since last release
	NoGitChangeValidation struct{}
	// AlreadyReleasedValidation check if project with same tag already released
	AlreadyReleasedValidation struct{}
	// UncommittedValidation check if there are uncommitted files
	UncommittedValidation struct{}
)

//
// validatorImpl
//

// NewValidator return new instance of validator
func NewValidator(fn ValidateFn) Validator {
	return &validatorImpl{fn: fn}
}

func (v *validatorImpl) Validate(c *Context) error {
	return v.fn(c)
}

//
// Validators
//

var _ Validator = (Validators)(nil)

// Validate release
func (v Validators) Validate(c *Context) error {
	for _, validator := range v {
		if err := validator.Validate(c); err != nil {
			return err
		}
	}
	return nil
}

//
// NoGitChangeValidation
//

// Validate if no git change
func (*NoGitChangeValidation) Validate(c *Context) error {
	if len(c.Git.Logs) < 1 {
		return errors.New("No change to be released")
	}
	return nil
}

//
// AlreadyReleasedValidation
//

// Validate if already release
func (*AlreadyReleasedValidation) Validate(c *Context) error {
	if c.TagName == c.Git.CurrentTag && c.Git.CurrentTag != "" {
		return fmt.Errorf("%s already released", c.TagName)
	}
	return nil
}

//
// UncommittedValidation
//

// Validate if uncommitted change
func (*UncommittedValidation) Validate(c *Context) error {
	if c.Git.Status != "" {
		return fmt.Errorf("Please commit changes first:\n%s", c.Git.Status)
	}
	return nil
}
