package restful

import (
	"github.com/typical-go/typical-go/command/stmt"
	"github.com/typical-go/typical-go/pkg/runn"
	"github.com/typical-go/typical-go/typibuild"

	"github.com/typical-go/typical-rest-server/EXPERIMENTAL/typictx"
)

type Archetype struct {
}

// Modules return list of module
func (a Archetype) Modules() (module []typictx.Module) {
	return
}

// Configs return list of configs
func (a Archetype) Configs() (configs []typictx.Config) {
	return
}

// Install archetype to the project
func (a Archetype) Install(proj *typibuild.Project) (err error) {
	appPath := proj.Path + "/app"
	return runn.Execute(
		stmt.MakeDirectory{Path: appPath + "/controller"},
		stmt.MakeDirectory{Path: appPath + "/service"},
		stmt.MakeDirectory{Path: appPath + "/helper"},
		stmt.MakeDirectory{Path: appPath + "/repository"},
	)
}
