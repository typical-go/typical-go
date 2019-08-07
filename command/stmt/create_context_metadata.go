package stmt

import (
	"encoding/json"
	"io/ioutil"

	"github.com/typical-go/typical-go/typibuild"
)

type CreateContextMetadata struct {
	Project *typibuild.Project
	Target  string
}

func (c CreateContextMetadata) Run() error {
	b, _ := json.MarshalIndent(c.Project, "", "    ")
	return ioutil.WriteFile(c.Target, b, 0644)
}
