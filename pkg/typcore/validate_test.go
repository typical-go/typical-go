package typcore_test

import (
	"errors"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/common"
	"github.com/typical-go/typical-go/pkg/typapp"
	"github.com/typical-go/typical-go/pkg/typbuildtool"
	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/typical-go/typical-go/typicalgo"
)

func TestDescriptor_Validate_DefaultValue(t *testing.T) {
	os.Mkdir("typicalgo", os.ModePerm)
	os.Mkdir("pkg", os.ModePerm)
	ioutil.WriteFile("go.mod", []byte("module github.com/typical-go/typical-go\ngo 1.13"), 0644)
	defer func() {
		os.Remove("typicalgo")
		os.Remove("pkg")
		os.Remove("go.mod")
	}()

	d := &typcore.Descriptor{
		App:       typicalgo.New(),
		BuildTool: typbuildtool.New(),
	}
	require.NoError(t, common.Validate(d))

	require.Equal(t, "typcore", d.Name)
	require.Equal(t, "0.0.1", d.Version)
	require.EqualValues(t, []string{"typicalgo", "pkg"}, d.ProjectSources)
	require.EqualValues(t, "github.com/typical-go/typical-go", d.ModulePackage)
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
				Name:          name,
				ModulePackage: "some-package",
				App:           app{},
				BuildTool:     typbuildtool.New(),
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
				Name:          name,
				ModulePackage: "some-package",
				App:           app{},
				BuildTool:     typbuildtool.New(),
			}
			require.EqualError(t, common.Validate(d), "Descriptor: Invalid `Name`")
		}
	})
}

func TestDecriptor_Validate_ReturnError(t *testing.T) {
	testcases := []struct {
		typcore.Descriptor
		errMsg string
	}{
		{
			typcore.Descriptor{
				Name:          "Typical Go",
				ModulePackage: "some-package",
				App:           typapp.New(nil),
				BuildTool:     typbuildtool.New(),
			},
			"Descriptor: Invalid `Name`",
		},
		{
			typcore.Descriptor{
				Name:      "some-name",
				App:       typapp.New(nil),
				BuildTool: typbuildtool.New(),
			},
			"`go.mod` is missing and the project not in $GOPATH",
		},
		{
			typcore.Descriptor{
				Name:          "some-name",
				ModulePackage: "some-package",
				App:           typapp.New(nil),
				BuildTool:     invalidBuildTool{"Build: some-error"},
			},
			"Descriptor: Build: some-error",
		},
		{
			typcore.Descriptor{
				Name:          "some-name",
				ModulePackage: "some-package",
				App:           invalidApp{"App: some-error"},
				BuildTool:     typbuildtool.New(),
			},
			"Descriptor: App: some-error",
		},
		{
			typcore.Descriptor{
				Name:          "some-name",
				ModulePackage: "some-package",
				App:           app{sources: []string{"bad-src"}},
				BuildTool:     typbuildtool.New(),
			},
			"Descriptor: Source 'bad-src' is not exist",
		},
		{
			typcore.Descriptor{
				Name:          "some-name",
				ModulePackage: "some-package",
				BuildTool:     typbuildtool.New(),
			},
			"Descriptor: App can't be nil",
		},
		{
			typcore.Descriptor{
				Name:          "some-name",
				ModulePackage: "some-package",
				App:           app{},
			},
			"Descriptor: BuildTool can't be nil",
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

type invalidBuildTool struct {
	errMessage string
}

func (i invalidBuildTool) Validate() error {
	return errors.New(i.errMessage)
}

func (i invalidBuildTool) Run(*typcore.TypicalContext) error {
	return nil
}

type invalidApp struct {
	errMessage string
}

func (i invalidApp) Validate() error {
	return errors.New(i.errMessage)
}

func (i invalidApp) Run(*typcore.Descriptor) error {
	return nil
}

func (i invalidApp) Sources() []string {
	return nil
}

type app struct {
	sources []string
}

func (a app) Run(*typcore.Descriptor) error {
	return nil
}

func (a app) Sources() []string {
	return a.sources
}
