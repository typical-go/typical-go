package command

import (
	"strings"

	"github.com/typical-go/typical-go/archetypes/restful"
	"github.com/typical-go/typical-go/pkg/runn"

	"github.com/typical-go/typical-go/command/stmt"
	"github.com/typical-go/typical-go/typibuild"
)

// NewProject new project
func NewProject(parentPath, packageName string) (err error) {

	name := getNameFromPath(packageName)
	projectPath := parentPath + "/" + name

	// TODO: parse from project.json
	proj := &typibuild.Project{
		Name:        name,
		Version:     "0.0.1",
		Description: "Hello world of typical generation",
		Modules:     []string{},
		PackageName: packageName,
		Path:        projectPath,
	}

	var archetype typibuild.ArcheType

	// TODO: initiate based on project.json
	archetype = restful.Archetype{}

	return runn.Execute(
		stmt.MakeDirectory{Path: proj.Path},
		stmt.MakeDirectory{Path: proj.Path + "/app"},
		stmt.MakeDirectory{Path: proj.Path + "/cmd"},
		stmt.MakeDirectory{Path: proj.Path + "/cmd/typical-app"},
		stmt.MakeDirectory{Path: proj.Path + "/cmd/typical-dev-tool"},
		stmt.MakeDirectory{Path: proj.Path + "/config"},
		stmt.MakeDirectory{Path: proj.Path + "/typical"},
		stmt.MakeDirectory{Path: proj.Path + "/.typical"},
		stmt.CreateTypicalContext{Project: proj, Target: proj.Path + "/typical/context.go"},
		stmt.CreateAppEntryPoint{Project: proj, Target: proj.Path + "/cmd/typical-app/main.go"},
		stmt.CreateTypicalDevToolEntryPoint{Project: proj, Target: proj.Path + "/cmd/typical-dev-tool/main.go"},
		stmt.CreateTypicalWrapper{Target: proj.Path + "/typicalw"},
		stmt.TriggerArchetypeInstall{Project: proj, Archetype: archetype},
		stmt.ChangeDirectory{ProjectPath: proj.Path},
		stmt.GoModInit{ProjectPath: proj.Path, PackageName: packageName},
		stmt.GoFmt{ProjectPath: proj.Path},
	)
}

func getNameFromPath(path string) string {
	// TODO: handle window path format
	chunks := strings.Split(path, "/")

	return chunks[len(chunks)-1]
}
