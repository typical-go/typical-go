package stdbuild

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/typical-go/typical-go/pkg/typenv"
	"github.com/urfave/cli/v2"
)

// CmdClean is command to  clean
func CmdClean() *cli.Command {
	return &cli.Command{
		Name:    "clean",
		Aliases: []string{"c"},
		Usage:   "Clean the project from generated file during build time",
		Action:  cleanProject,
	}
}

func cleanProject(c *cli.Context) error {
	removeFile(typenv.AppBin)
	removeAllFile(typenv.Layout.Temp)
	removeFile(".env")
	removeFile(typenv.GeneratedConstructor)
	return nil
}

func removeFile(name string) {
	log.Infof("Remove: %s", name)
	if err := os.Remove(name); err != nil {
		log.Error(err.Error())
	}
}

func removeAllFile(path string) {
	log.Infof("Remove All: %s", path)
	if err := os.RemoveAll(path); err != nil {
		log.Error(err.Error())
	}
}
