package stmt

import (
	"github.com/typical-go/typical-go/typibuild"
)

type TriggerArchetypeInstall struct {
	Project   *typibuild.Project
	Archetype typibuild.ArcheType
}

func (i TriggerArchetypeInstall) Run() error {
	return i.Archetype.Install(i.Project)
}
