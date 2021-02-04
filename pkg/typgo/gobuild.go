package typgo

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/urfave/cli/v2"
)

type (
	// GoBuild command
	GoBuild struct {
		MainPackage string // By default is "cmd/PROJECT_NAME"
		Output      string // By default is "bin/PROJECT_NAME"
		// By default is set variable typgo.ProjectName to PROJECT_NAME
		// and typgo.ProjectVersion to PROJECT-VERSION
		Ldflags fmt.Stringer
	}
	// BuildVars to injected variable when build
	BuildVars map[string]string
)

var _ Tasker = (*GoBuild)(nil)
var _ Action = (*GoBuild)(nil)
var _ Basher = (*GoBuild)(nil)

// Task for gobuild
func (p *GoBuild) Task(b *Descriptor) *cli.Command {
	return &cli.Command{
		Name:    "build",
		Aliases: []string{"b"},
		Usage:   "build the project",
		Action:  b.Action(p),
	}
}

// Execute standard compile
func (p *GoBuild) Execute(c *Context) error {
	if p.MainPackage == "" {
		p.MainPackage = fmt.Sprintf("./cmd/%s", c.Descriptor.ProjectName)
	}
	if p.Output == "" {
		p.Output = fmt.Sprintf("bin/%s", c.Descriptor.ProjectName)
	}
	if p.Ldflags == nil {
		p.Ldflags = BuildVars{
			"github.com/typical-go/typical-go/pkg/typgo.ProjectName":    c.Descriptor.ProjectName,
			"github.com/typical-go/typical-go/pkg/typgo.ProjectVersion": c.Descriptor.ProjectVersion,
		}
	}
	return c.Execute(p.Bash())
}

// Bash for go-build
func (p *GoBuild) Bash() *Bash {
	args := []string{"build"}
	if p.Ldflags != nil {
		args = append(args, "-ldflags", p.Ldflags.String())
	}
	if p.Output != "" {
		args = append(args, "-o", p.Output, p.MainPackage)
	}
	return &Bash{
		Name:   "go",
		Args:   args,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}

}

//
// BuildVars
//

var _ fmt.Stringer = (BuildVars)(nil)

func (b BuildVars) String() string {
	var args []string
	for _, key := range b.Keys() {
		args = append(args, fmt.Sprintf("-X %s=%v", key, b[key]))
	}
	return strings.Join(args, " ")
}

// Keys return sorted key
func (b BuildVars) Keys() []string {
	keys := make([]string, 0, len(b))
	for k := range b {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
