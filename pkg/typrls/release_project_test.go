package typrls_test

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/execkit"
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/typical-go/typical-go/pkg/typrls"
	"github.com/urfave/cli/v2"
)

func TestReleaseProject(t *testing.T) {
	var debug strings.Builder
	testcases := []struct {
		TestName string
		typrls.ReleaseProject
		Context     *typgo.Context
		Debug       string
		ExpectedErr string
	}{
		{
			ReleaseProject: typrls.ReleaseProject{
				Tagger:     typrls.DefaultTagger,
				Summarizer: typrls.DefaultSummarizer,
				Releaser: typrls.NewReleaser(func(c *typrls.Context) error {
					fmt.Fprintln(&debug, "1")
					return nil
				}),
				Publisher: typrls.NewPublisher(func(c *typrls.Context) error {
					fmt.Fprintln(&debug, "2")
					return nil
				}),
			},
			Context: &typgo.Context{
				Context:  createContext(),
				BuildSys: &typgo.BuildSys{Descriptor: &typgo.Descriptor{}},
			},
			Debug: "1\n2\n",
		},
		{
			TestName: "release error",
			ReleaseProject: typrls.ReleaseProject{
				Tagger:     typrls.DefaultTagger,
				Summarizer: typrls.DefaultSummarizer,
				Releaser: typrls.NewReleaser(func(c *typrls.Context) error {
					return errors.New("release-error")
				}),
				Publisher: typrls.NewPublisher(func(c *typrls.Context) error {
					return errors.New("publish-error")
				}),
			},
			Context: &typgo.Context{
				Context:  createContext(),
				BuildSys: &typgo.BuildSys{Descriptor: &typgo.Descriptor{}},
			},
			ExpectedErr: "release-error",
		},
		{
			TestName: "publish error",
			ReleaseProject: typrls.ReleaseProject{
				Tagger:     typrls.DefaultTagger,
				Summarizer: typrls.DefaultSummarizer,
				Releaser: typrls.NewReleaser(func(c *typrls.Context) error {
					return nil
				}),
				Publisher: typrls.NewPublisher(func(c *typrls.Context) error {
					return errors.New("publish-error")
				}),
			},
			Context: &typgo.Context{
				Context:  createContext(),
				BuildSys: &typgo.BuildSys{Descriptor: &typgo.Descriptor{}},
			},
			ExpectedErr: "publish-error",
		},
		{
			TestName: "empty publisher",
			ReleaseProject: typrls.ReleaseProject{
				Tagger:     typrls.DefaultTagger,
				Summarizer: typrls.DefaultSummarizer,
				Releaser: typrls.NewReleaser(func(c *typrls.Context) error {
					return nil
				}),
			},
			Context: &typgo.Context{
				Context:  createContext(),
				BuildSys: &typgo.BuildSys{Descriptor: &typgo.Descriptor{}},
			},
		},
		{
			TestName: "skip publish",
			ReleaseProject: typrls.ReleaseProject{
				Tagger:     typrls.DefaultTagger,
				Summarizer: typrls.DefaultSummarizer,
				Releaser: typrls.NewReleaser(func(c *typrls.Context) error {
					return nil
				}),
				Publisher: typrls.NewPublisher(func(c *typrls.Context) error {
					return errors.New("publish-error")
				}),
			},
			Context: &typgo.Context{
				Context:  createContext("-skip-publish"),
				BuildSys: &typgo.BuildSys{Descriptor: &typgo.Descriptor{}},
			},
		},
		{
			TestName: "empty releaser",
			ReleaseProject: typrls.ReleaseProject{
				Tagger:     typrls.DefaultTagger,
				Summarizer: typrls.DefaultSummarizer,
				Publisher: typrls.NewPublisher(func(c *typrls.Context) error {
					return errors.New("publish-error")
				}),
			},
			Context: &typgo.Context{
				Context:  createContext(),
				BuildSys: &typgo.BuildSys{Descriptor: &typgo.Descriptor{}},
			},
			ExpectedErr: "publish-error",
		},
	}
	for _, tt := range testcases {
		t.Run(tt.TestName, func(t *testing.T) {
			defer debug.Reset()
			unpatch := execkit.Patch([]*execkit.RunExpectation{})
			defer unpatch(t)

			err := tt.Execute(tt.Context)
			if tt.ExpectedErr != "" {
				require.EqualError(t, err, tt.ExpectedErr)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.Debug, debug.String())
			}
		})
	}

}

