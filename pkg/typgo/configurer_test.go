package typgo_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typgo"
)

func TestConfigurers(t *testing.T) {
	cfg1 := &typgo.Configuration{}
	cfg2 := &typgo.Configuration{}
	cfg3 := &typgo.Configuration{}
	cfg4 := &typgo.Configuration{}
	configs := typgo.Configurers{
		cfg1,
		cfg2,
		typgo.Configurers{
			cfg3,
			cfg4,
		},
	}
	require.Equal(t,
		[]*typgo.Configuration{cfg1, cfg2, cfg3, cfg4},
		configs.Configurations(),
	)

}
