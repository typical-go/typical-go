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

func cmdRun() *cli.Command {
	return &cli.Command{
		Name:   "run",
		Usage:  "Run build-tool for project in current working directory",
		Flags:  []cli.Flag{srcFlag, projPkgFlag, typicalTmpFlag},
		Action: run,
	}
}

func run(c *cli.Context) error {
	p, err := getParam(c)
	if err != nil {
		return err
	}

	chksumTarget := fmt.Sprintf("%s/checksum", p.TypicalTmp)
	bin := fmt.Sprintf("%s/bin/%s", p.TypicalTmp, filepath.Base(p.Src))

	chksum := generateChecksum(p.Src)
	chksum0, _ := ioutil.ReadFile(chksumTarget)
	_, err = os.Stat(chksumTarget)

	if os.IsNotExist(err) || bytes.Compare(chksum, chksum0) != 0 {
		if err = ioutil.WriteFile(chksumTarget, chksum, 0777); err != nil {
			return err
		}

		fmt.Printf("Build %s as %s\n", p.Src, bin)
		if err := execkit.Run(c.Context, &execkit.GoBuild{
			Output:      bin,
			MainPackage: "./" + p.Src,
			Ldflags: execkit.BuildVars{
				"github.com/typical-go/typical-go/pkg/typgo.ProjectPkg": p.ProjectPkg,
				"github.com/typical-go/typical-go/pkg/typgo.TypicalTmp": p.TypicalTmp,
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
