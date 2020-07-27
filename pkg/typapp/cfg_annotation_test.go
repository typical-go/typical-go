package typapp_test

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/execkit"
	"github.com/typical-go/typical-go/pkg/typannot"
	"github.com/typical-go/typical-go/pkg/typapp"
	"github.com/typical-go/typical-go/pkg/typgo"
)

func TestCfgAnnotation_Annotate(t *testing.T) {
	target := "some-target"
	unpatch := execkit.Patch([]*execkit.RunExpectation{})
	defer unpatch(t)
	defer os.Clearenv()
	defer os.Remove(target)
	defer os.Remove(".env")

	cfgAnnotation := &typapp.CfgAnnotation{Target: target, DotEnv: true}
	c := &typannot.Context{
		Context: &typgo.Context{
			BuildSys: &typgo.BuildSys{
				Descriptor: &typgo.Descriptor{Name: "some-project"},
			},
		},
		ASTStore: &typannot.ASTStore{
			Annots: []*typannot.Annot{
				{
					TagName:  "@cfg",
					TagParam: `ctor_name:"ctor1" prefix:"SS"`,
					Decl: &typannot.Decl{
						Name:    "SomeSample",
						Package: "mypkg",
						Type: &typannot.StructType{
							Fields: []*typannot.Field{
								{Name: "SomeField1", Type: "string", StructTag: `default:"some-text"`},
								{Name: "SomeField2", Type: "int", StructTag: `default:"9876"`},
							},
						},
					},
				},
			},
		},
	}

	require.NoError(t, cfgAnnotation.Annotate(c))

	b, _ := ioutil.ReadFile(target)
	require.Equal(t, `package main

// Autogenerated by Typical-Go. DO NOT EDIT.

import (
	"github.com/kelseyhightower/envconfig"
)

func init() { 
	typapp.AppendCtor(
		&typapp.Constructor{
			Name: "ctor1",
			Fn: func() (*mypkg.SomeSample, error) {
				var cfg mypkg.SomeSample
				if err := envconfig.Process("SS", &cfg); err != nil {
					return nil, err
				}
				return &cfg, nil
			},
		},
	)
}`, string(b))

	b, _ = ioutil.ReadFile(".env")
	require.Equal(t, "SS_SOMEFIELD1=some-text\nSS_SOMEFIELD2=9876\n", string(b))
	require.Equal(t, "some-text", os.Getenv("SS_SOMEFIELD1"))
	require.Equal(t, "9876", os.Getenv("SS_SOMEFIELD2"))
}

func TestCfgAnnotation_Annotate_Default(t *testing.T) {
	os.MkdirAll("cmd/some-project", 0777)

	unpatch := execkit.Patch([]*execkit.RunExpectation{})
	defer unpatch(t)
	defer os.RemoveAll("cmd/some-project")

	cfgAnnotation := &typapp.CfgAnnotation{}
	c := &typannot.Context{
		Context: &typgo.Context{
			BuildSys: &typgo.BuildSys{
				Descriptor: &typgo.Descriptor{Name: "some-project"},
			},
		},
		ASTStore: &typannot.ASTStore{
			Annots: []*typannot.Annot{
				{
					TagName: "@cfg",
					Decl: &typannot.Decl{
						Name:    "SomeSample",
						Package: "mypkg",
						Type: &typannot.StructType{
							Fields: []*typannot.Field{
								{Name: "SomeField1", Type: "string", StructTag: `default:"some-text"`},
								{Name: "SomeField2", Type: "int", StructTag: `default:"9876"`},
							},
						},
					},
				},
			},
		},
	}

	require.NoError(t, cfgAnnotation.Annotate(c))

	b, _ := ioutil.ReadFile("cmd/some-project/config_annotated.go")

	require.Equal(t, `package main

// Autogenerated by Typical-Go. DO NOT EDIT.

import (
	"github.com/kelseyhightower/envconfig"
)

func init() { 
	typapp.AppendCtor(
		&typapp.Constructor{
			Name: "",
			Fn: func() (*mypkg.SomeSample, error) {
				var cfg mypkg.SomeSample
				if err := envconfig.Process("SOMESAMPLE", &cfg); err != nil {
					return nil, err
				}
				return &cfg, nil
			},
		},
	)
}`, string(b))

}

func TestCreateAndLoadDotEnv_EnvFileExist(t *testing.T) {
	target := "some-env"
	ioutil.WriteFile(target, []byte("key1=val111\nkey2=val222"), 0777)
	var out strings.Builder
	typapp.Stdout = &out
	defer os.Remove(target)
	defer func() { typapp.Stdout = os.Stdout }()

	typapp.CreateAndLoadDotEnv(target, []*typapp.Config{
		{
			Fields: []*typapp.Field{
				{Key: "key1", Default: "val1"},
				{Key: "key2", Default: "val2"},
				{Key: "key3", Default: "val3"},
			},
		},
	})

	b, _ := ioutil.ReadFile(target)
	require.Equal(t, "key1=val111\nkey2=val222\nkey3=val3\n", string(b))
	require.Equal(t, "UPDATE_ENV: +key3\n", out.String())
}