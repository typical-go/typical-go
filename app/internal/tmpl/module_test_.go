package tmpl

// ModuleTest template
const ModuleTest = `package {{.Name}}

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typcore"
)

func TestModule(t *testing.T) {
	m := &Module{}
	require.True(t, typcore.IsProvider(m))
	require.True(t, typcore.IsDestroyer(m))
	require.True(t, typcore.IsConfigurer(m))
	require.True(t, typcore.IsBuildCommander(m))
}
`
