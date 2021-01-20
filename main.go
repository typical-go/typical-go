package main

import (
	"log"

	"github.com/typical-go/typical-go/internal/app"
)

func main() {
	if err := app.Main(); err != nil {
		log.Fatal(err.Error())
	}
}
