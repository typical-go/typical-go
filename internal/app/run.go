package app

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/typical-go/typical-go/pkg/execkit"
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/urfave/cli/v2"
)

func cmdRun() *cli.Command {
	return &cli.Command{
		Name:  "run",
		Usage: "Run build-tool for project in current working directory",
		Flags: []cli.Flag{
			projectPkgFlag,
			typicalBuildFlag,
			typicalTmpFlag,
		},
		Action: Run,
	}
}

// Run typical-go
func Run(c *cli.Context) error {
	p, err := GetParam(c)
	if err != nil {
		return err
	}

	chksumTarget := fmt.Sprintf("%s/checksum", p.TypicalTmp)
	bin := fmt.Sprintf("%s/bin/%s", p.TypicalTmp, filepath.Base(p.TypicalBuild))

	chksum := generateChecksum(p.TypicalBuild)
	chksum0, _ := ioutil.ReadFile(chksumTarget)
	_, err = os.Stat(chksumTarget)

	if os.IsNotExist(err) || bytes.Compare(chksum, chksum0) != 0 {
		if err = ioutil.WriteFile(chksumTarget, chksum, 0777); err != nil {
			return err
		}

		fmt.Fprintf(Stdout, "Build %s to %s\n", p.TypicalBuild, bin)

		buildVars := typgo.BuildVars{
			"github.com/typical-go/typical-go/pkg/typgo.ProjectPkg": p.ProjectPkg,
			"github.com/typical-go/typical-go/pkg/typgo.TypicalTmp": p.TypicalTmp,
		}

		args := []string{"build"}
		args = append(args, "-ldflags", buildVars.String())
		args = append(args, "-o", bin, "./"+p.TypicalBuild)

		if err := execkit.Run(c.Context, &execkit.Command{
			Name:   "go",
			Args:   args,
			Stdout: os.Stdout,
			Stderr: os.Stderr,
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
