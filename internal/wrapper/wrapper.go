package wrapper

import (
	"bytes"
	"context"
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/typical-go/typical-go/pkg/execkit"
	"github.com/typical-go/typical-go/pkg/typtmpl"
)

const (
	typicalw = "typicalw"

	typicalTmpParam = "typical-tmp"
	projPkgParam    = "proj-pkg"
	srcParam        = "src"
)

type wrapContext struct {
	context.Context
	args         []string
	typicalTmp   string
	projectPkg   string
	src          string
	chksumTarget string
	bin          string
}

func wrap(c *wrapContext) (err error) {

	// if err := w.generateTypicalwIfNotExist(typicalTmp, projectPkg); err != nil {
	// 	return err
	// }

	chksum := generateChecksum(c.src)
	chksum0, _ := ioutil.ReadFile(c.chksumTarget)
	_, err = os.Stat(c.chksumTarget)

	if os.IsNotExist(err) || bytes.Compare(chksum, chksum0) != 0 {
		if err = ioutil.WriteFile(c.chksumTarget, chksum, 0777); err != nil {
			return err
		}

		fmt.Printf("Build %s as %s\n", c.src, c.bin)
		if err := execkit.Run(c.Context, &execkit.GoBuild{
			Output: c.bin,
			Source: "./" + c.src,
			Ldflags: execkit.BuildVars{
				"github.com/typical-go/typical-go/pkg/typgo.ProjectPkg": c.projectPkg,
				"github.com/typical-go/typical-go/pkg/typgo.TypicalTmp": c.typicalTmp,
			},
		}); err != nil {
			return err
		}
	}

	return execkit.Run(c.Context, &execkit.Command{
		Name:   c.bin,
		Args:   c.args,
		Stdout: os.Stdout,
		Stdin:  os.Stdin,
		Stderr: os.Stderr,
	})
}

func generateTypicalw(target, typicalTmp, projectPkg string) error {
	f, err := os.OpenFile(target, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
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
