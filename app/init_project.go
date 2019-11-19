package app

import (
	"strings"

	"github.com/typical-go/typical-go/pkg/utility/runn"
)

// InitProject iniate new project
func InitProject(parentPath, packageName string) error {
	return initproject{}.Run()
}

// 	name := getNameFromPath(packageName)
// 	projectPath := parentPath + "/" + name

// 	return runn.Execute(
// 		stmt.MakeDirectory{Path: projectPath},
// 		stmt.MakeDirectory{Path: projectPath + "/app"},
// 		stmt.MakeDirectory{Path: projectPath + "/cmd"},
// 		stmt.MakeDirectory{Path: projectPath + "/cmd/typical-app"},
// 		stmt.MakeDirectory{Path: projectPath + "/cmd/typical-dev-tool"},
// 		stmt.MakeDirectory{Path: projectPath + "/config"},
// 		stmt.MakeDirectory{Path: projectPath + "/typical"},
// 		stmt.MakeDirectory{Path: projectPath + "/.typical"},
// 		stmt.CreateTypicalContext{Target: projectPath + "/typical/context.go"},
// 		stmt.CreateAppEntryPoint{PackageName: packageName, Target: projectPath + "/cmd/typical-app/main.go"},
// 		stmt.CreateTypicalDevToolEntryPoint{PackageName: packageName, Target: projectPath + "/cmd/typical-dev-tool/main.go"},
// 		stmt.CreateTypicalWrapper{Target: projectPath + "/typicalw"},
// 		stmt.ChangeDirectory{ProjectPath: projectPath},
// 		stmt.GoModInit{ProjectPath: projectPath, PackageName: packageName},
// 		stmt.GoFmt{ProjectPath: projectPath},
// 	)
// }

type initproject struct {
}

func (i initproject) Run() (err error) {
	return runn.Execute(
		i.generateAppPackage,
		i.generateTypicalContext,
		i.generateCmdPackage,
		i.generateIgnoreFile,
		i.generateTypicalWrapper,
	)
}

func (i initproject) generateAppPackage() (err error) {
	return
}

func (i initproject) generateTypicalContext() (err error) {
	return
}

func (i initproject) generateCmdPackage() (err error) {
	return
}

func (i initproject) generateIgnoreFile() (err error) {
	return
}

func (i initproject) generateTypicalWrapper() (err error) {
	return
}

func getNameFromPath(path string) string {
	// TODO: handle window path format
	chunks := strings.Split(path, "/")
	return chunks[len(chunks)-1]
}
