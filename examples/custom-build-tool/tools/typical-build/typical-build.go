package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"

	"os"
)

const (
	output      = "bin/custom-build-tool"
	mainPackage = "./cmd/custom-build-tool"
)

func main() {
	bash(fmt.Sprintf("go build -o %s %s", output, mainPackage))
	bash(fmt.Sprintf("./%s", output))
}

func bash(commandLine string) {
	slices := strings.Fields(commandLine)
	cmd := exec.Command(slices[0], slices[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	log.Fatal(cmd.Run())
}
