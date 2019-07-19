package restapp

import (
	"github.com/typical-go/typical-go/command/stmt"
	"github.com/typical-go/typical-go/typicore"
)

type RestAppArchetype struct {
	Metadata typicore.ContextMetadata
}

func (a *RestAppArchetype) Name() string {
	return "rest-app"
}

func (a *RestAppArchetype) Statements() []typicore.Statement {
	return []typicore.Statement{
		stmt.MakeDirectory{Path: a.Metadata.ProjectPath + "/app/controller"},
	}
}
