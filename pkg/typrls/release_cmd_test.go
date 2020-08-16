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

func TestReleaseCmd_Execute_NoReleaser(t *testing.T) {
	rls := &typrls.ReleaseProject{
		Tagger:     typrls.DefaultTagger,
		Summarizer: typrls.DefaultSummarizer,
	}
	require.EqualError(t, rls.Execute(nil), "typrls: missing releaser")
}

func TestReleaseCmd_Execute(t *testing.T) {
	testcases := []struct {
		TestName string
		*typrls.ReleaseProject
		Ctx             *typgo.Context
		RunExpectations []*execkit.RunExpectation
		Expected        *typrls.Context
		ExpectedErr     string
	}{
		{
			TestName:       "missing summary",
			ReleaseProject: &typrls.ReleaseProject{},
			Ctx: &typgo.Context{
				Context: createContext(),
				BuildSys: &typgo.BuildSys{
					Descriptor: &typgo.Descriptor{},
				},
			},
			ExpectedErr: "typrls: missing summary",
		},
		{
			TestName: "missing tag",
			ReleaseProject: &typrls.ReleaseProject{
				Summarizer: typrls.DefaultSummarizer,
			},
			Ctx: &typgo.Context{
				Context: createContext(),
				BuildSys: &typgo.BuildSys{
					Descriptor: &typgo.Descriptor{},
				},
			},
			ExpectedErr: "typrls: missing tag",
		},
		{
			TestName: "bad summary",
			ReleaseProject: &typrls.ReleaseProject{
				Tagger: typrls.DefaultTagger,
				Summarizer: typrls.NewSummarizer(func(*typrls.Context) (string, error) {
					return "", errors.New("bad-summary")
				}),
				Releaser: &typrls.CrossCompiler{},
			},
			Ctx: &typgo.Context{
				Context: createContext(),
				BuildSys: &typgo.BuildSys{
					Descriptor: &typgo.Descriptor{},
				},
			},
			ExpectedErr: "bad-summary",
		},
		{
			TestName: "invalid",
			ReleaseProject: &typrls.ReleaseProject{
				Tagger:     typrls.DefaultTagger,
				Summarizer: typrls.DefaultSummarizer,
				Validator: typrls.NewValidator(func(*typrls.Context) error {
					return errors.New("some-error")
				}),
			},
			Ctx: &typgo.Context{
				Context: createContext(),
				BuildSys: &typgo.BuildSys{
					Descriptor: &typgo.Descriptor{},
				},
			},
			ExpectedErr: "some-error",
		},
		{
			TestName: "with alpha tag",
			ReleaseProject: &typrls.ReleaseProject{
				Tagger: typrls.DefaultTagger,
				Summarizer: typrls.NewSummarizer(func(*typrls.Context) (string, error) {
					return "some-summary", nil
				}),
			},
			Ctx: &typgo.Context{
				Context: createContext("-alpha"),
				BuildSys: &typgo.BuildSys{
					Descriptor: &typgo.Descriptor{},
				},
			},
			RunExpectations: []*execkit.RunExpectation{
				{CommandLine: []string{"git", "fetch"}},
				{CommandLine: []string{"git", "describe", "--tags", "--abbrev=0"}, OutputBytes: []byte("some-tag")},
				{CommandLine: []string{"git", "status", "--porcelain"}, OutputBytes: []byte("some-status")},
				{CommandLine: []string{"git", "--no-pager", "log", "some-tag..HEAD", "--oneline"}, OutputBytes: []byte("5378feb one\n")},
			},
			Expected: &typrls.Context{
				TagName: "v0.0.1_alpha",
				Alpha:   true,
				Summary: "some-summary",
				Git: &typrls.Git{
					Status:     "some-status",
					CurrentTag: "some-tag",
					Logs: []*typrls.Log{
						{ShortCode: "5378feb", Message: "one"},
					},
				},
				ReleaseFolder: "release",
			},
		},
		{
			TestName: "success",
			ReleaseProject: &typrls.ReleaseProject{
				Tagger: typrls.DefaultTagger,
				Summarizer: typrls.NewSummarizer(func(*typrls.Context) (string, error) {
					return "some-summary", nil
				}),
			},
			Ctx: &typgo.Context{
				Context: createContext(),
				BuildSys: &typgo.BuildSys{
					Descriptor: &typgo.Descriptor{
						ProjectVersion: "9.9.9",
					},
				},
			},
			RunExpectations: []*execkit.RunExpectation{
				{CommandLine: []string{"git", "fetch"}},
				{CommandLine: []string{"git", "describe", "--tags", "--abbrev=0"}, OutputBytes: []byte("some-tag-1")},
				{CommandLine: []string{"git", "status", "--porcelain"}, OutputBytes: []byte("some-status-1")},
			},
			Expected: &typrls.Context{
				TagName:       "v9.9.9",
				Alpha:         false,
				Summary:       "some-summary",
				Git:           &typrls.Git{Status: "some-status-1", CurrentTag: "some-tag-1"},
				ReleaseFolder: "release",
			},
		},
		{
			TestName: "with custom release folder",
			ReleaseProject: &typrls.ReleaseProject{
				Tagger: typrls.DefaultTagger,
				Summarizer: typrls.NewSummarizer(func(*typrls.Context) (string, error) {
					return "some-summary", nil
				}),
				ReleaseFolder: "some-release",
			},
			Ctx: &typgo.Context{
				Context: createContext(),
				BuildSys: &typgo.BuildSys{
					Descriptor: &typgo.Descriptor{
						ProjectVersion: "9.9.9",
					},
				},
			},
			RunExpectations: []*execkit.RunExpectation{
				{CommandLine: []string{"git", "fetch"}},
				{CommandLine: []string{"git", "describe", "--tags", "--abbrev=0"}, OutputBytes: []byte("some-tag-1")},
				{CommandLine: []string{"git", "status", "--porcelain"}, OutputBytes: []byte("some-status-1")},
			},
			Expected: &typrls.Context{
				TagName:       "v9.9.9",
				Alpha:         false,
				Summary:       "some-summary",
				Git:           &typrls.Git{Status: "some-status-1", CurrentTag: "some-tag-1"},
				ReleaseFolder: "some-release",
			},
		},
		{
			TestName: "override release-tag",
			ReleaseProject: &typrls.ReleaseProject{
				Tagger: typrls.DefaultTagger,
				Summarizer: typrls.NewSummarizer(func(*typrls.Context) (string, error) {
					return "some-summary", nil
				}),
			},
			Ctx: &typgo.Context{
				Context: createContext("-tag-name=some-tag"),
				BuildSys: &typgo.BuildSys{
					Descriptor: &typgo.Descriptor{
						ProjectVersion: "9.9.9",
					},
				},
			},
			RunExpectations: []*execkit.RunExpectation{
				{CommandLine: []string{"git", "fetch"}},
				{CommandLine: []string{"git", "describe", "--tags", "--abbrev=0"}, OutputBytes: []byte("some-tag-3")},
				{CommandLine: []string{"git", "status", "--porcelain"}, OutputBytes: []byte("some-status-3")},
			},
			Expected: &typrls.Context{
				TagName:       "some-tag",
				Alpha:         false,
				Summary:       "some-summary",
				Git:           &typrls.Git{Status: "some-status-3", CurrentTag: "some-tag-3"},
				ReleaseFolder: "release",
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
				require.Equal(t, tt.Expected.TagName, rlsCtx.TagName)
				require.Equal(t, tt.Expected.Alpha, rlsCtx.Alpha)
				require.Equal(t, tt.Expected.Summary, rlsCtx.Summary)
				require.Equal(t, tt.Expected.Git, rlsCtx.Git)
				require.Equal(t, tt.Expected.ReleaseFolder, rlsCtx.ReleaseFolder)
			}
		})
	}
}

func createContext(args ...string) *cli.Context {
	flagSet := flag.NewFlagSet("test", 0)
	flagSet.Bool(typrls.AlphaFlag, false, "")
	flagSet.Bool(typrls.ForceFlag, false, "")
	flagSet.String(typrls.TagNameFlag, "", "")
	flagSet.Parse(args)
	return cli.NewContext(nil, flagSet, nil)
}

func TestReleaseCmd_Before(t *testing.T) {
	cmd := &typrls.ReleaseCmd{
		Before: typgo.NewAction(func(*typgo.Context) error {
			return errors.New("some-error")
		}),
	}
	command := cmd.Command(&typgo.BuildSys{})
	require.EqualError(t, command.Before(&cli.Context{}), "some-error")
}
