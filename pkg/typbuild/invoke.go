package typbuild

import (
	log "github.com/sirupsen/logrus"
	"github.com/typical-go/typical-go/pkg/common"
	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/urfave/cli/v2"
	"go.uber.org/dig"
)

// Invoke function
func (b *Build) Invoke(bctx *typcore.BuildContext, c *cli.Context, fn interface{}) (err error) {
	di := dig.New()

	// provide the cli.Context
	if err = di.Provide(func() *cli.Context { return c }); err != nil {
		return
	}

	// provide functions
	if bctx.Configuration != nil {
		if err = provide(di, bctx.Configuration.Provide()...); err != nil {
			return
		}
	}

	startFn := func() error {
		return di.Invoke(fn)
	}
	errs := common.NewApplication(startFn).Run()
	for _, err := range errs {
		log.Error(err.Error())
	}
	return
}

func invoke(di *dig.Container, fns ...interface{}) (err error) {
	for _, fn := range fns {
		if err = di.Invoke(fn); err != nil {
			return
		}
	}
	return
}

func provide(di *dig.Container, fns ...interface{}) (err error) {
	for _, fn := range fns {
		if err = di.Provide(fn); err != nil {
			return
		}
	}
	return
}
