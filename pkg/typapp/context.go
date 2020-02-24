package typapp

import (
	log "github.com/sirupsen/logrus"

	"github.com/typical-go/typical-go/pkg/common"
	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/typical-go/typical-go/pkg/typdep"
	"github.com/urfave/cli/v2"
)

// Context of App
type Context struct {
	*typcore.Descriptor
	*App
}

// ActionFunc to return ActionFunc to invoke function fn
func (c *Context) ActionFunc(invocation *typdep.Invocation) func(*cli.Context) error {
	return func(cliCtx *cli.Context) (err error) {
		return c.Invoke(cliCtx, invocation)
	}
}

// Invoke function with Dependency Injection
func (c *Context) Invoke(cliCtx *cli.Context, invocation *typdep.Invocation) (err error) {
	di := typdep.New()

	if cliCtx != nil {
		if err = typdep.NewConstructor(func() *cli.Context {
			return cliCtx
		}).Provide(di); err != nil {
			return
		}
	}

	if c.Configuration != nil {
		// provide configuration to dependency-injection container
		if err = typdep.ProvideAll(di, c.Configuration.Store().Provide()...); err != nil {
			return
		}
	}

	// provide registered function in descriptor to dependency-injection container
	if err = typdep.ProvideAll(di, c.Provide()...); err != nil {
		return
	}

	// invoke preparation as register in descriptor
	if err = typdep.InvokeAll(di, c.Prepare()...); err != nil {
		return
	}

	startFn := func() error { return invocation.Invoke(di) }
	stopFn := func() error { return typdep.InvokeAll(di, c.Destroy()...) }
	for _, err := range common.StartGracefully(startFn, stopFn) {
		log.Error(err.Error())
	}
	return
}
