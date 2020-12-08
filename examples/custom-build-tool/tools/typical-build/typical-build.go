package main

import (
	"context"
	"os"

	"github.com/typical-go/typical-go/pkg/typgo"
)

func main() {
	ctx := context.Background()
	output := "bin/custom-build-tool"
	mainPackage := "./cmd/custom-build-tool"
	typgo.RunBash(ctx, &typgo.GoBuild{MainPackage: mainPackage, Output: output})
	typgo.RunBash(ctx, &typgo.Bash{Name: output, Stdout: os.Stdout, Stderr: os.Stderr})
}
