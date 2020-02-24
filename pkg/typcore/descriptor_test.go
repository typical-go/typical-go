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
				App:       app{},
				BuildTool: typbuildtool.New(),
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
				App:       app{},
				BuildTool: typbuildtool.New(),
			}
			require.EqualError(t, common.Validate(d), "Descriptor: Invalid `Name`")
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
				App:       typapp.New(nil),
				BuildTool: typbuildtool.New(),
			},
			expectedErr: "Descriptor: Invalid `Name`",
		},
		{
			Descriptor: &typcore.Descriptor{
				Name:      "some-name",
				App:       typapp.New(nil),
				BuildTool: buildTool{errMessage: "Build: some-error"},
			},
			expectedErr: "Descriptor: Build: some-error",
		},
		{
			Descriptor: &typcore.Descriptor{
				Name:      "some-name",
				App:       app{errMessage: "App: some-error"},
				BuildTool: typbuildtool.New(),
			},
			expectedErr: "Descriptor: App: some-error",
		},
		{
			Descriptor: &typcore.Descriptor{
				Name:      "some-name",
				BuildTool: typbuildtool.New(),
			},
			expectedErr: "Descriptor: App can't be nil",
		},
		{
			Descriptor: &typcore.Descriptor{
				Name: "some-name",
				App:  app{},
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
		App:       &app{},
		BuildTool: &buildTool{},
	}
)

type buildTool struct {
	errMessage string
}

func (i buildTool) Validate() error {
	if i.errMessage != "" {
		return errors.New(i.errMessage)
	}
	return nil
}

func (i buildTool) Run(*typcore.TypicalContext) error {
	return nil
}

type app struct {
	errMessage string
	sources    []string
}

func (i app) Validate() error {
	if i.errMessage != "" {
		return errors.New(i.errMessage)
	}
	return nil
}

func (i app) Run(*typcore.Descriptor) error {
	return nil
}

func (i app) Sources() []string {
	return i.sources
}
