package typgo

import (
	"fmt"
	"os"

	"github.com/typical-go/typical-go/pkg/buildkit"
	"github.com/typical-go/typical-go/pkg/execkit"
	"github.com/typical-go/typical-go/pkg/typtmpl"
	"github.com/typical-go/typical-go/pkg/typvar"
)

var _ Build = (*StdBuild)(nil)

// StdBuild is standard build module for go project
type StdBuild struct{}

// Execute build
func (b *StdBuild) Execute(c *Context, phase Phase) (ok bool, err error) {
	switch phase {
	case RunPhase:
		return true, executeRun(c)
	case TestPhase:
		return true, executeTest(c)
	case CleanPhase:
		return true, executeClean(c)
	}
	return false, nil
}

func executeRun(c *Context) (err error) {
	binary := fmt.Sprintf("%s/%s", typvar.BinFolder, c.Descriptor.Name)
	srcDir := fmt.Sprintf("%s/%s", typvar.CmdFolder, c.Descriptor.Name)
	srcMain := fmt.Sprintf("./%s/main.go", srcDir)

	// NOTE: create main.go if not exist
	if _, err = os.Stat(srcMain); os.IsNotExist(err) {
		os.MkdirAll(srcDir, 0777)
		appMain := &typtmpl.AppMain{
			DescPkg: typvar.ProjectPkg + "/typical",
		}

		if err = typtmpl.WriteFile(srcMain, 0777, appMain); err != nil {
			return fmt.Errorf("%s: %w", srcDir, err)
		}
	}

	gobuild := buildkit.NewGoBuild(binary, "./"+srcDir).Command()
	gobuild.Stderr = os.Stderr
	gobuild.Stdout = os.Stderr

	gobuild.Print(os.Stdout)

	ctx := c.Ctx()
	if err = gobuild.Run(ctx); err != nil {
		return fmt.Errorf("GoBuild: %w", err)
	}

	binExec := &execkit.Command{
		Name:   binary,
		Args:   c.Args().Slice(),
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}

	binExec.Print(os.Stdout)
	fmt.Printf("\n\n")

	if err = binExec.Run(ctx); err != nil {
		return fmt.Errorf("%s: %w", binary, err)
	}

	return
}

func executeTest(c *Context) (err error) {
	var (
		targets []string
	)

	for _, layout := range c.Descriptor.Layouts {
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
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.Print(os.Stdout)
	fmt.Println()

	return cmd.Run(c.Ctx())
}

func executeClean(c *Context) (err error) {
	c.Infof("Remove All in '%s'", typvar.BinFolder)
	if err := os.RemoveAll(typvar.BinFolder); err != nil {
		c.Warn(err.Error())
	}

	c.Infof("Remove All: %s", typvar.TypicalTmp)
	if err := os.RemoveAll(typvar.TypicalTmp); err != nil {
		c.Warn(err.Error())
	}
	return
}
