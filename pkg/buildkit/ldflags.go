package buildkit

import (
	"fmt"
	"strings"
)

// Ldflags for linker flags
type Ldflags struct {
	args []string
}

// NewLdflags return new instance of Ldflags
func NewLdflags() *Ldflags {
	return &Ldflags{}
}

// SetVariable to set variable using linker
func (l *Ldflags) SetVariable(name string, value interface{}) {
	l.appendf("-X %s=%v", name, value)
}

func (l *Ldflags) String() string {
	return strings.Join(l.args, " ")
}

// NotEmpty return true if ldflags is not empty
func (l *Ldflags) NotEmpty() bool {
	return len(l.args) > 0
}

func (l *Ldflags) append(arg string) {
	l.args = append(l.args, arg)
}

func (l *Ldflags) appendf(format string, a ...interface{}) {
	l.args = append(l.args, fmt.Sprintf(format, a...))
}
