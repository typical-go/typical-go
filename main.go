package main

import "github.com/typical-go/typical-go/app"

func main() {
	parentPath := "sample"
	packageName := "github.com/typical-go/hello-world"
	app.NewProject(parentPath, packageName)
}
