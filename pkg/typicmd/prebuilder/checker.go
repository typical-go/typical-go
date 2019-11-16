package prebuilder

import "github.com/typical-go/typical-go/pkg/typctx"

type checker struct {
	*typctx.Context
	mockTarget      bool
	configuration   bool
	testTarget      bool
	buildToolBinary bool
	contextChecksum bool
	buildCommands   bool
	readmeFile      bool
}

func (r *checker) checkBuildTool() bool {
	return r.mockTarget ||
		r.configuration ||
		r.testTarget ||
		r.buildToolBinary ||
		r.contextChecksum ||
		r.buildCommands
}

func (r *checker) checkReadme() bool {
	return r.Context.ReadmeGenerator != nil &&
		(r.buildCommands || r.configuration || r.readmeFile)
}
