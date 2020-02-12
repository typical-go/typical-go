package typcore_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/common"
	"github.com/typical-go/typical-go/pkg/typcore"
)

func TestDescriptor_Validate_DefaultValue(t *testing.T) {
	d := &typcore.Descriptor{
		Package: "some-package",
	}
	require.NoError(t, common.Validate(d))

	require.Equal(t, "typcore", d.Name)
	require.Equal(t, "0.0.1", d.Version)
	require.Equal(t, []string{"app"}, d.Sources)
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
			require.NoError(t, common.Validate(&typcore.Descriptor{Name: name, Package: "some-package"}))
		}
	})
	t.Run("Invalid Names", func(t *testing.T) {
		invalids := []string{
			"Asdf!",
		}
		for _, name := range invalids {
			require.EqualError(t,
				common.Validate(&typcore.Descriptor{Name: name, Package: "some-package"}),
				"Descriptor: Invalid `Name`",
			)
		}
	})
}

func TestDecriptor_Validate_ReturnError(t *testing.T) {
	testcases := []struct {
		typcore.Descriptor
		errMsg string
	}{
		{
			typcore.Descriptor{Name: "Asdf", Package: "some-package"},
			"Descriptor: Name: 'Asdf' must be lower case and contain aplabhet or underscore ('_')",
		},
		{
			typcore.Descriptor{Name: "some-name"},
			"Descriptor: Package can't be empty",
		},
		{
			typcore.Descriptor{
				Name:    "some-name",
				Package: "some-package",
				Build:   invalidBuild{"Build: some-error"},
			},
			"Descriptor: Build: some-error",
		},
		{
			typcore.Descriptor{
				Name:    "some-name",
				Package: "some-package",
				App:     invalidApp{"App: some-error"},
			},
			"Descriptor: App: some-error",
		},
	}
	for i, tt := range testcases {
		err := tt.Validate()
		if tt.errMsg == "" {
			require.NoError(t, err, i)
		} else {
			require.EqualError(t, err, tt.errMsg, i)
		}
	}
}

type invalidBuild struct {
	errMessage string
}

func (i invalidBuild) Validate() error {
	return errors.New(i.errMessage)
}

func (i invalidBuild) Run(*typcore.BuildContext) error {
	return nil
}

type invalidApp struct {
	errMessage string
}

func (i invalidApp) Validate() error {
	return errors.New(i.errMessage)
}

func (i invalidApp) Run(*typcore.AppContext) error {
	return nil
}

func (i invalidApp) Sources() []string {
	return nil
}
