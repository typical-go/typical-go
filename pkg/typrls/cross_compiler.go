package typrls

import (
	"fmt"
	"os"
	"strings"

	"github.com/typical-go/typical-go/pkg/execkit"
)

type (
	// CrossCompiler compile project to various platform
	CrossCompiler struct {
		Targets     []Target
		MainPackage string
	}
	// Target of release contain "$GOOS/$GOARC"
	Target string
)

//
// Compile
//

var _ Releaser = (*CrossCompiler)(nil)

// Release for compile
func (o *CrossCompiler) Release(c *Context) error {
	defer os.Unsetenv("GOOS")
	defer os.Unsetenv("GOARC")

	for _, target := range o.Targets {
		goos := target.OS()
		goarch := target.Arch()
		output := fmt.Sprintf("%s/%s_%s_%s_%s",
			c.ReleaseFolder, c.BuildSys.ProjectName, c.TagName, goos, goarch)

		os.Setenv("GOOS", goos)
		os.Setenv("GOARC", goarch)

		err := c.Execute(&execkit.GoBuild{
			Output:      output,
			MainPackage: o.getMainPackage(c),
			Ldflags: execkit.BuildVars{
				"github.com/typical-go/typical-go/pkg/typgo.AppName":    c.BuildSys.ProjectName,
				"github.com/typical-go/typical-go/pkg/typgo.AppVersion": c.TagName,
			},
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (o *CrossCompiler) getMainPackage(c *Context) string {
	if o.MainPackage == "" {
		o.MainPackage = fmt.Sprintf("./cmd/%s", c.BuildSys.ProjectName)
	}
	return o.MainPackage
}

//
// OSTarget
//

// OS operating system
func (t Target) OS() string {
	i := strings.Index(string(t), "/")
	if i < 0 {
		return ""
	}
	return string(t)[:i]
}

// Arch architecture
func (t Target) Arch() string {
	i := strings.Index(string(t), "/")
	if i < 0 {
		return ""
	}
	return string(t)[i+1:]
}
