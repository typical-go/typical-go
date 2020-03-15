package typbuildtool

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/typical-go/typical-go/pkg/common"
	"github.com/typical-go/typical-go/pkg/git"

	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/urfave/cli/v2"
)

// TypicalBuildTool is typical Build Tool for golang project
type TypicalBuildTool struct {
	commanders []Commander
	builder    Builder
	tester     Tester
	mocker     Mocker
	releaser   Releaser
	publishers []Publisher

	binFolder string

	includeBranch   bool
	includeCommitID bool
}

// New return new instance of build
func New() *TypicalBuildTool {
	return &TypicalBuildTool{
		builder:   NewBuilder(),
		tester:    NewTester(),
		mocker:    NewMocker(),
		releaser:  NewReleaser(),
		binFolder: "bin",
	}
}

// AppendCommanders to return BuildTool with appended commander
func (b *TypicalBuildTool) AppendCommanders(commanders ...Commander) *TypicalBuildTool {
	b.commanders = append(b.commanders, commanders...)
	return b
}

// WithBuilder return  BuildTool with new builder
func (b *TypicalBuildTool) WithBuilder(builder Builder) *TypicalBuildTool {
	b.builder = builder
	return b
}

// WithReleaser return BuildTool with new releaser
func (b *TypicalBuildTool) WithReleaser(releaser Releaser) *TypicalBuildTool {
	b.releaser = releaser
	return b
}

// WithPublishers return BuildTool with new publishers
func (b *TypicalBuildTool) WithPublishers(publishers ...Publisher) *TypicalBuildTool {
	b.publishers = publishers
	return b
}

// WithMocker return BuildTool with new mocker
func (b *TypicalBuildTool) WithMocker(mocker Mocker) *TypicalBuildTool {
	b.mocker = mocker
	return b
}

// WithTester return BuildTool with new tester
func (b *TypicalBuildTool) WithTester(tester Tester) *TypicalBuildTool {
	b.tester = tester
	return b
}

// WithBinFolder return BuildTool with new binFolder
func (b *TypicalBuildTool) WithBinFolder(binFolder string) *TypicalBuildTool {
	b.binFolder = binFolder
	return b
}

// BinFolder value
func (b *TypicalBuildTool) BinFolder() string {
	return b.binFolder
}

// Validate build
func (b *TypicalBuildTool) Validate() (err error) {

	if err = common.Validate(b.builder); err != nil {
		return fmt.Errorf("BuildTool: Builder: %w", err)
	}

	if err = common.Validate(b.mocker); err != nil {
		return fmt.Errorf("BuildTool: Mocker: %w", err)
	}

	if err = common.Validate(b.tester); err != nil {
		return fmt.Errorf("BuildTool: Tester: %w", err)
	}

	if err = common.Validate(b.releaser); err != nil {
		return fmt.Errorf("BuildTool: Releaser: %w", err)
	}

	return
}

// RunBuildTool to run the build-tool
func (b *TypicalBuildTool) RunBuildTool(c *typcore.Context) (err error) {
	app := cli.NewApp()
	app.Name = c.Name
	app.Usage = "" // NOTE: intentionally blank
	app.Description = c.Description
	app.Version = c.Version
	app.Commands = b.Commands(c)

	return app.Run(os.Args)
}

// Commands to return command
func (b *TypicalBuildTool) Commands(c *typcore.Context) (cmds []*cli.Command) {

	if b.builder != nil {
		cmds = append(cmds,
			b.buildCommand(c),
			b.runCommand(c),
		)
	}

	cmds = append(cmds, b.cleanCommand(c))

	if b.tester != nil {
		cmds = append(cmds, b.testCommand(c))
	}

	if b.mocker != nil {
		cmds = append(cmds, b.mockCommand(c))
	}

	if b.releaser != nil {
		cmds = append(cmds, b.releaseCommand(c))
	}
	for _, commanders := range b.commanders {
		cmds = append(cmds, commanders.Commands(c)...)
	}
	return cmds
}

// Build task
func (b *TypicalBuildTool) Build(c *Context) (dist BuildDistribution, err error) {
	return b.builder.Build(c)
}

// Publish the release
func (b *TypicalBuildTool) Publish(pc *PublishContext) (err error) {
	for _, publisher := range b.publishers {
		if err = publisher.Publish(pc); err != nil {
			return
		}
	}
	return
}

// Test the project
func (b *TypicalBuildTool) Test(c *Context) error {
	if b.tester == nil {
		panic("TypicalBuildTool: missing tester")
	}
	return b.tester.Test(c)
}

func (b *TypicalBuildTool) buildCommand(c *typcore.Context) *cli.Command {
	return &cli.Command{
		Name:    "build",
		Aliases: []string{"b"},
		Usage:   "Build the binary",
		Action: func(cliCtx *cli.Context) (err error) {
			_, err = b.Build(b.createContext(c, cliCtx))
			return
		},
	}
}

