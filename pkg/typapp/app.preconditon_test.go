package typapp

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typcfg"
)

func TestConfigConstructor(t *testing.T) {
	require.Equal(t, `func() (cfg *typapp.App, err error){
		cfg = new(typapp.App)
		if err = typcfg.Process("NAME", cfg); err != nil {
			return nil, err
		}
		return  
	}`, cfgCtorDef(typcfg.NewConfiguration("NAME", &App{})))

}
