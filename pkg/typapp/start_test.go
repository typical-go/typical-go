package typapp_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typapp"
)

func ExampleStart() {
	// append contructor definition
	typapp.AppendCtor(&typapp.Constructor{
		Fn: func() string {
			return "World"
		},
	})

	// append destructor definition
	typapp.AppendDtor(&typapp.Destructor{
		Fn: func() {
			fmt.Println("clean something")
		},
	})

	// start the application
	typapp.Start(func(text string) {
		fmt.Printf("Hello %s\n", text)
	})

	// Output: Hello World
	// clean something
}

func TestStart(t *testing.T) {
	defer typapp.ClearCtors()
	defer typapp.ClearDtors()

	var debugger []string
	typapp.AppendCtor(&typapp.Constructor{
		Fn: func() string { return "success" },
	})
	typapp.AppendDtor(&typapp.Destructor{
		Fn: func() { debugger = append(debugger, "clean") },
	})

	typapp.Start(func(s string) {
		debugger = append(debugger, s)
	})
	require.Equal(t, []string{"success", "clean"}, debugger)
}
