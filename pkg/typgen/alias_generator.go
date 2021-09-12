package typgen

import "sort"

type (
	AliasGenerator struct {
		Map                map[string]string // key is import package , value is import alias
		lastGeneratedAlias string
	}
)

func NewAliasGenerator(m map[string]string) *AliasGenerator {
	if m == nil {
		m = make(map[string]string)
	}
	return &AliasGenerator{
		Map: m,
	}
}

func (i *AliasGenerator) Generate(val string) string {
	alias, ok := i.Map[val]
	if !ok {
		alias = i.next(i.lastGeneratedAlias)
		i.Map[val] = alias
		i.lastGeneratedAlias = alias
		return alias
	}
	return alias
}

func (i *AliasGenerator) next(last string) string {
	if last == "" {
		return "a"
	} else if last[len(last)-1] == 'z' {
		return last[:len(last)-1] + "aa"
	} else {
		return last[:len(last)-1] + string(last[len(last)-1]+1)
	}
}

func (i *AliasGenerator) Keys() []string {
	var keys []string
	for k := range i.Map {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func (i *AliasGenerator) Imports() []*Import {
	var imports []*Import
	for _, key := range i.Keys() {
		imports = append(imports, &Import{
			Name: i.Map[key],
			Path: key,
		})
	}
	return imports
}
