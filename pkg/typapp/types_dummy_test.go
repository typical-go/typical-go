package typapp_test

import (
	"github.com/typical-go/typical-go/pkg/typapp"
	"github.com/typical-go/typical-go/pkg/typdep"
	"github.com/urfave/cli/v2"
)

type entryPointer struct {
	invocation *typdep.Invocation
}

func dummyEntryPointer(invocation *typdep.Invocation) *entryPointer {
	return &entryPointer{
		invocation: invocation,
	}
}

func (e *entryPointer) EntryPoint() *typdep.Invocation {
	return e.invocation
}

type provider struct {
	constructors []*typdep.Constructor
}

func dummyProvider(constructors ...*typdep.Constructor) *provider {
	return &provider{
		constructors: constructors,
	}
}

func (p *provider) Provide() []*typdep.Constructor {
	return p.constructors
}

type preparer struct {
	invocations []*typdep.Invocation
}

func dummyPreparer(invocations ...*typdep.Invocation) *preparer {
	return &preparer{
		invocations: invocations,
	}
}

func (p *preparer) Prepare() []*typdep.Invocation {
	return p.invocations
}

type destroyer struct {
	invocations []*typdep.Invocation
}

func dummyDestroyers(invocations ...*typdep.Invocation) *destroyer {
	return &destroyer{
		invocations: invocations,
	}
}

func (d *destroyer) Destroy() []*typdep.Invocation {
	return d.invocations
}

type commander struct {
	cmds []*cli.Command
}

func dummyCommander(cmds ...*cli.Command) *commander {
	return &commander{
		cmds: cmds,
	}
}

func (c *commander) Commands(*typapp.Context) []*cli.Command {
	return c.cmds
}
