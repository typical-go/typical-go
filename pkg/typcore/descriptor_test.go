package typcore_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/common"
	"github.com/typical-go/typical-go/pkg/typapp"
	"github.com/typical-go/typical-go/pkg/typbuildtool"
	"github.com/typical-go/typical-go/pkg/typcore"
)

func TestDescriptor(t *testing.T) {
	t.Run("SHOULD implement AppLauncher", func(t *testing.T) {
		var _ typcore.AppLauncher = &typcore.Descriptor{}
	})
	t.Run("SHOULD implement BuildToolLauncher", func(t *testing.T) {
		var _ typcore.BuildToolLauncher = &typcore.Descriptor{}
	})
}

func TestDescriptor_ValidateName(t *testing.T) {
	t.Run("Valid Names", func(t *testing.T) {
		valids := []string{
			"asdf",
			"Asdf",
			"As_df",
			"as-df",
		}
		for _, name := range valids {
			d := &typcore.Descriptor{
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
			d := &typcore.Descriptor{
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
		*typcore.Descriptor
		expectedErr string
	}{
		{
			Descriptor: validDescriptor,
		},
		{
			Descriptor: &typcore.Descriptor{
				Name:      "Typical Go",
				App:       typapp.Create(nil),
				BuildTool: typbuildtool.Create(),
			},
			expectedErr: "Descriptor: Invalid name",
		},
		{
			Descriptor: &typcore.Descriptor{
				Name:      "some-name",
				App:       typapp.Create(nil),
				BuildTool: dummyBuildTool{errMessage: "Build: some-error"},
			},
			expectedErr: "Descriptor: Build: some-error",
		},
		{
			Descriptor: &typcore.Descriptor{
				Name:      "some-name",
				App:       dummyApp{errMessage: "App: some-error"},
				BuildTool: dummyBuildTool{},
			},
			expectedErr: "Descriptor: App: some-error",
		},
		{
			Descriptor: &typcore.Descriptor{
				Name:      "some-name",
				BuildTool: dummyBuildTool{},
			},
			expectedErr: "Descriptor: App can't be nil",
		},
		{
			Descriptor: &typcore.Descriptor{
				Name: "some-name",
				App:  dummyApp{},
			},
			expectedErr: "Descriptor: BuildTool can't be nil",
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
	validDescriptor = &typcore.Descriptor{
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

func (i dummyBuildTool) RunBuildTool(*typcore.Context) error {
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

func (i dummyApp) RunApp(*typcore.Descriptor) error {
	return nil
}

func (i dummyApp) Sources() []string {
	return i.sources
}
