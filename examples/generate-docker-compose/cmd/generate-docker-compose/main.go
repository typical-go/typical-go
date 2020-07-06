package main

import (
	"github.com/typical-go/typical-go/examples/generate-docker-compose/internal/pinger"
	"github.com/typical-go/typical-go/pkg/typapp"
)

func main() {
	typapp.Start(pinger.Main)
}
