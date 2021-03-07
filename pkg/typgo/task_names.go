package typgo

import (
	"strings"
)

type (
	// TaskNames run command in current BuildSys
	TaskNames []string
)

var _ Action = (TaskNames)(nil)

// Execute BuildCmdRuns
func (r TaskNames) Execute(c *Context) error {
	for _, name := range r {
		args := []string{c.App.Name}
		args = append(args, strings.Fields(name)...)
		if err := c.App.Run(args); err != nil {
			return err
		}
	}
	return nil
}
