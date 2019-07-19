package command

import (
	"strings"

	"github.com/typical-go/typical-go/command/stmt"
	"github.com/typical-go/typical-go/typicore"
)

// NewProject new project
func NewProject(parentPath, packageName string) (err error) {

	name := getNameFromPath(packageName)
	projectPath := parentPath + "/" + name

	metadata := &typicore.ContextMetadata{
		Name:        name,
		Version:     "0.0.1",
		Description: "Hello world of typical generation",
		ArcheType: "github.com/typical-go/typical-go/EXPERIMENTAL/restapp.RestAppArchetype",
		Modules:     []string{},
		PackageName: packageName,
		ProjectPath: projectPath,
	}

	err = execute(
		stmt.MakeDirectory{Path: projectPath},
		stmt.MakeDirectory{Path: projectPath + "/app"},
		stmt.MakeDirectory{Path: projectPath + "/cmd"},
		stmt.MakeDirectory{Path: projectPath + "/cmd/app"},
		stmt.MakeDirectory{Path: projectPath + "/cmd/typical-dev-tool"},
		stmt.MakeDirectory{Path: projectPath + "/config"},
		stmt.MakeDirectory{Path: projectPath + "/typical"},
		stmt.MakeDirectory{Path: projectPath + "/.typical"},
		stmt.CreateContextMetadata{Metadata: metadata, Target: projectPath + "/.typical/metadata.json"},
		stmt.CreateTypicalContext{Metadata: metadata, Target: projectPath + "/typical/init.go"},
		stmt.CreateAppEntryPoint{Metadata: metadata, Target: projectPath + "/cmd/app/main.go"},
		stmt.CreateTypicalDevToolEntryPoint{Metadata: metadata, Target: projectPath + "/cmd/typical-dev-tool/main.go"},
		stmt.GoModInit{ProjectPath: projectPath, PackageName: packageName},
		stmt.GoFmt{ProjectPath: projectPath},
	)
	return
}

func getNameFromPath(path string) string {
	// TODO: handle window path format
	chunks := strings.Split(path, "/")

	return chunks[len(chunks)-1]
}
