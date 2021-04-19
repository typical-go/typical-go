package app

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

// Start app
func Start() {
	// TODO: change start app implementation
	fmt.Println("Hello world!")
	fmt.Print("Press 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

// Stop app
func Stop() {
	// TODO: change graceful shutdown implementation
	fmt.Printf("Stop app at %s", time.Now())
}
