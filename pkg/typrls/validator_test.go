package typrls_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typrls"
)

func TestValidator(t *testing.T) {
	testcases := []struct {
		TestName string
		typrls.Validator
		Context     *typrls.Context
		ExpectedErr string
	}{
		{
			Validator: typrls.NewValidator(func(*typrls.Context) error {
				return errors.New("some-error")
			}),
			ExpectedErr: "some-error",
		},
		{
			TestName: "composite validation: later error",
			Validator: typrls.Validators{
				typrls.NewValidator(func(*typrls.Context) error { return nil }),
				typrls.NewValidator(func(*typrls.Context) error { return errors.New("some-error-2") }),
			},
			ExpectedErr: "some-error-2",
		},
		{
			TestName: "composite validation: first error",
			Validator: typrls.Validators{
				typrls.NewValidator(func(*typrls.Context) error { return errors.New("some-error-1") }),
				typrls.NewValidator(func(*typrls.Context) error { return errors.New("some-error-2") }),
			},
			ExpectedErr: "some-error-1",
		},
		{
			TestName:  "uncommitted change",
			Validator: &typrls.UncommittedValidation{},
			Context: &typrls.Context{
				Git: &typrls.Git{
					Status: "some-status",
				},
			},
			ExpectedErr: "Please commit changes first:\nsome-status",
		},
		{
			TestName:  "no uncommitted change",
			Validator: &typrls.UncommittedValidation{},
			Context:   &typrls.Context{Git: &typrls.Git{}},
		},
		{
			TestName:  "already release",
			Validator: &typrls.AlreadyReleasedValidation{},
			Context: &typrls.Context{
				ReleaseTag: "some-tag",
				Git: &typrls.Git{
					CurrentTag: "some-tag",
				},
			},
			ExpectedErr: "some-tag already released",
		},
		{
			TestName:  "not release yet",
			Validator: &typrls.AlreadyReleasedValidation{},
			Context: &typrls.Context{
				ReleaseTag: "different-tag",
				Git: &typrls.Git{
					CurrentTag: "some-tag",
				},
			},
		},
		{
			TestName:    "not git change since last release",
			Validator:   &typrls.NoGitChangeValidation{},
			Context:     &typrls.Context{Git: &typrls.Git{}},
			ExpectedErr: "No change to be released",
		},
		{
			TestName:  "there are git change since last release",
			Validator: &typrls.NoGitChangeValidation{},
			Context: &typrls.Context{
				Git: &typrls.Git{
					Logs: []*typrls.Log{
						{ShortCode: "some-short", Message: "some-message"},
					},
				},
			},
		},
	}
	for _, tt := range testcases {
		t.Run(tt.TestName, func(t *testing.T) {
			err := tt.Validate(tt.Context)
			if tt.ExpectedErr != "" {
				require.EqualError(t, err, tt.ExpectedErr)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
