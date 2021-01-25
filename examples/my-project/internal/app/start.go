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

// Shutdown app
func Shutdown() {
	// TODO: change graceful shutdown implementation
	fmt.Printf("Shutdown app at %s", time.Now())
}
