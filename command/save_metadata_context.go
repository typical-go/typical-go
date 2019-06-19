package command

import (
	"github.com/typical-go/typical-go/context"

	"encoding/json"
	"io/ioutil"
)

type saveMetadataContext struct {
	source  string
	context context.Context
}

// SaveMetadataContext save context to corresponding source
func SaveMetadataContext(context context.Context, source string) Command {
	return &saveMetadataContext{
		source:  source,
		context: context,
	}
}

func (c saveMetadataContext) Run() error {
	b, _ := json.MarshalIndent(c.context, "", "    ")
	return ioutil.WriteFile(c.source, b, 0644)
}
