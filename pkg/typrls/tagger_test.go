package typrls_test

import (
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/typical-go/typical-go/pkg/oskit"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/execkit"
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/typical-go/typical-go/pkg/typrls"
	"github.com/urfave/cli/v2"
)

func TestTagger(t *testing.T) {
	typrls.Now = func() time.Time {
		return time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	}
	defer func() { typrls.Now = time.Now }()
	testcases := []struct {
		TestName string
		typrls.Tagger
		Context         *typgo.Context
		Alpha           bool
		RunExpectations []*execkit.RunExpectation
		Expected        string
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
			TestName: "include git id and date",
			Tagger:   (&typrls.StdTagger{}).IncludeGitID().IncludeDate(),
			Context: &typgo.Context{
				BuildSys: &typgo.BuildSys{
					Descriptor: &typgo.Descriptor{ProjectVersion: "0.0.1"},
				},
				Context: &cli.Context{},
			},
			RunExpectations: []*execkit.RunExpectation{
				{CommandLine: "git rev-parse HEAD", OutputBytes: []byte("1234567890")},
			},
			Expected: "v0.0.1+123456_20091110",
		},
		{
			TestName: "with git id error",
			Tagger:   (&typrls.StdTagger{}).IncludeGitID(),
			Context: &typgo.Context{
				BuildSys: &typgo.BuildSys{
					Descriptor: &typgo.Descriptor{ProjectVersion: "0.0.1"},
				},
				Context: &cli.Context{},
			},
			RunExpectations: []*execkit.RunExpectation{
				{CommandLine: "git rev-parse HEAD", ReturnError: errors.New("some-error")},
			},
			Expected: "v0.0.1+",
		},
	}
	for _, tt := range testcases {
		t.Run(tt.TestName, func(t *testing.T) {
			var out strings.Builder
			defer oskit.PatchStdout(&out)()
			defer execkit.Patch(tt.RunExpectations)(t)

			require.Equal(t, tt.Expected, tt.CreateTag(tt.Context, tt.Alpha))
		})
	}
}
