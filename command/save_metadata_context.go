package command

import (
	"encoding/json"
	"io/ioutil"

	"github.com/typical-go/typical-code-generator/metadata"
)

type saveMetadataContext struct {
	source  string
	context metadata.Context
}

// SaveMetadataContext save metadata (context) to corresponding source
func SaveMetadataContext(source string, context metadata.Context) Command {
	return &saveMetadataContext{
		source:  source,
		context: context,
	}
}

func (c saveMetadataContext) Run() error {
	b, _ := json.MarshalIndent(c.context, "", "    ")
	return ioutil.WriteFile(c.source, b, 0644)
}
