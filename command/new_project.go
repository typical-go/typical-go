package command

import (
	"strings"

	"github.com/typical-go/typical-go/command/stmt"
	"github.com/typical-go/typical-go/typicore"
)

// NewProject new project
func NewProject(path string) (err error) {

	metadata := &typicore.ContextMetadata{
		Name:        getNameFromPath(path),
		Version:     "0.0.1",
		Description: "Hello world of typical generation",
		AppModule:   "github.com/typical-go/EXPERIMENTAL/typictx.TypiApp",
		Modules:     []string{},
	}

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
		stmt.CreateEntryPoint{Target: path + "/cmd/app/main.go"},
		stmt.CreateEntryPoint{Target: path + "/cmd/typical-dev-tool/main.go"},
		stmt.CreateTypicalContext{Metadata: metadata, Target: path + "/typical/init.go"},
	)
	return
}

func getNameFromPath(path string) string {
	// TODO: handle window path format
	chunks := strings.Split(path, "/")

	return chunks[len(chunks)-1]
}
