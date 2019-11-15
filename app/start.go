package app

import "os"

// Start the application
func Start() {
	parentPath := "sample"
	packageName := "github.com/typical-go/hello-world"
	os.RemoveAll(parentPath)
	NewProject(parentPath, packageName)
}
