package main

import (
	"github.com/typical-go/typical-go/examples/serve-react-demo/internal/server"
	"github.com/typical-go/typical-go/pkg/typapp"
)

func main() {
	typapp.Start(server.Main)
}
