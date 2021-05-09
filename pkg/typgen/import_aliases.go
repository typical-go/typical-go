package typgen

type (
	// ImportAliases ...
	ImportAliases struct {
		Map  map[string]string // key is import package , value is import alias
		last string
	}
)

// NewImportAliases return new constructor of ImportAliasess
func NewImportAliases() *ImportAliases {
	return &ImportAliases{
		Map: make(map[string]string),
	}
}

// Append package to map
func (i *ImportAliases) Append(pkg string) string {
	alias, ok := i.Map[pkg]
	if !ok {
		alias = nextAlias(i.last)
		i.Map[pkg] = alias
		i.last = alias
		return alias
	}
	return alias
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
