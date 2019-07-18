package command

import (
	"strings"

	"github.com/typical-go/typical-go/command/stmt"
	"github.com/typical-go/typical-go/typicore"
)

// NewProject new project
func NewProject(path string) (err error) {

	metadata := &typicore.ContextMetadata{
		Name:      getNameFromPath(path),
		Version:   "0.0.1",
		AppModule: "github.com/typical-go/EXPERIMENTAL/typictx.TypiApp",
		Modules:   []string{},
	}

	err = Start(
		&stmt.MakeDirectory{path},
		&stmt.MakeDirectory{path + "/app"},
		&stmt.MakeDirectory{path + "/cmd"},
		&stmt.MakeDirectory{path + "/cmd/app"},
		&stmt.MakeDirectory{path + "/cmd/typical-dev-tool"},
		&stmt.MakeDirectory{path + "/config"},
		&stmt.MakeDirectory{path + "/typical"},
		&stmt.MakeDirectory{path + "/.typical"},
		&stmt.CreateContextMetadata{metadata, path + "/.typical/metadata.json"},
		&stmt.CreateEntryPoint{path + "/cmd/app/main.go"},
		&stmt.CreateEntryPoint{path + "/cmd/typical-dev-tool/main.go"},
	)
	return
}

func getNameFromPath(path string) string {
	// TODO: handle window path format
	chunks := strings.Split(path, "/")

	return chunks[len(chunks)-1]
}
