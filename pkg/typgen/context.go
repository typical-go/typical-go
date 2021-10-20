package typgen

import "github.com/typical-go/typical-go/pkg/typgo"

type (
	Context struct {
		*typgo.Context
		*InitFile
		Annotations []*Annotation
	}
)
