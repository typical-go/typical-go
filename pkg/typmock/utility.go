package typmock

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/typical-go/typical-go/pkg/execkit"
	"github.com/typical-go/typical-go/pkg/typast"
	"github.com/typical-go/typical-go/pkg/typvar"

	"github.com/iancoleman/strcase"
	"github.com/typical-go/typical-go/pkg/buildkit"
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/urfave/cli/v2"
)

// Utility to generate mock class
func Utility() typgo.Utility {
	return typgo.NewUtility(commands)
}

func commands(c *typgo.BuildCli) []*cli.Command {
	return []*cli.Command{
		{
			Name:        "mock",
			Usage:       "Generate mock class",
			UsageText:   "mock [package_names]",
			Description: "If package_names is missing then check every package",
			Action:      c.ActionFn("mock", mock),
		},
	}
}

func mock(c *typgo.Context) (err error) {

	mockery := createMockery(c)

	mockgen := fmt.Sprintf("%s/bin/mockgen", typvar.TypicalTmp)
	if err = installIfNotExist(c.Ctx(), mockgen); err != nil {
		return
	}

	targetMap := mockery.TargetMap

	if c.Args().Len() > 0 {
		targetMap = mockery.Filter(c.Args().Slice()...)
	}

	for key, targets := range targetMap {
		mockPkg := fmt.Sprintf("%s_mock", key)

		fmt.Printf("\nRemove all: %s\n", mockPkg)
		os.RemoveAll(mockPkg)

		for _, t := range targets {
			srcPkg := fmt.Sprintf("%s/%s", typvar.ProjectPkg, t.Dir)
			dest := fmt.Sprintf("%s%s/%s.go", t.Parent, t.MockPkg, strcase.ToSnake(t.Source))
			name := fmt.Sprintf("%s.%s", srcPkg, t.Source)

			cmd := &execkit.Command{
				Name: mockgen,
				Args: []string{
					"-destination", dest,
					"-package", t.MockPkg,
					srcPkg,
					t.Source,
				},
				Stderr: os.Stderr,
			}

			cmd.Print(os.Stdout)

			if err = cmd.Run(c.Ctx()); err != nil {
				c.Warnf("Fail to mock '%s': %s", name, err.Error())
			}
		}
	}
	return
}

func createMockery(c *typgo.Context) *Mockery {
	mockery := &Mockery{
		TargetMap:  make(TargetMap),
		ProjectPkg: typvar.ProjectPkg,
	}

	for _, annot := range c.ASTStore.Annots {
		if isMock(annot) {
			pkg := annot.Decl.Pkg
			dir := filepath.Dir(annot.Decl.Path)

			parent := ""
			if dir != "." {
				parent = dir[:len(dir)-len(pkg)]
			}

			mockery.Put(&Mock{
				Dir:     dir,
				Pkg:     pkg,
				Source:  annot.Decl.Name,
				Parent:  parent,
				MockPkg: fmt.Sprintf("%s_mock", pkg),
			})

		}
	}

	return mockery
}

func isMock(annot *typast.Annot) bool {
	return strings.EqualFold(annot.TagName, MockTag) &&
		annot.Decl.Type == typast.Interface
}

func installIfNotExist(ctx context.Context, mockgen string) (err error) {
	if _, err = os.Stat(mockgen); os.IsNotExist(err) {
		gobuild := &buildkit.GoBuild{
			Out:    mockgen,
			Source: "github.com/golang/mock/mockgen",
		}
		cmd := gobuild.Command()
		return cmd.Run(ctx)
	}
	return
}
