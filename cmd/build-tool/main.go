package main

import (
	_ "github.com/typical-go/typical-go/cmd/internal/dependency"
	"github.com/typical-go/typical-go/pkg/typicmd/buildtool"
	"github.com/typical-go/typical-go/typical"
)

func main() {
	buildtool.Run(typical.Context)
}
