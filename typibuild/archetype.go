package typibuild

import "github.com/typical-go/typical-rest-server/EXPERIMENTAL/typictx"

type ArcheType interface {
	Modules() []typictx.Module
	Configs() []typictx.Config
	Install(project *Project) error
}
