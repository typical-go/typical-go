package main

import (
	"context"
	"os"

	"github.com/typical-go/typical-go/pkg/execkit"
)

func main() {
	ctx := context.Background()
	output := "bin/custom-build-tool"
	mainPackage := "./cmd/custom-build-tool"
	execkit.Run(ctx, &execkit.GoBuild{MainPackage: mainPackage, Output: output})
	execkit.Run(ctx, &execkit.Command{Name: output, Stdout: os.Stdout, Stderr: os.Stderr})
}
