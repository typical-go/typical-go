package app

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/typical-go/typical-go/pkg/execkit"
	"github.com/typical-go/typical-go/pkg/tmplkit"
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
			&cli.BoolFlag{Name: "go-mod", Usage: "Iniate go.mod before setup"},
			&cli.BoolFlag{Name: "new", Usage: "Setup new project with standard layout and typical-build"},
		},
		Action: Setup,
	}
}

// Setup typical-go
func Setup(c *cli.Context) error {
	if c.Bool("go-mod") {
		if err := initGoMod(c); err != nil {
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

// initGoMod initiate gomodob
func initGoMod(c *cli.Context) error {
	fmt.Fprintf(Stdout, "Initiate go.mod\n")
	pkg := c.String(ProjectPkgParam)
	if pkg == "" {
		return errors.New("project-pkg is empty")
	}
	dir := filepath.Base(pkg)
	os.Mkdir(dir, 0777)
	var stderr strings.Builder
	if err := execkit.Run(c.Context, &execkit.Command{
		Name:   "go",
		Args:   []string{"mod", "init", pkg},
		Stderr: &stderr,
		Dir:    dir,
	}); err != nil {
		return fmt.Errorf("%s: %s", err.Error(), stderr.String())
	}
	return nil
}

func createWrapper(p *Param) error {
	path := fmt.Sprintf("%s/typicalw", p.SetupTarget)
	fmt.Fprintf(Stdout, "Create '%s'\n", path)
	return tmplkit.WriteFile(path, typicalwTmpl, p)
}

func newProject(p *Param) {
	mainPkg := p.SetupTarget + "/cmd/" + p.ProjectName
	main := mainPkg + "/main.go"
	fmt.Fprintf(Stdout, "Create '%s'\n", main)
	os.MkdirAll(mainPkg, 0777)
	tmplkit.WriteFile(main, mainTmpl, p)

	appPkg := p.SetupTarget + "/internal/app"
	appStart := appPkg + "/start.go"
	fmt.Fprintf(Stdout, "Create '%s'\n", appStart)
	os.MkdirAll(appPkg, 0777)
	ioutil.WriteFile(appStart, []byte(appStartSrc), 0777)

	generatedPkg := p.SetupTarget + "/internal/generated/typical"
	generatedDoc := generatedPkg + "/doc.go"
	fmt.Fprintf(Stdout, "Create '%s'\n", generatedDoc)
	os.MkdirAll(generatedPkg, 0777)
	ioutil.WriteFile(generatedDoc, []byte(generatedDocSrc), 0777)

	typicalBuildPkg := p.SetupTarget + "/tools/typical-build"
	typicalBuild := typicalBuildPkg + "/typical-build.go"
	fmt.Fprintf(Stdout, "Create '%s'\n", typicalBuild)
	os.MkdirAll(typicalBuildPkg, 0777)
	tmplkit.WriteFile(typicalBuild, typicalBuildTmpl, p)

}
