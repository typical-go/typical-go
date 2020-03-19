package typbuildtool

import (
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/urfave/cli/v2"

	"github.com/typical-go/typical-go/pkg/exor"
	"github.com/typical-go/typical-go/pkg/typast"
	"github.com/typical-go/typical-go/pkg/typcore"
)

// Module of build-tool
type Module struct {
	stdout         io.Writer
	stderr         io.Writer
	coverProfile   string
	releaseTargets []ReleaseTarget
	releaseFolder  string
}

// StandardBuild return new instance of Module
func StandardBuild() *Module {
	return &Module{
		stdout:       os.Stdout,
		stderr:       os.Stderr,
		coverProfile: "cover.out",
		releaseTargets: []ReleaseTarget{
			"linux/amd64",
			"darwin/amd64",
		},
		releaseFolder: "release",
	}
}

// WithStdout return StdBuilder with new stdout
func (b *Module) WithStdout(stdout io.Writer) *Module {
	b.stdout = stdout
	return b
}

// WithStderr return StdBuilder with new stderr
func (b *Module) WithStderr(stderr io.Writer) *Module {
	b.stderr = stderr
	return b
}

// WithCoverProfile return StdTester with new cover profile
func (b *Module) WithCoverProfile(coverProfile string) *Module {
	b.coverProfile = coverProfile
	return b
}

// WithReleaseTargets to set target and return its instance
func (b *Module) WithReleaseTargets(releaseTargets ...ReleaseTarget) *Module {
	b.releaseTargets = releaseTargets
	return b
}

// WithReleaseFolder return StdReleaser with new release folder
func (b *Module) WithReleaseFolder(releaseFolder string) *Module {
	b.releaseFolder = releaseFolder
	return b
}

// Validate the releaser
func (b *Module) Validate() (err error) {
	if len(b.releaseTargets) < 1 {
		return errors.New("Missing 'Targets'")
	}
	for _, target := range b.releaseTargets {
		if err = target.Validate(); err != nil {
			return fmt.Errorf("Target: %w", err)
		}
	}
	return
}

// Commands of build-tool
func (b *Module) Commands(c *Context) []*cli.Command {
	return []*cli.Command{
		{
			Name:  "mock",
			Usage: "Generate mock class",
			Action: func(cliCtx *cli.Context) (err error) {
				return b.Mock(&BuildContext{
					Context: c,
					Cli:     cliCtx,
				})
			},
		},
	}
}

// Build the project
func (b *Module) Build(c *BuildContext) (dists []BuildDistribution, err error) {
	binary := fmt.Sprintf("%s/%s", c.BinFolder(), c.Name)
	srcDir := fmt.Sprintf("%s/%s", c.CmdFolder(), c.Name)
	src := fmt.Sprintf("./%s/main.go", srcDir)
	ctx := c.Cli.Context

	// NOTE: create main.go if not exist
	if _, err = os.Stat(src); os.IsNotExist(err) {
		os.MkdirAll(srcDir, 0777)
		if err = typcore.WriteAppMain(ctx, src, c.ProjectPackage+"/typical"); err != nil {
			return
		}
	}

	gobuild := exor.NewGoBuild(binary, src).
		WithStdout(b.stdout).
		WithStderr(b.stderr)

	if err = gobuild.Execute(ctx); err != nil {
		return
	}

	return []BuildDistribution{
		NewBuildDistribution(binary),
	}, nil
}

// Clean build result
func (b *Module) Clean(c *BuildContext) (err error) {
	c.Infof("Remove All in '%s'", c.BinFolder())
	if err := os.RemoveAll(c.BinFolder()); err != nil {
		c.Error(err.Error())
	}
	return
}

// Test the project
func (b *Module) Test(c *BuildContext) (err error) {
	var targets []string
	for _, source := range c.ProjectSources {
		targets = append(targets, fmt.Sprintf("./%s/...", source))
	}

	gotest := exor.NewGoTest(targets...).
		WithCoverProfile(b.coverProfile).
		WithRace(true).
		WithStdout(os.Stdout).
		WithStderr(os.Stderr)

	return gotest.Execute(c.Cli.Context)
}

// Mock the project
func (b *Module) Mock(c *BuildContext) (err error) {
	ctx := c.Cli.Context
	store := NewMockStore()
	if err = c.Ast().EachAnnotation("mock", typast.InterfaceType, func(decl *typast.Declaration, ann *typast.Annotation) (err error) {
		var (
			pkg     = decl.File.Name.Name
			dir     = filepath.Dir(decl.Path)
			dirDest = dir[:len(dir)-len(pkg)]
			srcPkg  = fmt.Sprintf("%s/%s", c.ProjectPackage, dir)
			mockPkg = fmt.Sprintf("mock_%s", pkg)
			mockDir = fmt.Sprintf("%s%s", dirDest, mockPkg)
			dest    = fmt.Sprintf("%s/%s.go", mockDir, strcase.ToSnake(decl.SourceName))
		)

		store.Put(&MockTarget{
			SrcPkg:  srcPkg,
			SrcName: decl.SourceName,
			MockPkg: mockPkg,
			MockDir: mockDir,
			Dest:    dest,
		})
		return
	}); err != nil {
		return
	}

	mockgen := fmt.Sprintf("%s/bin/mockgen", c.TmpFolder())

	if _, err = os.Stat(mockgen); os.IsNotExist(err) {
		c.Info("Build mockgen")
		if err = exor.NewGoBuild(mockgen, "github.com/golang/mock/mockgen").Execute(ctx); err != nil {
			return
		}
	}

	for pkg, targets := range store.Map() {

		c.Infof("Remove package: %s", pkg)
		os.RemoveAll(pkg)

		for _, target := range targets {
			c.Infof("Generate mock: %s", target.Dest)
			cmd := exec.CommandContext(ctx, mockgen,
				"-destination", target.Dest,
				"-package", target.MockPkg,
				target.SrcPkg,
				target.SrcName,
			)
			cmd.Stderr = os.Stderr
			if err := cmd.Run(); err != nil {
				c.Errorf("Mock '%s' failed: %w", target, err)
			}
		}
	}

	return
}

// Release this project
func (b *Module) Release(c *ReleaseContext) (files []string, err error) {

	for _, target := range b.releaseTargets {
		var binary string
		if binary, err = b.build(c.BuildContext, c.Tag, target); err != nil {
			err = fmt.Errorf("Failed build release: %w", err)
			return
		}
		files = append(files, binary)
	}

	return
}

func (b *Module) build(c *BuildContext, tag string, target ReleaseTarget) (binary string, err error) {
	ctx := c.Cli.Context
	goos := target.OS()
	goarch := target.Arch()
	binary = strings.Join([]string{c.Name, tag, goos, goarch}, "_")
	// TODO: Support CGO
	cmd := exec.CommandContext(ctx, "go", "build",
		"-o", fmt.Sprintf("%s/%s", b.releaseFolder, binary),
		"-ldflags", "-w -s",
		fmt.Sprintf("./%s/%s", c.CmdFolder(), c.Name),
	)
	cmd.Env = append(os.Environ(), "GOOS="+goos, "GOARCH="+goarch)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	if err = cmd.Run(); err != nil {
		return
	}
	return
}
