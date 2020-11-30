package typgo

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/typical-go/typical-go/pkg/execkit"
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

var _ CliCommander = (*GoBuild)(nil)
var _ Action = (*GoBuild)(nil)
var _ execkit.Commander = (*GoBuild)(nil)

// Cli compile
func (p *GoBuild) Cli(b *BuildSys) *cli.Command {
	return &cli.Command{
		Name:    "compile",
		Aliases: []string{"c"},
		Usage:   "Compile the project",
		Action:  b.Action(p),
	}
}

// Execute standard compile
func (p *GoBuild) Execute(c *Context) error {
	if p.MainPackage == "" {
		p.MainPackage = fmt.Sprintf("./cmd/%s", c.BuildSys.ProjectName)
	}
	if p.Output == "" {
		p.Output = fmt.Sprintf("bin/%s", c.BuildSys.ProjectName)
	}
	if p.Ldflags == nil {
		p.Ldflags = BuildVars{
			"github.com/typical-go/typical-go/pkg/typgo.ProjectName":    c.BuildSys.ProjectName,
			"github.com/typical-go/typical-go/pkg/typgo.ProjectVersion": c.BuildSys.ProjectVersion,
		}
	}
	return c.Execute(p.Command())
}

// Command bash command
func (p *GoBuild) Command() *execkit.Command {
	args := []string{"build"}
	if p.Ldflags != nil {
		args = append(args, "-ldflags", p.Ldflags.String())
	}
	if p.Output != "" {
		args = append(args, "-o", p.Output, p.MainPackage)
	}
	return &execkit.Command{
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
