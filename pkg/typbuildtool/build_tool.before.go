package typbuildtool

import (
	"fmt"

	"github.com/typical-go/typical-go/pkg/typcfg"
)

// Precondition for this project
func (b *BuildTool) Precondition(c *CliContext) (err error) {
	if !b.enablePrecondition {
		c.Info("Skip the preconditon")
		return
	}

	app := c.Core.App
	if configurer, ok := app.(typcfg.Configurer); ok {
		if err = typcfg.Write(b.configFile, configurer); err != nil {
			return
		}
	}

	if err = typcfg.Write(b.configFile, b); err != nil {
		return
	}

	if preconditioner, ok := app.(Preconditioner); ok {
		if err = preconditioner.Precondition(c); err != nil {
			return fmt.Errorf("Precondition-App: %w", err)
		}
	}

	typcfg.Load(b.configFile)

	return
}

// Configurations of Build-Tool
func (b *BuildTool) Configurations() (cfgs []*typcfg.Configuration) {
	for _, module := range b.buildSequences {
		if configurer, ok := module.(typcfg.Configurer); ok {
			cfgs = append(cfgs, configurer.Configurations()...)
		}
	}

	for _, utility := range b.utilities {
		if configurer, ok := utility.(typcfg.Configurer); ok {
			cfgs = append(cfgs, configurer.Configurations()...)
		}
	}

	return
}
