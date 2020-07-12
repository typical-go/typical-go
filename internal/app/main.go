package app

import (
	"bytes"
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/typical-go/typical-go/pkg/execkit"
	"github.com/typical-go/typical-go/pkg/typapp"
	"github.com/typical-go/typical-go/pkg/typtmpl"
	"github.com/urfave/cli/v2"
)

const (
	typicalw = "typicalw"

	typicalTmpParam    = "typical-tmp"
	projPkgParam       = "project-pkg"
	srcParam           = "src"
	createWrapperParam = "create:wrapper"
)

// Main function to run the typical-go
func Main() (err error) {

	app := cli.NewApp()
	app.Name = typapp.Name
	app.Usage = ""       // NOTE: intentionally blank
	app.Description = "" // NOTE: intentionally blank
	app.Version = typapp.Version
	app.Flags = []cli.Flag{
		&cli.StringFlag{Name: typicalTmpParam, Value: ".typical-tmp"},
		&cli.StringFlag{Name: srcParam, Value: "tools/typical-build"},
		&cli.StringFlag{Name: projPkgParam, Usage: "same with module package in go.mod if empty"},
		&cli.BoolFlag{Name: createWrapperParam},
	}
	app.Action = execute

	return app.Run(os.Args)
}

func execute(c *cli.Context) (err error) {
	var (
		typicalTmp   = c.String(typicalTmpParam)
		projectPkg   = c.String(projPkgParam)
		src          = c.String(srcParam)
		chksumTarget = fmt.Sprintf("%s/checksum", typicalTmp)
		bin          = fmt.Sprintf("%s/bin/%s", typicalTmp, filepath.Base(src))
	)

	if projectPkg == "" {
		if projectPkg, err = retrieveProjPkg(c.Context); err != nil {
			return err
		}
	}

	if c.Bool(createWrapperParam) {
		return typtmpl.ExecuteToFile(typicalw, &typtmpl.Typicalw{
			Src:        src,
			TypicalTmp: typicalTmp,
			ProjectPkg: projectPkg,
		})
	}

	chksum := generateChecksum(src)
	chksum0, _ := ioutil.ReadFile(chksumTarget)
	_, err = os.Stat(chksumTarget)

	if os.IsNotExist(err) || bytes.Compare(chksum, chksum0) != 0 {
		if err = ioutil.WriteFile(chksumTarget, chksum, 0777); err != nil {
			return err
		}

		fmt.Printf("Build %s as %s\n", src, bin)
		if err := execkit.Run(c.Context, &execkit.GoBuild{
			Output: bin,
			Source: "./" + src,
			Ldflags: execkit.BuildVars{
				"github.com/typical-go/typical-go/pkg/typgo.ProjectPkg": projectPkg,
				"github.com/typical-go/typical-go/pkg/typgo.TypicalTmp": typicalTmp,
			},
		}); err != nil {
			return err
		}
	}

	return execkit.Run(c.Context, &execkit.Command{
		Name:   bin,
		Args:   c.Args().Slice(),
		Stdout: os.Stdout,
		Stdin:  os.Stdin,
		Stderr: os.Stderr,
	})
}

func retrieveProjPkg(ctx context.Context) (string, error) {
	var stderr strings.Builder
	var stdout strings.Builder
	cmd := execkit.Command{
		Name:   "go",
		Args:   []string{"list", "-m"},
		Stdout: &stdout,
		Stderr: &stderr,
	}
	if err := cmd.Run(ctx); err != nil {
		return "", errors.New(stderr.String())
	}
	return strings.TrimSpace(stdout.String()), nil
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
