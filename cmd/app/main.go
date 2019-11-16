package main

import (
	_ "github.com/typical-go/typical-go/internal/dependency"
	"github.com/typical-go/typical-go/pkg/typapp"
	"github.com/typical-go/typical-go/typical"
)

func main() {
	typapp.Run(typical.Context)
}
