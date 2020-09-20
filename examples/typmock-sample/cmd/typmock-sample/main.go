package main

import (
	_ "github.com/typical-go/typical-go/examples/typmock-sample/internal/generated/typical"
	"github.com/typical-go/typical-go/examples/typmock-sample/internal/helloworld"
	"github.com/typical-go/typical-go/pkg/typapp"
)

func main() {
	typapp.Start(helloworld.Main)
}
