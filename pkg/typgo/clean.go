package typgo

import (
	"os"
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
	removeAll(c, TypicalTmp)
	return nil
}

func removeAll(c *Context, folder string) {
	if err := os.RemoveAll(folder); err == nil {
		c.Infof("RemoveAll: %s", folder)
	}
}
