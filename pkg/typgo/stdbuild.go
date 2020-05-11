package typgo

import (
	"fmt"
	"io"
	"os"

	"github.com/typical-go/typical-go/pkg/buildkit"
	"github.com/typical-go/typical-go/pkg/execkit"
	"github.com/typical-go/typical-go/pkg/typtmpl"
	"github.com/typical-go/typical-go/pkg/typvar"
)

var (
	_ Cleaner = (*StdBuild)(nil)
	_ Tester  = (*StdBuild)(nil)
	_ Runner  = (*StdBuild)(nil)
)

// StdBuild is standard build module for go project
type StdBuild struct {
	stdout io.Writer
	stderr io.Writer
	stdin  io.Reader
}

// StandardBuild return new instance of Module
func StandardBuild() *StdBuild {
	return &StdBuild{
		stdout: os.Stdout,
		stderr: os.Stderr,
		stdin:  os.Stdin,
	}
}

// Run the project locally
func (b *StdBuild) Run(c *Context) (err error) {
	c.Info("Standard-Build: Build the project")
	binary := fmt.Sprintf("%s/%s", typvar.BinFolder, c.BuildTool.Name)
	srcDir := fmt.Sprintf("%s/%s", typvar.CmdFolder, c.BuildTool.Name)
	src := fmt.Sprintf("./%s/main.go", srcDir)

	// NOTE: create main.go if not exist
	if _, err = os.Stat(src); os.IsNotExist(err) {
		os.MkdirAll(srcDir, 0777)
		appMain := &typtmpl.AppMain{
			DescPkg: typvar.ProjectPkg + "/typical",
		}

		if err = typtmpl.WriteFile(src, 0777, appMain); err != nil {
			return fmt.Errorf("%s: %w", srcDir, err)
		}
	}

	gobuild := buildkit.NewGoBuild(binary, src).Command()
	gobuild.Stderr = b.stderr
	gobuild.Stdout = b.stdout

	gobuild.Print(os.Stdout)

	ctx := c.Cli.Context
	if err = gobuild.Run(ctx); err != nil {
		return fmt.Errorf("GoBuild: %w", err)
	}

	binExec := &execkit.Command{
		Name:   binary,
		Args:   c.Cli.Args().Slice(),
		Stdout: b.stdout,
		Stderr: b.stderr,
	}

	binExec.Print(os.Stdout)
	fmt.Printf("\n\n")

	if err = binExec.Run(ctx); err != nil {
		return fmt.Errorf("%s: %w", binary, err)
	}

	return
}

// Test the project
func (b *StdBuild) Test(c *Context) (err error) {
	var (
		targets []string
	)

	for _, layout := range c.BuildTool.Layouts {
		targets = append(targets, fmt.Sprintf("./%s/...", layout))
	}

	if len(targets) < 1 {
		c.Info("Nothing to test")
		return
	}

	gotest := buildkit.GoTest{
		Targets:      targets,
		Timeout:      typvar.TestTimeout,
		CoverProfile: typvar.TestCoverProfile,
		Race:         true,
	}

	cmd := gotest.Command()
	cmd.Stdout = b.stdout
	cmd.Stderr = b.stderr

	cmd.Print(os.Stdout)
	fmt.Println()

	ctx := c.Cli.Context

	return cmd.Run(ctx)
}

// Clean build result
func (b *StdBuild) Clean(c *Context) (err error) {
	c.Infof("Remove All in '%s'", typvar.BinFolder)
	if err := os.RemoveAll(typvar.BinFolder); err != nil {
		c.Warn(err.Error())
	}
	return
}
