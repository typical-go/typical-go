package typcore_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typbuild"
	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/urfave/cli/v2"
)

func TestDescriptor_Validate_DefaultValue(t *testing.T) {
	d := &typcore.Descriptor{
		Name:    "some-name",
		Package: "some-package",
		Build:   typbuild.New(),
	}
	require.True(t, typcore.IsValidator(d))
	require.NoError(t, typcore.Validate(d))
	require.Equal(t, "0.0.1", d.Version)
}

func TestDecriptor_Validate_ReturnError(t *testing.T) {
	testcases := []struct {
		typcore.Descriptor
		errMsg string
	}{
		{
			typcore.Descriptor{Package: "some-package"},
			"Descriptor: Name can't be empty",
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
				Build:   typbuild.New(),
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

func (i invalidBuild) Invoke(bctx *typcore.BuildContext, c *cli.Context, fn interface{}) (err error) {
	return nil
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

func (i invalidApp) Invoke(actx *typcore.AppContext, c *cli.Context, fn interface{}) (err error) {
	return nil
}

func (i invalidApp) Run(*typcore.AppContext) error {
	return nil
}
