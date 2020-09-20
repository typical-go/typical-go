package main

import (
	"fmt"

	"github.com/typical-go/typical-go/pkg/typgo"
)

func main() {
	fmt.Println("-------------------------------")
	fmt.Println("typgo.AppName: " + typgo.AppName)
	fmt.Println("typgo.AppVersion: " + typgo.AppVersion)
	fmt.Println("-------------------------------")
	fmt.Println()
	fmt.Println("Hello, Typical")
}
