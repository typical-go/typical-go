package execkit

import (
	"context"
	"fmt"
	"os"
	"strings"
)

type (
	// GoBuild builder
	GoBuild struct {
		Ldflags fmt.Stringer
		Output  string
		Source  string
	}
	// BuildVars to injected variable when build
	BuildVars map[string]string
)

//
// GoBuild
//

var _ fmt.Stringer = (*GoBuild)(nil)
var _ Runner = (*Command)(nil)

// Command of GoBuild
func (g *GoBuild) Command() *Command {
	return &Command{
		Name:   "go",
		Args:   g.Args(),
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
}

// Args is arguments for go build
func (g *GoBuild) Args() []string {
	args := []string{"build"}
	if g.Ldflags != nil {
		args = append(args, "-ldflags", g.Ldflags.String())
	}
	args = append(args, "-o", g.Output, g.Source)
	return args
}

// Run gobuild
func (g *GoBuild) Run(ctx context.Context) error {
	return g.Command().Run(ctx)
}

func (g GoBuild) String() string {
	return g.Command().String()
}

//
// BuildVars
//

var _ fmt.Stringer = (BuildVars)(nil)

func (b BuildVars) String() string {
	var args []string
	for name, value := range b {
		args = append(args, fmt.Sprintf("-X %s=%v", name, value))

	}
	return strings.Join(args, " ")
}

// BuildVar return ldflag argument for set build variable
func BuildVar(name string, value interface{}) string {
	return fmt.Sprintf("-X %s=%v", name, value)
}
