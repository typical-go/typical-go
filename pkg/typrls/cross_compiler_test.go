package typrls_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/typical-go/typical-go/pkg/typrls"
)

func TestCrossCompile(t *testing.T) {

	testcases := []struct {
		TestName string
		typrls.CrossCompiler
		TagName         string
		RunExpectations []*typgo.RunExpectation
		ExpectedErr     string
	}{
		{
			CrossCompiler: typrls.CrossCompiler{
				Targets: []typrls.Target{"darwin/amd64", "linux/amd64"},
			},
			TagName: "v0.0.1",
			RunExpectations: []*typgo.RunExpectation{
				{CommandLine: "go build -ldflags \"-X github.com/typical-go/typical-go/pkg/typgo.ProjectName=some-project -X github.com/typical-go/typical-go/pkg/typgo.ProjectVersion=v0.0.1\" -o /some-project_v0.0.1_darwin_amd64 ./cmd/some-project"},
				{CommandLine: "go build -ldflags \"-X github.com/typical-go/typical-go/pkg/typgo.ProjectName=some-project -X github.com/typical-go/typical-go/pkg/typgo.ProjectVersion=v0.0.1\" -o /some-project_v0.0.1_linux_amd64 ./cmd/some-project"},
			},
		},
		{
			TestName: "go build error",
			CrossCompiler: typrls.CrossCompiler{
				Targets: []typrls.Target{"darwin/amd64"},
			},
			TagName: "v0.0.1",
			RunExpectations: []*typgo.RunExpectation{
				{
					CommandLine: "go build -ldflags \"-X github.com/typical-go/typical-go/pkg/typgo.ProjectName=some-project -X github.com/typical-go/typical-go/pkg/typgo.ProjectVersion=v0.0.1\" -o /some-project_v0.0.1_darwin_amd64 ./cmd/some-project",
					ReturnError: errors.New("some-error"),
				},
			},
			ExpectedErr: "some-error",
		},
	}
	for _, tt := range testcases {
		t.Run(tt.TestName, func(t *testing.T) {
			unpatch := typgo.PatchBash(tt.RunExpectations)
			defer unpatch(t)
			c, _ := typgo.DummyContext()
			err := tt.Release(&typrls.Context{
				TagName: tt.TagName,
				Context: c,
			})
			if tt.ExpectedErr != "" {
				require.EqualError(t, err, tt.ExpectedErr)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestTarget(t *testing.T) {
	testcases := []struct {
		TestName string
		typrls.Target
		ExpectedOS   string
		ExpectedArch string
	}{
		{Target: "darwin/amd64", ExpectedOS: "darwin", ExpectedArch: "amd64"},
		{Target: "linux/amd64", ExpectedOS: "linux", ExpectedArch: "amd64"},
		{Target: "no-slash", ExpectedOS: "", ExpectedArch: ""},
	}
	for _, tt := range testcases {
		t.Run(tt.TestName, func(t *testing.T) {
			require.Equal(t, tt.ExpectedOS, tt.OS())
			require.Equal(t, tt.ExpectedArch, tt.Arch())
		})
	}
}
