package typmock_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typmock"
)

func TestMockery(t *testing.T) {
	mockery := &typmock.Mockery{
		TargetMap: make(typmock.TargetMap),
	}
	mockery.Put(&typmock.Mock{Pkg: "pkg1", Source: "target1"})
	mockery.Put(&typmock.Mock{Pkg: "pkg1", Source: "target2"})
	mockery.Put(&typmock.Mock{Pkg: "pkg2", Source: "target3"})
	mockery.Put(&typmock.Mock{Pkg: "pkg1", Source: "target4"})
	mockery.Put(&typmock.Mock{Pkg: "pkg1", Source: "target5"})
	mockery.Put(&typmock.Mock{Pkg: "pkg2", Source: "target6"})

	pkg1 := []*typmock.Mock{
		{Pkg: "pkg1", Source: "target1"},
		{Pkg: "pkg1", Source: "target2"},
		{Pkg: "pkg1", Source: "target4"},
		{Pkg: "pkg1", Source: "target5"},
	}

	pkg2 := []*typmock.Mock{
		{Pkg: "pkg2", Source: "target3"},
		{Pkg: "pkg2", Source: "target6"},
	}

	require.Equal(t,
		typmock.TargetMap{"pkg1": pkg1, "pkg2": pkg2},
		mockery.Filter(),
	)
	require.Equal(t,
		typmock.TargetMap{"pkg1": pkg1, "pkg2": pkg2},
		mockery.Filter("pkg1", "pkg2"),
	)
	require.Equal(t,
		typmock.TargetMap{"pkg1": pkg1},
		mockery.Filter("pkg1"),
	)
	require.Equal(t,
		typmock.TargetMap{},
		mockery.Filter("not-found"),
	)

}
