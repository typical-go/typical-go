package typgo

import (
	"os"

	"github.com/urfave/cli/v2"
)

type (
	// Cleaner responsible to clean
	Cleaner interface {
		Clean(*Context) error
	}
	// StdClean standard clean
	StdClean struct{}
	// CleanFn function
	CleanFn     func(*Context) error
	cleanerImpl struct {
		fn CleanFn
	}
	// Cleans for composite clean
	Cleans []Cleaner
)

//
// Cleans
//

var _ Cleaner = (Cleans)(nil)

// Clean the cleans
func (s Cleans) Clean(c *Context) error {
	for _, cleaner := range s {
		if err := cleaner.Clean(c); err != nil {
			return err
		}
	}
	return nil
}

//
// cleanerImpl
//

var _ Cleaner = (*cleanerImpl)(nil)

// NewClean return new instance of Cleaner
func NewClean(fn CleanFn) Cleaner {
	return &cleanerImpl{fn: fn}
}

func (s *cleanerImpl) Clean(c *Context) error {
	return s.fn(c)
}

//
// StdClean
//

var _ Cleaner = (*StdClean)(nil)

// Clean project
func (s *StdClean) Clean(c *Context) error {
	removeAll(c, BinFolder)
	removeAll(c, TypicalTmp)
	return nil
}

func removeAll(c *Context, folder string) {
	if err := os.RemoveAll(folder); err == nil {
		c.Infof("RemoveAll: %s", folder)
	}
}

func remove(c *Context, file string) {
	if err := os.Remove(file); err == nil {
		c.Infof("Remove: %s", file)
	}
}

//
// commands
//

func cmdClean(c *BuildCli) *cli.Command {
	return &cli.Command{
		Name:   "clean",
		Usage:  "Clean the project",
		Action: c.ActionFn("CLEAN", c.Clean.Clean),
	}
}
