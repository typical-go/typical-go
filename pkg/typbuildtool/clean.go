package typbuildtool

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/typical-go/typical-go/pkg/typenv"
	"github.com/urfave/cli/v2"
)

func cmdClean() *cli.Command {
	return &cli.Command{
		Name:    "clean",
		Aliases: []string{"c"},
		Usage:   "Clean the project from generated file during build time",
		Action:  cleanProject,
		Subcommands: []*cli.Command{
			{Name: "app", Usage: "Remove app", Action: removeApp},
			{Name: "metadata", Usage: "Remove metadata", Action: removeTmp},
			{Name: "env", Usage: "Remove envfile", Action: removeEnvFile},
			{Name: "constructor", Usage: "Remove constructor", Action: removeConstructor},
		},
	}
}

func cleanProject(c *cli.Context) error {
	removeApp(c)
	removeTmp(c)
	removeEnvFile(c)
	removeConstructor(c)
	return nil
}

func removeApp(c *cli.Context) error {
	removeFile(typenv.AppBin)
	return nil
}

func removeTmp(c *cli.Context) error {
	removeAllFile(typenv.Layout.Temp)
	return nil
}

func removeEnvFile(c *cli.Context) error {
	removeFile(".env")
	return nil
}

func removeConstructor(c *cli.Context) error {
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
