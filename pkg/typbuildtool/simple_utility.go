package typbuildtool

import (
	"github.com/typical-go/typical-go/pkg/typcfg"
	"github.com/urfave/cli/v2"
)

var (
	_ Utility           = (*SimpleUtility)(nil)
	_ typcfg.Configurer = (*SimpleUtility)(nil)
)

// SimpleUtility return command based on command function
type SimpleUtility struct {
	fn         UtilityFn
	configurer typcfg.Configurer
}

// UtilityFn is a function to return command
type UtilityFn func(ctx *Context) []*cli.Command

// NewUtility return new instance SimpleUtility
func NewUtility(fn UtilityFn) *SimpleUtility {
	return &SimpleUtility{
		fn: fn,
	}
}

// Configure the utility
func (s *SimpleUtility) Configure(configurer typcfg.Configurer) *SimpleUtility {
	s.configurer = configurer
	return s
}

// Commands return list of command
func (s *SimpleUtility) Commands(ctx *Context) (cmds []*cli.Command) {
	return s.fn(ctx)
}

// Configurations of utility
func (s *SimpleUtility) Configurations() (cfgs []*typcfg.Configuration) {
	if s.configurer != nil {
		return s.configurer.Configurations()
	}
	return
}