func (b *TypicalBuildTool) runCommand(c *typcore.Context) *cli.Command {
	return &cli.Command{
		Name:            "run",
		Aliases:         []string{"r"},
		Usage:           "Run the binary",
		SkipFlagParsing: true,
		Action: func(cliCtx *cli.Context) (err error) {
			var dist BuildDistribution
			bc := b.createContext(c, cliCtx)

			if dist, err = b.Build(bc); err != nil {
				return
			}

			fmt.Println()
			fmt.Println()

			return dist.Run(bc)
		},
	}
}

func (b *TypicalBuildTool) testCommand(c *typcore.Context) *cli.Command {
	return &cli.Command{
		Name:    "test",
		Aliases: []string{"t"},
		Usage:   "Run the testing",
		Action: func(cliCtx *cli.Context) error {
			return b.Test(b.createContext(c, cliCtx))
		},
	}
}

func (b *TypicalBuildTool) releaseCommand(c *typcore.Context) *cli.Command {
	return &cli.Command{
		Name:  "release",
		Usage: "Release the distribution",
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "no-test", Usage: "Release without run unit test"},
			&cli.BoolFlag{Name: "no-publish", Usage: "Release without create github release"},
			&cli.BoolFlag{Name: "force", Usage: "Release by passed all validation"},
			&cli.BoolFlag{Name: "alpha", Usage: "Release for alpha version"},
		},
		Action: func(cliCtx *cli.Context) (err error) {

			bc := b.createContext(c, cliCtx)

			if !cliCtx.Bool("no-test") {
				if err = b.Test(bc); err != nil {
					return
				}
			}

			ctx := cliCtx.Context
			force := cliCtx.Bool("force")
			alpha := cliCtx.Bool("alpha")

			if err = git.Fetch(ctx); err != nil {
				return fmt.Errorf("Failed git fetch: %w", err)
			}
			defer git.Fetch(ctx)

			tag := b.Tag(ctx, c.Version, alpha)

			if status := git.Status(ctx); status != "" && !force {
				return fmt.Errorf("Please commit changes first:\n%s", status)
			}

			var latest string
			if latest = git.LatestTag(ctx); latest == tag && !force {
				return fmt.Errorf("%s already released", latest)
			}

			var gitLogs []*git.Log
			if gitLogs = git.RetrieveLogs(ctx, latest); len(gitLogs) < 1 && !force {
				return errors.New("No change to be released")
			}

			rlsCtx := &ReleaseContext{
				Context: bc,
				Alpha:   alpha,
				Tag:     tag,
				GitLogs: gitLogs,
			}

			var releaseFiles []string
			releaseFiles, err = b.releaser.Release(rlsCtx)

			if !cliCtx.Bool("no-publish") {
				publishCtx := &PublishContext{
					ReleaseContext: rlsCtx,
					ReleaseFiles:   releaseFiles,
				}
				if err = b.Publish(publishCtx); err != nil {
					err = fmt.Errorf("Failed to publish: %w", err)
					return
				}
			}

			return
		},
	}
}

func (b *TypicalBuildTool) mockCommand(c *typcore.Context) *cli.Command {
	return &cli.Command{
		Name:  "mock",
		Usage: "Generate mock class",
		Action: func(cliCtx *cli.Context) (err error) {
			if b.mocker == nil {
				panic("Mocker is nil")
			}
			return b.mocker.Mock(b.createContext(c, cliCtx))
		},
	}
}

func (b *TypicalBuildTool) cleanCommand(c *typcore.Context) *cli.Command {
	return &cli.Command{
		Name:    "clean",
		Aliases: []string{"c"},
		Usage:   "Clean the project from generated file during build time",
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "short", Aliases: []string{"s"}, Usage: "Short version of clean only clean build-tool"},
		},
		Action: func(cliCtx *cli.Context) (err error) {

			return b.Clean(b.createContext(c, cliCtx))
		},
	}
}

// Clean the project
func (b *TypicalBuildTool) Clean(c *Context) (err error) {
	removeAll(b.binFolder)
	if c.Cli.Bool("short") {
		remove(typcore.BuildToolBin(c.TempFolder))
	} else {
		removeAll(c.TempFolder)
	}
	return
}

// Tag return relase tag
func (b *TypicalBuildTool) Tag(ctx context.Context, version string, alpha bool) string {
	var builder strings.Builder
	builder.WriteString("v")
	builder.WriteString(version)
	if b.includeBranch {
		builder.WriteString("_")
		builder.WriteString(git.Branch(ctx))
	}
	if b.includeCommitID {
		builder.WriteString("_")
		builder.WriteString(git.LatestCommit(ctx))
	}
	if alpha {
		builder.WriteString("_alpha")
	}
	return builder.String()
}

// createContext to create new instance of Context
func (b *TypicalBuildTool) createContext(tc *typcore.Context, cc *cli.Context) *Context {
	return &Context{
		Context:          tc,
		TypicalBuildTool: b,
		Cli:              cc,
	}
}

func removeAll(path string) {
	log.Infof("Remove All: %s", path)
	if err := os.RemoveAll(path); err != nil {
		log.Error(err.Error())
	}
}

func remove(path string) {
	log.Infof("Remove: %s", path)
	if err := os.Remove(path); err != nil {
		log.Error(err.Error())
	}
}
