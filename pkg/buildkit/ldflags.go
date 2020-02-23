package buildkit

import (
	"fmt"
	"strings"
)

// Ldflags for linker flags
type ldflags struct {
	lines []string
}

// SetVariable to set variable using linker
func (l *ldflags) SetVariable(name string, value interface{}) {
	l.appendf("-X %s=%v", name, value)
}

func (l *ldflags) String() string {
	return strings.Join(l.lines, " ")
}

func (l *ldflags) append(arg string) {
	l.lines = append(l.lines, arg)
}

func (l *ldflags) appendf(format string, a ...interface{}) {
	l.lines = append(l.lines, fmt.Sprintf(format, a...))
}
