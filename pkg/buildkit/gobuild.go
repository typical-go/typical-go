package buildkit

import (
	"fmt"
	"strings"

	"github.com/typical-go/typical-go/pkg/execkit"
)

// GoBuild builder
type GoBuild struct {
	ldflags []string
	out     string
	src     string
}

// NewGoBuild return new instance of gobuild
func NewGoBuild(out, src string) *GoBuild {
	return &GoBuild{
		out: out,
		src: src,
	}
}

// SetVariable to set variable using linker
func (g *GoBuild) SetVariable(name string, value interface{}) *GoBuild {
	g.ldflags = append(g.ldflags, fmt.Sprintf("-X %s=%v", name, value))
	return g
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
	if len(g.ldflags) > 0 {
		args = append(args, "-ldflags", strings.Join(g.ldflags, " "))
	}
	args = append(args, "-o", g.out, g.src)
	return args
}
