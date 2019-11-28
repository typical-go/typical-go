package typprebuilder

import (
	"fmt"

	"github.com/typical-go/typical-go/pkg/typbuildtool"
	"github.com/typical-go/typical-go/pkg/typcfg"
	"github.com/typical-go/typical-go/pkg/utility/coll"

	log "github.com/sirupsen/logrus"

	"github.com/typical-go/typical-go/pkg/typctx"
	"github.com/typical-go/typical-go/pkg/typenv"
	"github.com/typical-go/typical-go/pkg/typprebuilder/walker"
)

type prebuilder struct {
	ProjectFiles       *walker.ProjectFiles
	Dirs               coll.Strings
	ApplicationImports coll.KeyStrings
	ContextImport      string
	ConfigFields       []typcfg.Field
	BuildCommands      []string
}

func (p *prebuilder) Initiate(ctx *typctx.Context) (err error) {
	var files coll.Strings
	if p.Dirs, files, err = scanProject(typenv.Layout.App); err != nil {
		return
	}
	if p.ProjectFiles, err = walker.WalkProject(files); err != nil {
		return
	}
	p.ContextImport = ctx.Package + "/typical"
	log.Debug("Create imports for Application")
	for _, dir := range p.Dirs {
		p.ApplicationImports.Append(coll.KeyString{String: ctx.Package + "/" + dir})
	}
	p.ApplicationImports.Append(coll.KeyString{String: p.ContextImport})
	p.ConfigFields = ConfigFields(ctx)
	for _, command := range typbuildtool.ModuleCommands(ctx) {
		for _, subcommand := range command.Subcommands {
			s := fmt.Sprintf("%s_%s", command.Name, subcommand.Name)
			p.BuildCommands = append(p.BuildCommands, s)
		}
	}
	return
}
