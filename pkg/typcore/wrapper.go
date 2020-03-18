package typcore

import (
	"os"

	"github.com/typical-go/typical-go/pkg/exor"
)

// TypicalWrapper responsible to wrap the typical project
type TypicalWrapper struct{}

// NewWrapper return new instance of TypicalWrapper
func NewWrapper() *TypicalWrapper {
	return &TypicalWrapper{}
}

// Wrap the project
func (*TypicalWrapper) Wrap(c *WrapContext) (err error) {

	if c.ProjectPackage == "" {
		c.ProjectPackage = RetrieveProjectPackage()
	}

	// NOTE: create tmp folder if not exist
	os.MkdirAll(c.TmpFolder+"/build-tool", os.ModePerm)
	os.MkdirAll(c.TmpFolder+"/bin", os.ModePerm)

	cksmFile := c.TmpFolder + "/checksum"
	out := c.TmpFolder + "/bin/build-tool"
	srcPath := c.TmpFolder + "/build-tool/main.go"
	descriptorPkg := c.ProjectPackage + "/typical"

	var cksm *Checksum
	if cksm, err = CreateChecksum("typical"); err != nil {
		return
	}

	if _, err = os.Stat(out); os.IsNotExist(err) || !cksm.IsSame(cksmFile) {
		c.Info("Update checksum")
		if err = cksm.Save(cksmFile); err != nil {
			return
		}

		if _, err = os.Stat(srcPath); os.IsNotExist(err) {
			c.Infof("Generate build-tool main source: %s", srcPath)
			if err = WriteBuildToolMain(c.Ctx, srcPath, descriptorPkg); err != nil {
				return
			}
		}

		c.Info("Build the build-tool")
		return exor.NewGoBuild(out, srcPath).
			SetVariable("github.com/typical-go/typical-go/pkg/typcore.DefaultProjectPackage", c.ProjectPackage).
			SetVariable("github.com/typical-go/typical-go/pkg/typbuildtool.DefaultTmpFolder", c.TmpFolder).
			WithStdout(os.Stdout).
			WithStderr(os.Stderr).
			WithStdin(os.Stdin).
			Execute(c.Ctx)
	}

	return
}
