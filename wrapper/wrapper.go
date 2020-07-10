package wrapper

import (
	"bytes"
	"crypto/sha256"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/typical-go/typical-go/pkg/execkit"
	"github.com/typical-go/typical-go/pkg/typlog"
	"github.com/typical-go/typical-go/pkg/typtmpl"
	"github.com/urfave/cli/v2"
)

const (
	projectPkgVar = "github.com/typical-go/typical-go/pkg/typgo.ProjectPkg"
	typicalTmpVar = "github.com/typical-go/typical-go/pkg/typgo.TypicalTmp"
	gitignore     = ".gitignore"
	typicalw      = "typicalw"

	typicalTmpParam = "typical-tmp"
	projPkgParam    = "proj-pkg"
	srcParam        = "src"
)

type (
	wrapper struct {
		Name    string
		Version string
		typlog.Logger
	}
)

func (w *wrapper) app() *cli.App {
	app := cli.NewApp()
	app.Name = w.Name
	app.Usage = ""       // NOTE: intentionally blank
	app.Description = "" // NOTE: intentionally blank
	app.Version = w.Version

	app.Commands = []*cli.Command{
		{
			Name:  "wrap",
			Usage: "wrap the project with its build-tool",
			Flags: []cli.Flag{
				&cli.StringFlag{Name: typicalTmpParam, Value: ".typical-tmp"},
				&cli.StringFlag{Name: srcParam, Value: "tools/typical-build"},
				&cli.StringFlag{Name: projPkgParam},
			},
			Action: w.wrap,
		},
	}

	return app
}

func (w *wrapper) wrap(c *cli.Context) (err error) {
	var (
		typicalTmp   = c.String(typicalTmpParam)
		projectPkg   = c.String(projPkgParam)
		src          = c.String(srcParam)
		chksumTarget = fmt.Sprintf("%s/checksum", typicalTmp)
		bin          = fmt.Sprintf("%s/bin/%s", typicalTmp, filepath.Base(src))
	)

	if projectPkg != "" {
		if projectPkg, err = w.retrieveProjPkg(c); err != nil {
			return err
		}
	}

	// if err := w.generateTypicalwIfNotExist(typicalTmp, projectPkg); err != nil {
	// 	return err
	// }

	chksum := generateChecksum(src)
	chksum0, _ := ioutil.ReadFile(chksumTarget)
	_, err = os.Stat(chksumTarget)

	if os.IsNotExist(err) || bytes.Compare(chksum, chksum0) != 0 {
		if err = ioutil.WriteFile(chksumTarget, chksum, 0777); err != nil {
			return err
		}

		w.Info("Build the build-tool")
		gobuild := &execkit.GoBuild{
			Out:    bin,
			Source: "./" + src,
			Ldflags: []string{
				execkit.BuildVar(projectPkgVar, projectPkg),
				execkit.BuildVar(typicalTmpVar, typicalTmp),
			},
		}
		if err := gobuild.Run(c.Context); err != nil {
			return err
		}
	}

	typicalBuild := &execkit.Command{
		Name:   bin,
		Args:   c.Args().Slice(),
		Stdout: os.Stdout,
		Stdin:  os.Stdin,
		Stderr: os.Stderr,
	}
	return typicalBuild.Run(c.Context)
}

func (w *wrapper) retrieveProjPkg(c *cli.Context) (string, error) {
	var stderr strings.Builder
	var stdout strings.Builder
	cmd := execkit.Command{
		Name:   "go",
		Args:   []string{"list", "-m"},
		Stdout: &stdout,
		Stderr: &stderr,
	}
	if err := cmd.Run(c.Context); err != nil {
		return "", errors.New(stderr.String())
	}
	return strings.TrimSpace(stdout.String()), nil
}

func (w *wrapper) generateTypicalwIfNotExist(typicalTmp, projectPkg string) error {
	if _, err := os.Stat(typicalw); !os.IsNotExist(err) {
		return nil
	}

	w.Infof("Generate %s", typicalw)
	f, err := os.OpenFile("typicalw", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		return err
	}
	defer f.Close()

	tmpl := &typtmpl.Typicalw{
		TypicalSource: "github.com/typical-go/typical-go/cmd/typical-go",
		TypicalTmp:    typicalTmp,
		ProjectPkg:    projectPkg,
	}
	return tmpl.Execute(f)
}

func generateChecksum(source string) []byte {
	h := sha256.New()
	filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if b, err := ioutil.ReadFile(path); err == nil {
			h.Write(b)
		}
		return nil
	})
	return h.Sum(nil)
}
