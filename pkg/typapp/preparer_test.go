package typapp_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typapp"
)

func TestPreparation(t *testing.T) {
	prep := &typapp.Preparation{}
	require.Equal(t, []*typapp.Preparation{prep}, prep.Preparations())
}

func TestPrepares(t *testing.T) {
	prep1 := &typapp.Preparation{}
	prep2 := &typapp.Preparation{}
	prep3 := &typapp.Preparation{}
	prep4 := &typapp.Preparation{}
	prepares := typapp.Preparers{
		prep1,
		prep2,
		typapp.Preparers{
			prep3,
			prep4,
		},
	}
	require.Equal(t, []*typapp.Preparation{
		prep1,
		prep2,
		prep3,
		prep4,
	}, prepares.Preparations())
}
