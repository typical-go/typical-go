package typbuildtool

import "os"

// Clean build result
func (b *StdBuild) Clean(c *BuildContext) (err error) {
	c.Info("Standard-Build: Clean the project")
	c.Infof("Remove All in '%s'", c.BuildTool.binFolder)
	if err := os.RemoveAll(c.BuildTool.binFolder); err != nil {
		c.Warn(err.Error())
	}
	return
}
