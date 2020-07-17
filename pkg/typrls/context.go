package typrls

import (
	"github.com/typical-go/typical-go/pkg/git"
	"github.com/typical-go/typical-go/pkg/typgo"
)

type (
	// Context contain data for release
	Context struct {
		*typgo.Context
		// ReleaseTag is next release tag
		ReleaseTag string
		Git
	}
	// Git detail
	Git struct {
		Status     string
		CurrentTag string
		Logs       []*git.Log
	}
)
