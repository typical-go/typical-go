package typbuildtool_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typbuildtool"
)

func TestStdMocker(t *testing.T) {
	mocker := typbuildtool.NewMocker()
	mocker.Put(&typbuildtool.MockTarget{MockDir: "pkg1", SrcName: "target1"})
	mocker.Put(&typbuildtool.MockTarget{MockDir: "pkg1", SrcName: "target2"})
	mocker.Put(&typbuildtool.MockTarget{MockDir: "pkg2", SrcName: "target3"})
	mocker.Put(&typbuildtool.MockTarget{MockDir: "pkg1", SrcName: "target4"})
	mocker.Put(&typbuildtool.MockTarget{MockDir: "pkg1", SrcName: "target5"})
	mocker.Put(&typbuildtool.MockTarget{MockDir: "pkg2", SrcName: "target6"})

	require.Equal(t, map[string][]*typbuildtool.MockTarget{
		"pkg1": []*typbuildtool.MockTarget{
			{MockDir: "pkg1", SrcName: "target1"},
			{MockDir: "pkg1", SrcName: "target2"},
			{MockDir: "pkg1", SrcName: "target4"},
			{MockDir: "pkg1", SrcName: "target5"},
		},
		"pkg2": []*typbuildtool.MockTarget{
			{MockDir: "pkg2", SrcName: "target3"},
			{MockDir: "pkg2", SrcName: "target6"},
		},
	}, mocker.TargetMap())
}
