package stmt

import (
	"io/ioutil"
)

var typicalContextTemplate = []byte(`package typical

import(
	"github.com/typical-go/typical-rest-server/EXPERIMENTAL/typictx"
)

// Context instance of Context
var Context = typictx.Context{
	Name:        "",
	Description: "",
}
`)

// CreateTypicalContext to create Typical Context in Target file
type CreateTypicalContext struct {
	Target string
}

// Run the create typical context
func (c CreateTypicalContext) Run() (err error) {
	return ioutil.WriteFile(c.Target, typicalContextTemplate, 0644)
}
