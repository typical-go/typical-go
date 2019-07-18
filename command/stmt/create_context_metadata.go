package stmt

import (
	"encoding/json"
	"io/ioutil"

	"github.com/typical-go/typical-go/typicore"
)

type CreateContextMetadata struct {
	Metadata *typicore.ContextMetadata
	Target   string
}

func (c CreateContextMetadata) Run() error {
	b, _ := json.MarshalIndent(c.Metadata, "", "    ")
	return ioutil.WriteFile(c.Target, b, 0644)
}
