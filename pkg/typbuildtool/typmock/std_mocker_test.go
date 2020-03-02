package typmock_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typbuildtool/typmock"
)

func TestStdMocker(t *testing.T) {
	mocker := typmock.New()
	mocker.Put(&typmock.Target{MockDir: "pkg1", SrcName: "target1"})
	mocker.Put(&typmock.Target{MockDir: "pkg1", SrcName: "target2"})
	mocker.Put(&typmock.Target{MockDir: "pkg2", SrcName: "target3"})
	mocker.Put(&typmock.Target{MockDir: "pkg1", SrcName: "target4"})
	mocker.Put(&typmock.Target{MockDir: "pkg1", SrcName: "target5"})
	mocker.Put(&typmock.Target{MockDir: "pkg2", SrcName: "target6"})

	require.Equal(t, map[string][]*typmock.Target{
		"pkg1": []*typmock.Target{
			{MockDir: "pkg1", SrcName: "target1"},
			{MockDir: "pkg1", SrcName: "target2"},
			{MockDir: "pkg1", SrcName: "target4"},
			{MockDir: "pkg1", SrcName: "target5"},
		},
		"pkg2": []*typmock.Target{
			{MockDir: "pkg2", SrcName: "target3"},
			{MockDir: "pkg2", SrcName: "target6"},
		},
	}, mocker.TargetMap())
}
