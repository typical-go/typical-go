package typbuild

import (
	"os"

	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/urfave/cli/v2"
)

// Run build tool
func (b *Build) Run(bctx *typcore.BuildContext) (err error) {
	app := cli.NewApp()
	app.Name = bctx.Name
	app.Usage = "" // NOTE: intentionally blank
	app.Description = bctx.Description
	app.Version = bctx.Version
	app.Commands = b.BuildCommands(&Context{bctx})

	return app.Run(os.Args)
}
