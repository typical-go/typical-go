package stmt

import (
	"encoding/json"
	"io/ioutil"

	"github.com/typical-go/typical-go/typicore"
)

type CreateContextMetadata struct {
	Metadata *typicore.ContextMetadata
	Source   string
}

func (c CreateContextMetadata) Run() error {
	b, _ := json.MarshalIndent(c.Metadata, "", "    ")
	return ioutil.WriteFile(c.Source, b, 0644)
}
