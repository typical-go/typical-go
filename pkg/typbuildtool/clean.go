package typbuildtool

import (
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	"github.com/typical-go/typical-go/pkg/typenv"
	"github.com/urfave/cli"
)

func (t buildtool) cmdClean() cli.Command {
	return cli.Command{
		Name:      "clean",
		ShortName: "c",
		Usage:     "Clean the project from generated file during build time",
		Action:    t.cleanProject,
	}
}

func (t buildtool) cleanProject(ctx *cli.Context) (err error) {
	log.Info("Clean the application")
	log.Infof("  Remove %s", typenv.Bin)
	if err = os.RemoveAll(typenv.Bin); err != nil {
		return
	}
	log.Infof("  Remove %s", typenv.Metadata)
	if err = os.RemoveAll(typenv.Metadata); err != nil {
		return
	}
	log.Info("  Remove .env")
	os.Remove(".env")
	return filepath.Walk(typenv.Dependency.SrcPath, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			log.Infof("  Remove %s", path)
			return os.Remove(path)
		}
		return nil
	})
}
