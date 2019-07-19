package command

import (
	"strings"

	"github.com/typical-go/typical-go/command/stmt"
	"github.com/typical-go/typical-go/typicore"
)

// NewProject new project
func NewProject(parentPath, projectPath string) (err error) {

	metadata := &typicore.ContextMetadata{
		Name:        getNameFromPath(projectPath),
		Version:     "0.0.1",
		Description: "Hello world of typical generation",
		AppModule:   "github.com/typical-go/EXPERIMENTAL/typictx.TypiApp",
		Modules:     []string{},
		ProjectPath: projectPath,
	}

	path := parentPath + "/" + metadata.Name

	err = Start(
		stmt.MakeDirectory{Path: path},
		stmt.MakeDirectory{Path: path + "/app"},
		stmt.MakeDirectory{Path: path + "/cmd"},
		stmt.MakeDirectory{Path: path + "/cmd/app"},
		stmt.MakeDirectory{Path: path + "/cmd/typical-dev-tool"},
		stmt.MakeDirectory{Path: path + "/config"},
		stmt.MakeDirectory{Path: path + "/typical"},
		stmt.MakeDirectory{Path: path + "/.typical"},
		stmt.CreateContextMetadata{Metadata: metadata, Target: path + "/.typical/metadata.json"},
		// stmt.CreateEntryPoint{Tarath: path + "/cmd/typical-dev-tool/main.go"},
		stmt.CreateTypicalContext{Metadata: metadata, Target: path + "/typical/init.go"},
		stmt.CreateAppEntryPoint{Metadata: metadata, Target: path + "/cmd/app/main.go"},
	)
	return
}

func getNameFromPath(path string) string {
	// TODO: handle window path format
	chunks := strings.Split(path, "/")

	return chunks[len(chunks)-1]
}
