package app

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/typical-go/typical-go/pkg/common"
	"github.com/typical-go/typical-go/pkg/execkit"
	"github.com/urfave/cli/v2"
)

func cmdSetup() *cli.Command {
	return &cli.Command{
		Name:  "setup",
		Usage: "Setup typical-go",
		Flags: []cli.Flag{
			projectPkgFlag,
			typicalBuildFlag,
			typicalTmpFlag,
			&cli.StringFlag{Name: "gomod", Usage: "Iniate go.mod before setup if not empty"},
			&cli.BoolFlag{Name: "new", Usage: "Setup new project with standard layout and typical-build"},
		},
		Action: Setup,
	}
}

// Setup typical-go
func Setup(c *cli.Context) error {
	if gomod := c.String("gomod"); gomod != "" {
		if err := initGoMod(c.Context, gomod); err != nil {
			return err
		}
	}

	p, err := GetParam(c)
	if err != nil {
		return err
	}

	if c.Bool("new") {
		newProject(p)
	}
	return createWrapper(p)
}

func initGoMod(ctx context.Context, pkg string) error {
	var stderr strings.Builder
	fmt.Fprintf(Stdout, "Initiate go.mod\n")
	if err := execkit.Run(ctx, &execkit.Command{
		Name:   "go",
		Args:   []string{"mod", "init", pkg},
		Stderr: &stderr,
	}); err != nil {
		return fmt.Errorf("%s: %s", err.Error(), stderr.String())
	}
	return nil
}

func createWrapper(p *Param) error {
	path := fmt.Sprintf("%s/typicalw", p.SetupTarget)
	fmt.Fprintf(Stdout, "Create '%s'\n", path)
	return common.ExecuteTmplToFile(path, typicalwTmpl, p)
}

func newProject(p *Param) {
	mainPkg := p.SetupTarget + "/cmd/" + p.ProjectName
	main := mainPkg + "/main.go"
	fmt.Fprintf(Stdout, "Create '%s'\n", main)
	os.MkdirAll(mainPkg, 0777)
	common.ExecuteTmplToFile(main, mainTmpl, p)

	appPkg := p.SetupTarget + "/internal/app"
	appStart := appPkg + "/start.go"
	fmt.Fprintf(Stdout, "Create '%s'\n", appStart)
	os.MkdirAll(appPkg, 0777)
	ioutil.WriteFile(appStart, []byte(appStartSrc), 0777)

	generatedPkg := p.SetupTarget + "/internal/generated"
	generatedDoc := generatedPkg + "/doc.go"
	fmt.Fprintf(Stdout, "Create '%s'\n", generatedDoc)
	os.MkdirAll(generatedPkg, 0777)
	ioutil.WriteFile(generatedDoc, []byte(generatedDocSrc), 0777)

	typicalBuildPkg := p.SetupTarget + "/tools/typical-build"
	typicalBuild := typicalBuildPkg + "/typical-build.go"
	fmt.Fprintf(Stdout, "Create '%s'\n", typicalBuild)
	os.MkdirAll(typicalBuildPkg, 0777)
	common.ExecuteTmplToFile(typicalBuild, typicalBuildTmpl, p)

}
