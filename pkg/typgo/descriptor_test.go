package typgo_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/common"
	"github.com/typical-go/typical-go/pkg/typapp"
	"github.com/typical-go/typical-go/pkg/typgo"
)

func TestDescriptor_ValidateName(t *testing.T) {
	t.Run("Valid Names", func(t *testing.T) {
		valids := []string{
			"asdf",
			"Asdf",
			"As_df",
			"as-df",
		}
		for _, name := range valids {
			d := &typgo.Descriptor{
				Name:      name,
				App:       dummyApp{},
				BuildTool: dummyBuildTool{},
			}
			require.NoError(t, common.Validate(d))
		}
	})
	t.Run("Invalid Names", func(t *testing.T) {
		invalids := []string{
			"Asdf!",
		}
		for _, name := range invalids {
			d := &typgo.Descriptor{
				Name:      name,
				App:       dummyApp{},
				BuildTool: dummyBuildTool{},
			}
			require.EqualError(t, common.Validate(d), "Descriptor: Invalid name")
		}
	})
}

func TestDecriptor_Validate_ReturnError(t *testing.T) {
	testcases := []struct {
		*typgo.Descriptor
		expectedErr string
	}{
		{
			Descriptor: validDescriptor,
		},
		{
			Descriptor: &typgo.Descriptor{
				Name:      "Typical Go",
				App:       &typapp.App{},
				BuildTool: dummyBuildTool{},
			},
			expectedErr: "Descriptor: Invalid name",
		},
		{
			Descriptor: &typgo.Descriptor{
				Name:      "some-name",
				App:       &typapp.App{},
				BuildTool: dummyBuildTool{errMessage: "some-error"},
			},
			expectedErr: "Descriptor: BuildTool: some-error",
		},
		{
			Descriptor: &typgo.Descriptor{
				Name:      "some-name",
				App:       dummyApp{errMessage: "some-error"},
				BuildTool: dummyBuildTool{},
			},
			expectedErr: "Descriptor: App: some-error",
		},
		{
			Descriptor: &typgo.Descriptor{
				Name:      "some-name",
				BuildTool: dummyBuildTool{},
			},
			expectedErr: "Descriptor: App: nil",
		},
		{
			Descriptor: &typgo.Descriptor{
				Name: "some-name",
				App:  dummyApp{},
			},
			expectedErr: "Descriptor: BuildTool: nil",
		},
	}
	for i, tt := range testcases {
		err := tt.Validate()
		if tt.expectedErr == "" {
			require.NoError(t, err, i)
		} else {
			require.EqualError(t, err, tt.expectedErr, i)
		}
	}
}

var (
	validDescriptor = &typgo.Descriptor{
		Name:      "some-name",
		App:       &dummyApp{},
		BuildTool: &dummyBuildTool{},
	}
)

type dummyBuildTool struct {
	errMessage string
}

func (i dummyBuildTool) Validate() error {
	if i.errMessage != "" {
		return errors.New(i.errMessage)
	}
	return nil
}

func (i dummyBuildTool) Run(*typgo.Descriptor) error {
	return nil
}

type dummyApp struct {
	errMessage string
	sources    []string
}

func (i dummyApp) Validate() error {
	if i.errMessage != "" {
		return errors.New(i.errMessage)
	}
	return nil
}

func (i dummyApp) Run(*typgo.Descriptor) error {
	return nil
}
