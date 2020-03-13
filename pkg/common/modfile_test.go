package common_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/common"
)

func TestModfile(t *testing.T) {
	var b bytes.Buffer
	b.WriteString("module github.com/typical-go/typical-go\ngo 1.13")

	gomod := common.ParseModfile(&b)
	require.Equal(t, &common.Modfile{
		ProjectPackage: "github.com/typical-go/typical-go",
		GoVersion:      "1.13",
	}, gomod)
}
