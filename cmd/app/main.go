package main

import (
	_ "github.com/typical-go/typical-go/internal/dependency"
	"github.com/typical-go/typical-go/pkg/typicmd/application"
	"github.com/typical-go/typical-go/typical"
)

func main() {
	application.Run(typical.Context)
}
