package stmt

import (
	"io/ioutil"
)

type CreateEntryPoint struct {
	Source string
}

func (c CreateEntryPoint) Run() error {
	return ioutil.WriteFile(c.Source, []byte(`package main
func main(){

}`), 0644)
}
