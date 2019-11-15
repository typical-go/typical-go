package app

import "os"

func start() {
	parentPath := "sample"
	packageName := "github.com/typical-go/hello-world"
	os.RemoveAll(parentPath)
	NewProject(parentPath, packageName)
}
