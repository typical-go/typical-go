package execkit

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

type (
	// GoBuild builder
	GoBuild struct {
		Ldflags     fmt.Stringer
		Output      string
		MainPackage string
	}
	// BuildVars to injected variable when build
	BuildVars map[string]string
)

//
// GoBuild
//

var _ Commander = (*GoBuild)(nil)

// Command of GoBuild
func (g *GoBuild) Command() *Command {
	return &Command{
		Name:   "go",
		Args:   g.Args(),
		Stdout: os.Stdout,
		Stderr: os.Stderr,
		Stdin:  os.Stdin,
	}
}

// Args is arguments for go build
func (g *GoBuild) Args() []string {
	args := []string{"build"}
	if g.Ldflags != nil {
		args = append(args, "-ldflags", g.Ldflags.String())
	}
	args = append(args, "-o", g.Output, g.MainPackage)
	return args
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
