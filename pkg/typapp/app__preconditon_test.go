package typapp_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typapp"
	"github.com/typical-go/typical-go/pkg/typcfg"
)

func TestConfigConstructor(t *testing.T) {
	expected := `func() (cfg *typapp_test.config, err error){
		cfg = new(typapp_test.config)
		if err = typcfg.Process("NAME", cfg); err != nil {
			return nil, err
		}
		return  
	}`

	require.Equal(t, expected, typapp.ConfigContructor(typcfg.NewConfiguration("NAME", &config{})))

}

type config struct{}
