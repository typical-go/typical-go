package typmock_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typmock"
)

func TestMockery(t *testing.T) {
	mockery := typmock.NewMockery("")
	mockery.Put(&typmock.Target{Pkg: "pkg1", Source: "target1"})
	mockery.Put(&typmock.Target{Pkg: "pkg1", Source: "target2"})
	mockery.Put(&typmock.Target{Pkg: "pkg2", Source: "target3"})
	mockery.Put(&typmock.Target{Pkg: "pkg1", Source: "target4"})
	mockery.Put(&typmock.Target{Pkg: "pkg1", Source: "target5"})
	mockery.Put(&typmock.Target{Pkg: "pkg2", Source: "target6"})

	pkg1 := []*typmock.Target{
		{Pkg: "pkg1", Source: "target1"},
		{Pkg: "pkg1", Source: "target2"},
		{Pkg: "pkg1", Source: "target4"},
		{Pkg: "pkg1", Source: "target5"},
	}

	pkg2 := []*typmock.Target{
		{Pkg: "pkg2", Source: "target3"},
		{Pkg: "pkg2", Source: "target6"},
	}

	require.Equal(t, map[string][]*typmock.Target{"pkg1": pkg1, "pkg2": pkg2}, mockery.TargetMap())
	require.Equal(t, map[string][]*typmock.Target{"pkg1": pkg1, "pkg2": pkg2}, mockery.TargetMap("pkg1", "pkg2"))
	require.Equal(t, map[string][]*typmock.Target{"pkg1": pkg1}, mockery.TargetMap("pkg1"))
	require.Equal(t, map[string][]*typmock.Target{}, mockery.TargetMap("not-found"))

}
