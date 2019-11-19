package typbuildtool

import (
	"os"

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
		Subcommands: []cli.Command{
			{Name: "build-tool", Usage: "Remove build-tool", Action: t.removeBuildTool},
			{Name: "prebuilder", Usage: "Remove prebuilder", Action: t.removePrebuilder},
			{Name: "app", Usage: "Remove app", Action: t.removeApp},
			{Name: "metadata", Usage: "Remove metadata", Action: t.removeMetadata},
			{Name: "env", Usage: "Remove envfile", Action: t.removeEnvFile},
			{Name: "dependency", Usage: "Remove dependency", Action: t.removeEnvFile},
		},
	}
}

func (t buildtool) cleanProject(ctx *cli.Context) error {
	t.removeBuildTool(ctx)
	t.removePrebuilder(ctx)
	t.removeApp(ctx)
	t.removeMetadata(ctx)
	t.removeEnvFile(ctx)
	t.removeDependency(ctx)
	return nil
}

func (t buildtool) removeBuildTool(ctx *cli.Context) error {
	removeFile(typenv.BuildTool.BinPath)
	return nil
}

func (t buildtool) removePrebuilder(ctx *cli.Context) error {
	removeFile(typenv.Prebuilder.BinPath)
	return nil
}

func (t buildtool) removeApp(ctx *cli.Context) error {
	removeFile(typenv.App.BinPath)
	return nil
}

func (t buildtool) removeMetadata(ctx *cli.Context) error {
	removeAllFile(typenv.Metadata)
	return nil
}

func (t buildtool) removeEnvFile(ctx *cli.Context) error {
	removeFile(".env")
	return nil
}

func (t buildtool) removeDependency(ctx *cli.Context) error {
	removeAllFile(typenv.Dependency.SrcPath)
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
