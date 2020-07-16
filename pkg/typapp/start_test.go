package typapp_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typapp"
)

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
