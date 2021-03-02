package typgo

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/typical-go/typical-go/pkg/envkit"
	"github.com/urfave/cli/v2"
)

type (
	// Descriptor describe the project
	Descriptor struct {
		ProjectName    string // By default is same with project folder. Only allowed characters(a-z,A-Z), underscore or dash.
		ProjectVersion string // By default it is 0.0.1
		Environment    EnvLoader
		Tasks          []Tasker
		Stdout         io.Writer
	}
)

// Start typical build-tool
func Start(d *Descriptor) {
	if err := BuildTool(d).Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

// BuildTool app
func BuildTool(d *Descriptor) *cli.App {
	if d.Stdout == nil {
		d.Stdout = os.Stdout
	}

	logger := Logger{
		Stdout:      d.Stdout,
		ProjectName: d.ProjectName,
	}

	if err := setEnv(d.Environment, logger); err != nil {
		logger.Warn("load environment: " + err.Error())
	}

	cli.AppHelpTemplate = appHelpTemplate
	cli.SubcommandHelpTemplate = subcommandHelpTemplate

	app := cli.NewApp()
	for _, task := range d.Tasks {
		app.Commands = append(app.Commands, task.Task().CliCommand(d))
	}

	return app
}

func setEnv(envLoad EnvLoader, logger Logger) error {
	if envLoad == nil {
		return nil
	}
	env, err := envLoad.EnvLoad()
	if err != nil {
		return err
	}
	if err := envkit.Setenv(env); err != nil {
		return err
	}
	logger.Info("load environment")
	keys := envkit.SortedKeys(env)
	fmt.Fprintln(logger.Stdout, strings.Join(keys, ", "))
	return nil
}
