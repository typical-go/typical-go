package x

import (
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/typical-go/typical-go/app"

	"github.com/typical-go/typical-go/pkg/typcli"
	"github.com/urfave/cli"
)

// Module of x
func Module() interface{} {
	return &module{}
}

type module struct{}

func (module) BuildCommand(c *typcli.ContextCli) cli.Command {
	return cli.Command{
		Name:  "x",
		Usage: "Typical-Go Development Tool",
		Subcommands: []cli.Command{
			{Name: "init-project", Action: testInitProject},
		},
	}
}

func testInitProject(*cli.Context) error {
	parent := "sample"
	pkg := "github.com/typical-go/hello-world"
	log.Infof("Remove: %s", parent)
	os.RemoveAll(parent)
	log.Infof("Init Project: %s", pkg)
	return app.InitProject(parent, pkg)
}
