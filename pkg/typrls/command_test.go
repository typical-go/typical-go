package typrls_test

import (
	"errors"
	"flag"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/execkit"
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/typical-go/typical-go/pkg/typrls"
	"github.com/urfave/cli/v2"
)

func TestCommand_Execute(t *testing.T) {
	testcases := []struct {
		TestName        string
		*typrls.Command // NOTE: don't set releaser
		Ctx             *typgo.Context
		RunExpectations []*execkit.RunExpectation
		Expected        *typrls.Context
		ExpectedErr     string
	}{
		{
			TestName: "missing summary",
			Command:  &typrls.Command{},
			Ctx: &typgo.Context{
				Context:    createContext(),
				Descriptor: &typgo.Descriptor{},
			},
			ExpectedErr: "typrls: missing summary",
		},
		{
			TestName: "bad summary",
			Command: &typrls.Command{
				Summary: typrls.NewSummarizer(func(*typrls.Context) (string, error) {
					return "", errors.New("bad-summary")
				}),
			},
			Ctx: &typgo.Context{
				Context:    createContext(),
				Descriptor: &typgo.Descriptor{},
			},
			ExpectedErr: "bad-summary",
		},
		{
			TestName: "invalid",
			Command: &typrls.Command{
				Summary: typrls.DefaultSummary,
				Validation: typrls.NewValidator(func(*typrls.Context) error {
					return errors.New("some-error")
				}),
			},
			Ctx: &typgo.Context{
				Context:    createContext(),
				Descriptor: &typgo.Descriptor{},
			},
			ExpectedErr: "some-error",
		},
		{
			TestName: "with alpha tag",
			Command: &typrls.Command{
				Summary: typrls.NewSummarizer(func(*typrls.Context) (string, error) {
					return "some-summary", nil
				}),
			},
			Ctx: &typgo.Context{
				Context:    createContext("-alpha"),
				Descriptor: &typgo.Descriptor{},
			},
			RunExpectations: []*execkit.RunExpectation{
				{CommandLine: []string{"git", "fetch"}},
				{CommandLine: []string{"git", "describe", "--tags", "--abbrev=0"}, OutputBytes: []byte("some-tag")},
				{CommandLine: []string{"git", "status", "--porcelain"}, OutputBytes: []byte("some-status")},
				{CommandLine: []string{"git", "--no-pager", "log", "some-tag..HEAD", "--oneline"}, OutputBytes: []byte("5378feb one\n")},
			},
			Expected: &typrls.Context{
				ReleaseTag: "v0.0.1_alpha",
				Alpha:      true,
				Summary:    "some-summary",
				Git: &typrls.Git{
					Status:     "some-status",
					CurrentTag: "some-tag",
					Logs: []*typrls.Log{
						{ShortCode: "5378feb", Message: "one"},
					},
				},
			},
		},
		{
			TestName: "success",
			Command: &typrls.Command{
				Summary: typrls.NewSummarizer(func(*typrls.Context) (string, error) {
					return "some-summary", nil
				}),
			},
			Ctx: &typgo.Context{
				Context: createContext(),
				Descriptor: &typgo.Descriptor{
					Version: "9.9.9",
				},
			},
			RunExpectations: []*execkit.RunExpectation{
				{CommandLine: []string{"git", "fetch"}},
				{CommandLine: []string{"git", "describe", "--tags", "--abbrev=0"}, OutputBytes: []byte("some-tag-1")},
				{CommandLine: []string{"git", "status", "--porcelain"}, OutputBytes: []byte("some-status-1")},
			},
			Expected: &typrls.Context{
				ReleaseTag: "v9.9.9",
				Alpha:      false,
				Summary:    "some-summary",
				Git:        &typrls.Git{Status: "some-status-1", CurrentTag: "some-tag-1"},
			},
		},
		{
			TestName: "override release-tag",
			Command: &typrls.Command{
				Summary: typrls.NewSummarizer(func(*typrls.Context) (string, error) {
					return "some-summary", nil
				}),
			},
			Ctx: &typgo.Context{
				Context: createContext("-tag=some-tag"),
				Descriptor: &typgo.Descriptor{
					Version: "9.9.9",
				},
			},
			RunExpectations: []*execkit.RunExpectation{
				{CommandLine: []string{"git", "fetch"}},
				{CommandLine: []string{"git", "describe", "--tags", "--abbrev=0"}, OutputBytes: []byte("some-tag-3")},
				{CommandLine: []string{"git", "status", "--porcelain"}, OutputBytes: []byte("some-status-3")},
			},
			Expected: &typrls.Context{
				ReleaseTag: "some-tag",
				Alpha:      false,
				Summary:    "some-summary",
				Git:        &typrls.Git{Status: "some-status-3", CurrentTag: "some-tag-3"},
			},
		},
	}
	for _, tt := range testcases {
		t.Run(tt.TestName, func(t *testing.T) {
			unpatch := execkit.Patch(tt.RunExpectations)
			defer unpatch(t)

			var rlsCtx *typrls.Context
			tt.Releaser = typrls.NewReleaser(func(r *typrls.Context) error {
				rlsCtx = r
				return nil
			})

			err := tt.Execute(tt.Ctx)
			if tt.ExpectedErr != "" {
				require.EqualError(t, err, tt.ExpectedErr)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.Expected.ReleaseTag, rlsCtx.ReleaseTag)
				require.Equal(t, tt.Expected.Alpha, rlsCtx.Alpha)
				require.Equal(t, tt.Expected.Summary, rlsCtx.Summary)
				require.Equal(t, tt.Expected.Git, rlsCtx.Git)
			}
		})
	}
}

func createContext(args ...string) *cli.Context {
	flagSet := flag.NewFlagSet("test", 0)
	flagSet.Bool(typrls.AlphaFlag, false, "")
	flagSet.Bool(typrls.ForceFlag, false, "")
	flagSet.String(typrls.TagFlag, "", "")
	flagSet.Parse(args)
	return cli.NewContext(nil, flagSet, nil)
}
