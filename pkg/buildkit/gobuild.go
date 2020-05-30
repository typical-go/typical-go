package buildkit

import (
	"fmt"
	"strings"

	"github.com/typical-go/typical-go/pkg/execkit"
)

// GoBuild builder
type GoBuild struct {
	Ldflags []string
	Out     string
	Source  string
}

// BuildVar return ldflag argument for set build variable
func BuildVar(name string, value interface{}) string {
	return fmt.Sprintf("-X %s=%v", name, value)
}

// Command of GoBuild
func (g *GoBuild) Command() *execkit.Command {
	return &execkit.Command{
		Name: "go",
		Args: g.Args(),
	}
}

// Args is arguments for go build
func (g *GoBuild) Args() []string {
	args := []string{"build"}
	if len(g.Ldflags) > 0 {
		args = append(args, "-ldflags", strings.Join(g.Ldflags, " "))
	}
	args = append(args, "-o", g.Out, g.Source)
	return args
}
