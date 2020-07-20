package main

import (
	"github.com/typical-go/typical-go/examples/mock-command/internal/helloworld"
	"github.com/typical-go/typical-go/pkg/typapp"
)

func main() {
	typapp.Start(helloworld.Main)
}
