package tmplkit

import (
	"sort"
	"strings"
)

type (
	// Imports ...
	Imports struct {
		Map                map[string]string // key is import package , value is import alias
		lastGeneratedAlias string
	}
)

// NewImportAliases return new constructor of ImportAliasess
func NewImports(m map[string]string) *Imports {
	if m == nil {
		m = make(map[string]string)
	}
	return &Imports{
		Map: m,
	}
}

// AppendWithAlias Append package with generated alias name. Return generated alias name
func (i *Imports) AppendWithAlias(pkg string) string {
	alias, ok := i.Map[pkg]
	if !ok {
		alias = nextAlias(i.lastGeneratedAlias)
		i.Map[pkg] = alias
		i.lastGeneratedAlias = alias
		return alias
	}
	return alias
}

func (i Imports) Keys() []string {
	var keys []string
	for k := range i.Map {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func (i Imports) String() string {
	var o strings.Builder
	if len(i.Map) > 0 {
		o.WriteString("import (\n")
		for _, key := range i.Keys() {
			o.WriteString("\t")
			if i.Map[key] != "" {
				o.WriteString(i.Map[key])
				o.WriteString(" ")
			}
			o.WriteString("\"")
			o.WriteString(key)
			o.WriteString("\"\n")
		}
		o.WriteString(")\n\n")
	}
	return o.String()
}

func nextAlias(last string) string {
	if last == "" {
		return "a"
	} else if last[len(last)-1] == 'z' {
		return last[:len(last)-1] + "aa"
	} else {
		return last[:len(last)-1] + string(last[len(last)-1]+1)
	}
}
