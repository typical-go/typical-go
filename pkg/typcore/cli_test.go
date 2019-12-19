package typcore_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/typical-go/typical-go/pkg/typcore"
)

func TestNewCli(t *testing.T) {
	desc := &typcore.ProjectDescriptor{}
	obj := struct{}{}
	cli := typcore.NewCli(desc, obj)
	require.Equal(t, desc, cli.ProjectDescriptor())
	require.Equal(t, obj, cli.Object())
}