func TestReleaseProject_CustomReleaseFolder(t *testing.T) {
	unpatch := execkit.Patch(nil)
	defer unpatch(t)
	var rlsCtx *typrls.Context

	rel := &typrls.ReleaseProject{
		Tagger: typrls.DefaultTagger,
		Summarizer: typrls.NewSummarizer(func(*typrls.Context) string {
			return "some-summary"
		}),
		Releaser: typrls.NewReleaser(func(r *typrls.Context) error {
			rlsCtx = r
			return nil
		}),
	}

	rel.Execute(&typgo.Context{
		Context: createContext("-release-folder=some-release"),
		BuildSys: &typgo.BuildSys{
			Descriptor: &typgo.Descriptor{ProjectVersion: "9.9.9"},
		},
	})
	defer os.RemoveAll("some-release")
	require.Equal(t, "some-release", rlsCtx.ReleaseFolder)
}

func TestReleaseProject_Execute_Context(t *testing.T) {
	testcases := []struct {
		TestName string
		*typrls.ReleaseProject
		Ctx             *typgo.Context
		RunExpectations []*execkit.RunExpectation
		Expected        *typrls.Context
		ExpectedErr     string
	}{
		{
			ReleaseProject: &typrls.ReleaseProject{},
			Ctx: &typgo.Context{
				Context: createContext(),
				BuildSys: &typgo.BuildSys{
					Descriptor: &typgo.Descriptor{},
				},
			},
			Expected: &typrls.Context{
				TagName:       "v0.0.1",
				Alpha:         false,
				Summary:       "",
				Git:           &typrls.Git{},
				ReleaseFolder: "release",
			},
		},
		{
			TestName: "with alpha tag",
			ReleaseProject: &typrls.ReleaseProject{
				Tagger: typrls.DefaultTagger,
				Summarizer: typrls.NewSummarizer(func(*typrls.Context) string {
					return "some-summary"
				}),
			},
			Ctx: &typgo.Context{
				Context: createContext("-alpha"),
				BuildSys: &typgo.BuildSys{
					Descriptor: &typgo.Descriptor{},
				},
			},
			RunExpectations: []*execkit.RunExpectation{
				{CommandLine: "git fetch"},
				{CommandLine: "git describe --tags --abbrev=0", OutputBytes: []byte("some-tag")},
				{CommandLine: "git status --porcelain", OutputBytes: []byte("some-status")},
				{CommandLine: "git --no-pager log some-tag..HEAD --oneline", OutputBytes: []byte("5378feb one\n")},
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
				Summarizer: typrls.NewSummarizer(func(*typrls.Context) string {
					return "some-summary"
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
				{CommandLine: "git fetch"},
				{CommandLine: "git describe --tags --abbrev=0", OutputBytes: []byte("some-tag-1")},
				{CommandLine: "git status --porcelain", OutputBytes: []byte("some-status-1")},
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
			TestName: "override release-tag",
			ReleaseProject: &typrls.ReleaseProject{
				Tagger: typrls.DefaultTagger,
				Summarizer: typrls.NewSummarizer(func(*typrls.Context) string {
					return "some-summary"
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
				{CommandLine: "git fetch"},
				{CommandLine: "git describe --tags --abbrev=0", OutputBytes: []byte("some-tag-3")},
				{CommandLine: "git status --porcelain", OutputBytes: []byte("some-status-3")},
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
			}
		})
	}
}

func createContext(args ...string) *cli.Context {
	flagSet := flag.NewFlagSet("test", 0)
	flagSet.Bool(typrls.AlphaFlag, false, "")
	flagSet.Bool(typrls.SkipPublishFlag, false, "")
	flagSet.String(typrls.TagNameFlag, "", "")
	flagSet.String(typrls.ReleaseFolderFlag, "release", "")
	flagSet.Parse(args)
	return cli.NewContext(nil, flagSet, nil)
}

func TestReleaseCmd_Before(t *testing.T) {
	cmd := &typrls.ReleaseProject{
		Before: typgo.NewAction(func(*typgo.Context) error {
			return errors.New("some-error")
		}),
	}
	command := cmd.Command(&typgo.BuildSys{})
	require.EqualError(t, command.Before(&cli.Context{}), "some-error")
}
