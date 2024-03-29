package typrls

import (
	"fmt"
	"os"
	"strings"

	"github.com/typical-go/typical-go/pkg/typgo"
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
			c.ReleaseFolder, c.Descriptor.ProjectName, c.TagName, goos, goarch)

		c.Infof("\nGOOS=%s GOARC=%s", goos, goarch)
		os.Setenv("GOOS", goos)
		os.Setenv("GOARC", goarch)

		err := c.ExecuteCommand(&typgo.GoBuild{
			Output:      output,
			MainPackage: o.getMainPackage(c),
			Ldflags: typgo.BuildVars{
				"github.com/typical-go/typical-go/pkg/typgo.ProjectName":    c.Descriptor.ProjectName,
				"github.com/typical-go/typical-go/pkg/typgo.ProjectVersion": c.TagName,
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
		o.MainPackage = fmt.Sprintf("./cmd/%s", c.Descriptor.ProjectName)
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
