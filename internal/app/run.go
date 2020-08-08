package app

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/typical-go/typical-go/pkg/execkit"
	"github.com/urfave/cli/v2"
)

func runCommand() *cli.Command {
	return &cli.Command{
		Name:  "run",
		Usage: "Run build-tool for project in current working directory",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: srcParam, Usage: "Build-tool source", Value: "tools/typical-build"},
			&cli.StringFlag{Name: projPkgParam, Usage: "Project package name. Same with module package in go.mod by default"},
			&cli.StringFlag{Name: typicalTmpParam, Usage: "Temporary directory location to save builds-related files", Value: ".typical-tmp"},
			&cli.BoolFlag{Name: createWrapperParam, Usage: "Create wrapper script"},
		},
		Action: func(c *cli.Context) error {
			if c.Bool(createWrapperParam) {
				return wrapper(c)
			}
			return run(c)
		},
	}
}

func run(c *cli.Context) error {
	typicalTmp := getTypicalTmp(c)
	src := getSrc(c)

	projectPkg, err := getProjectPkg(c)
	if err != nil {
		return err
	}

	chksumTarget := fmt.Sprintf("%s/checksum", typicalTmp)
	bin := fmt.Sprintf("%s/bin/%s", typicalTmp, filepath.Base(src))

	chksum := generateChecksum(src)
	chksum0, _ := ioutil.ReadFile(chksumTarget)
	_, err = os.Stat(chksumTarget)

	if os.IsNotExist(err) || bytes.Compare(chksum, chksum0) != 0 {
		if err = ioutil.WriteFile(chksumTarget, chksum, 0777); err != nil {
			return err
		}

		fmt.Printf("Build %s as %s\n", src, bin)
		if err := execkit.Run(c.Context, &execkit.GoBuild{
			Output:      bin,
			MainPackage: "./" + src,
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
