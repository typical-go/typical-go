package typcore_test

import (
	"context"
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

func (i invalidBuild) Prebuild(ctx context.Context, b *typcore.BuildContext) error {
	return nil
}

func (i invalidBuild) BuildCommands(b *typcore.BuildContext) []*cli.Command {
	return nil
}

func (i invalidBuild) Releaser() typcore.Releaser {
	return nil
}

type invalidApp struct {
	errMessage string
}

func (i invalidApp) Validate() error {
	return errors.New(i.errMessage)
}

func (i invalidApp) EntryPoint() interface{} {
	return nil
}

func (i invalidApp) Provide() []interface{} {
	return nil
}

func (i invalidApp) Prepare() []interface{} {
	return nil
}

func (i invalidApp) Destroy() []interface{} {
	return nil
}

func (i invalidApp) AppCommands(*typcore.AppContext) []*cli.Command {
	return nil
}
