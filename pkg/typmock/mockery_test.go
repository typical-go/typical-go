package typmock_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typannot"
	"github.com/typical-go/typical-go/pkg/typmock"
)

func TestMockery(t *testing.T) {
	mockery := typmock.NewMockery("")
	mockery.Put(&typannot.Mock{Pkg: "pkg1", Source: "target1"})
	mockery.Put(&typannot.Mock{Pkg: "pkg1", Source: "target2"})
	mockery.Put(&typannot.Mock{Pkg: "pkg2", Source: "target3"})
	mockery.Put(&typannot.Mock{Pkg: "pkg1", Source: "target4"})
	mockery.Put(&typannot.Mock{Pkg: "pkg1", Source: "target5"})
	mockery.Put(&typannot.Mock{Pkg: "pkg2", Source: "target6"})

	pkg1 := []*typannot.Mock{
		{Pkg: "pkg1", Source: "target1"},
		{Pkg: "pkg1", Source: "target2"},
		{Pkg: "pkg1", Source: "target4"},
		{Pkg: "pkg1", Source: "target5"},
	}

	pkg2 := []*typannot.Mock{
		{Pkg: "pkg2", Source: "target3"},
		{Pkg: "pkg2", Source: "target6"},
	}

	require.Equal(t, map[string][]*typannot.Mock{"pkg1": pkg1, "pkg2": pkg2}, mockery.TargetMap())
	require.Equal(t, map[string][]*typannot.Mock{"pkg1": pkg1, "pkg2": pkg2}, mockery.TargetMap("pkg1", "pkg2"))
	require.Equal(t, map[string][]*typannot.Mock{"pkg1": pkg1}, mockery.TargetMap("pkg1"))
	require.Equal(t, map[string][]*typannot.Mock{}, mockery.TargetMap("not-found"))

}
