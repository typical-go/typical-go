package stmt

import (
	"encoding/json"
	"io/ioutil"

	"github.com/typical-go/typical-go/appx"
)

// SaveContext save context to corresponding source
func SaveContext(context appx.Context, source string) Statement {
	return &saveContext{
		source:  source,
		context: context,
	}
}

type saveContext struct {
	source  string
	context appx.Context
}

func (c saveContext) Run() error {
	b, _ := json.MarshalIndent(c.context, "", "    ")
	return ioutil.WriteFile(c.source, b, 0644)
}
