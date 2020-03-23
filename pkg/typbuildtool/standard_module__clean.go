package typbuildtool

import "os"

// Clean build result
func (b *StandardModule) Clean(c *BuildContext) (err error) {
	c.Info("Standard-Build: Clean the project")
	c.Infof("Remove All in '%s'", c.BinFolder())
	if err := os.RemoveAll(c.BinFolder()); err != nil {
		c.Warn(err.Error())
	}
	return
}
