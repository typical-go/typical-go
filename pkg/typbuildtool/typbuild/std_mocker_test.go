package typbuild_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typbuildtool/typbuild"
)

func TestStdMocker(t *testing.T) {
	mocker := typbuild.NewMocker()
	mocker.Put(&typbuild.MockTarget{MockDir: "pkg1", SrcName: "target1"})
	mocker.Put(&typbuild.MockTarget{MockDir: "pkg1", SrcName: "target2"})
	mocker.Put(&typbuild.MockTarget{MockDir: "pkg2", SrcName: "target3"})
	mocker.Put(&typbuild.MockTarget{MockDir: "pkg1", SrcName: "target4"})
	mocker.Put(&typbuild.MockTarget{MockDir: "pkg1", SrcName: "target5"})
	mocker.Put(&typbuild.MockTarget{MockDir: "pkg2", SrcName: "target6"})

	require.Equal(t, map[string][]*typbuild.MockTarget{
		"pkg1": []*typbuild.MockTarget{
			{MockDir: "pkg1", SrcName: "target1"},
			{MockDir: "pkg1", SrcName: "target2"},
			{MockDir: "pkg1", SrcName: "target4"},
			{MockDir: "pkg1", SrcName: "target5"},
		},
		"pkg2": []*typbuild.MockTarget{
			{MockDir: "pkg2", SrcName: "target3"},
			{MockDir: "pkg2", SrcName: "target6"},
		},
	}, mocker.TargetMap())
}
