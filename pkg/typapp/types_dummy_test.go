package typapp_test

import (
	"github.com/typical-go/typical-go/pkg/typapp"
	"github.com/typical-go/typical-go/pkg/typdep"
	"github.com/urfave/cli/v2"
)

type entryPointer struct {
	invocation *typapp.MainInvocation
}

func dummyEntryPointer(invocation *typapp.MainInvocation) *entryPointer {
	return &entryPointer{
		invocation: invocation,
	}
}

func (e *entryPointer) EntryPoint() *typapp.MainInvocation {
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
	invocations []*typapp.Preparation
}

func dummyPreparer(invocations ...*typapp.Preparation) *preparer {
	return &preparer{
		invocations: invocations,
	}
}

func (p *preparer) Prepare() []*typapp.Preparation {
	return p.invocations
}

type destroyer struct {
	invocations []*typapp.Destruction
}

func dummyDestroyers(invocations ...*typapp.Destruction) *destroyer {
	return &destroyer{
		invocations: invocations,
	}
}

func (d *destroyer) Destroy() []*typapp.Destruction {
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
