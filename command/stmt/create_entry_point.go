package stmt

import (
	"io/ioutil"
)

type CreateEntryPoint struct {
	Target string
}

func (c CreateEntryPoint) Run() error {
	return ioutil.WriteFile(c.Target, []byte(`package main
func main(){

}`), 0644)
}
