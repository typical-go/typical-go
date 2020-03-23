package typbuildtool

import "os"

// Clean build result
func (b *Module) Clean(c *BuildContext) (err error) {
	c.Infof("Remove All in '%s'", c.BinFolder())
	if err := os.RemoveAll(c.BinFolder()); err != nil {
		c.Warn(err.Error())
	}
	return
}
