package app

import (
	"strings"

	"github.com/typical-go/typical-go/app/stmt"
	"github.com/typical-go/typical-go/pkg/utility/runn"
)

// InitProject iniate new project
func InitProject(parentPath, packageName string) (err error) {
	name := getNameFromPath(packageName)
	projectPath := parentPath + "/" + name

	return runn.Execute(
		stmt.MakeDirectory{Path: projectPath},
		stmt.MakeDirectory{Path: projectPath + "/app"},
		stmt.MakeDirectory{Path: projectPath + "/cmd"},
		stmt.MakeDirectory{Path: projectPath + "/cmd/typical-app"},
		stmt.MakeDirectory{Path: projectPath + "/cmd/typical-dev-tool"},
		stmt.MakeDirectory{Path: projectPath + "/config"},
		stmt.MakeDirectory{Path: projectPath + "/typical"},
		stmt.MakeDirectory{Path: projectPath + "/.typical"},
		// stmt.CreateTypicalContext{Target: projectPath + "/typical/context.go"},
		// stmt.CreateAppEntryPoint{PackageName: packageName, Target: projectPath + "/cmd/typical-app/main.go"},
		// stmt.CreateTypicalDevToolEntryPoint{PackageName: packageName, Target: projectPath + "/cmd/typical-dev-tool/main.go"},
		stmt.CreateTypicalWrapper{Target: projectPath + "/typicalw"},
		stmt.ChangeDirectory{ProjectPath: projectPath},
		stmt.GoModInit{ProjectPath: projectPath, PackageName: packageName},
		stmt.GoFmt{ProjectPath: projectPath},
	)
}

func getNameFromPath(path string) string {
	// TODO: handle window path format
	chunks := strings.Split(path, "/")

	return chunks[len(chunks)-1]
}
