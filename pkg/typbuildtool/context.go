package typbuildtool

import (
	log "github.com/sirupsen/logrus"
	"github.com/typical-go/typical-go/pkg/common"
	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/typical-go/typical-go/pkg/typdep"
	"github.com/urfave/cli/v2"
)

// Context of buildtool
type Context struct {
	*typcore.Context
}

// ActionFunc to return ActionFunc to invoke function fn
func (c *Context) ActionFunc(v interface{}) func(*cli.Context) error {
	return func(cliCtx *cli.Context) (err error) {
		if invocation, ok := v.(*typdep.Invocation); ok {
			return c.Invoke(cliCtx, invocation)
		}
		return c.Invoke(cliCtx, typdep.NewInvocation(v))
	}
}

// Invoke function
func (c *Context) Invoke(cliCtx *cli.Context, invocation *typdep.Invocation) (err error) {
	di := typdep.New()

	if err = typdep.NewConstructor(func() *cli.Context {
		return cliCtx
	}).Provide(di); err != nil {
		return
	}

	// provide functions
	if c.Configuration != nil {
		if err = typdep.Provide(di, c.Configuration.Store().Provide()...); err != nil {
			return
		}
	}

	startFn := func() error { return invocation.Invoke(di) }
	for _, err := range common.StartGracefully(startFn, nil) {
		log.Error(err.Error())
	}
	return
}
