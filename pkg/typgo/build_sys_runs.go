package typgo

type (
	// BuildSysRuns run command in current BuildSys
	BuildSysRuns []string
)

var _ Action = (BuildSysRuns)(nil)

// Execute BuildSysRuns
func (r BuildSysRuns) Execute(c *Context) error {
	for _, name := range r {
		if err := c.BuildSys.Run(name, c.Context); err != nil {
			return err
		}
	}
	return nil
}
