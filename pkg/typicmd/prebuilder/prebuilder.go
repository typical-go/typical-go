package prebuilder

import (
	"fmt"

	"github.com/typical-go/typical-go/pkg/typimodule"
	"github.com/typical-go/typical-go/pkg/utility/collection"

	"github.com/typical-go/typical-go/pkg/typicmd/buildtool"
	"github.com/typical-go/typical-go/pkg/typicmd/prebuilder/golang"

	log "github.com/sirupsen/logrus"

	"github.com/typical-go/typical-go/pkg/typicmd/prebuilder/walker"
	"github.com/typical-go/typical-go/pkg/typictx"
	"github.com/typical-go/typical-go/pkg/typienv"
)

type prebuilder struct {
	ProjectFiles       *walker.ProjectFiles
	Dirs               collection.Strings
	ApplicationImports golang.Imports
	ContextImport      string
	ConfigFields       []typimodule.ConfigField
	BuildCommands      []string
}

func (p *prebuilder) Initiate(ctx *typictx.Context) (err error) {
	var files collection.Strings
	if p.Dirs, files, err = scanProject(typienv.AppName); err != nil {
		return
	}
	if p.ProjectFiles, err = walker.WalkProject(files); err != nil {
		return
	}
	p.ContextImport = ctx.Package + "/typical"
	log.Debug("Create imports for Application")
	for _, dir := range p.Dirs {
		p.ApplicationImports.AddImport("", ctx.Package+"/"+dir)
	}
	p.ApplicationImports.AddImport("", p.ContextImport)
	p.ConfigFields = ConfigFields(ctx)
	for _, command := range buildtool.Commands(ctx) {
		for _, subcommand := range command.Subcommands {
			s := fmt.Sprintf("%s_%s", command.Name, subcommand.Name)
			p.BuildCommands = append(p.BuildCommands, s)
		}
	}
	return
}
