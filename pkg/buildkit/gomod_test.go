package buildkit_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/buildkit"
)

func TestGoMod(t *testing.T) {
	var b bytes.Buffer
	b.WriteString("module github.com/typical-go/typical-go\ngo 1.13")

	gomod := buildkit.ParseGoMod(&b)
	require.Equal(t, &buildkit.GoMod{
		ProjectPackage: "github.com/typical-go/typical-go",
		GoVersion:     "1.13",
	}, gomod)
}
