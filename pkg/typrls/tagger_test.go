package typrls_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/execkit"
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/typical-go/typical-go/pkg/typrls"
	"github.com/urfave/cli/v2"
)

func TestTagger(t *testing.T) {
	testcases := []struct {
		TestName string
		typrls.Tagger
		Context      *typgo.Context
		Alpha        bool
		Expectations []*execkit.RunExpectation
		Expected     string
	}{
		{
			Tagger: &typrls.StdTagger{},
			Context: &typgo.Context{
				BuildSys: &typgo.BuildSys{
					Descriptor: &typgo.Descriptor{ProjectVersion: "0.0.1"},
				},
			},
			Expected: "v0.0.1",
		},
		{
			TestName: "with alpha",
			Tagger:   &typrls.StdTagger{},
			Context: &typgo.Context{
				BuildSys: &typgo.BuildSys{
					Descriptor: &typgo.Descriptor{ProjectVersion: "0.0.1"},
				},
			},
			Alpha:    true,
			Expected: "v0.0.1_alpha",
		},
		{
			TestName: "with git id",
			Tagger: &typrls.StdTagger{
				GitID: true,
			},
			Context: &typgo.Context{
				BuildSys: &typgo.BuildSys{
					Descriptor: &typgo.Descriptor{ProjectVersion: "0.0.1"},
				},
				Context: &cli.Context{},
			},
			Expectations: []*execkit.RunExpectation{
				{CommandLine: "git rev-parse HEAD", OutputBytes: []byte("1234567890")},
			},
			Expected: "v0.0.1+123456",
		},
		{
			TestName: "with git id error",
			Tagger: &typrls.StdTagger{
				GitID: true,
			},
			Context: &typgo.Context{
				BuildSys: &typgo.BuildSys{
					Descriptor: &typgo.Descriptor{ProjectVersion: "0.0.1"},
				},
				Context: &cli.Context{},
			},
			Expected: "v0.0.1+",
		},
	}
	for _, tt := range testcases {
		t.Run(tt.TestName, func(t *testing.T) {
			unpatch := execkit.Patch(tt.Expectations)
			defer unpatch(t)
			require.Equal(t, tt.Expected, tt.CreateTag(tt.Context, tt.Alpha))
		})
	}
}

func TestStdTagger(t *testing.T) {
	tag := &typrls.StdTagger{}
	tag.WithGitID()
	require.True(t, tag.GitID)
}
