package main

import (
	"log"

	_ "github.com/typical-go/typical-go/examples/typmock-sample/internal/generated/typical"
	"github.com/typical-go/typical-go/examples/typmock-sample/internal/helloworld"
	"github.com/typical-go/typical-go/pkg/typapp"
)

func main() {
	if err := typapp.Run(helloworld.Main); err != nil {
		log.Fatal(err)
	}
}
