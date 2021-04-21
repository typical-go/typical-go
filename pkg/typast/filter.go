package typast

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

type (
	Filter interface {
		IsAllowed(d *Directive) bool
	}
	NewFilter       func(d *Directive) bool
	Filters         []Filter
	TagNameFilter   []string
	FuncFilter      struct{}
	StructFilter    struct{}
	InterfaceFilter struct{}
	PublicFilter    struct{}
)

//
// NewFilter
//

var _ Filter = (NewFilter)(nil)

func (n NewFilter) IsAllowed(d *Directive) bool {
	return n(d)
}

//
// Filters
//

var _ Filter = (Filters)(nil)

func (f Filters) IsAllowed(d *Directive) bool {
	for _, filter := range f {
		if !filter.IsAllowed(d) {
			return false
		}
	}
	return true
}

//
// TagNameFilter
//

var _ Filter = (TagNameFilter)(nil)

func (t TagNameFilter) IsAllowed(d *Directive) bool {
	for _, tagName := range t {
		if strings.EqualFold(tagName, d.TagName) {
			return true
		}
	}
	return false
}

//
// FuncFilter
//

var _ Filter = (*FuncFilter)(nil)

func (*FuncFilter) IsAllowed(d *Directive) bool {
	funcDecl, ok := d.Type.(*FuncDecl)
	return ok && !funcDecl.IsMethod()
}

//
// StructFilter
//

var _ Filter = (*StructFilter)(nil)

func (*StructFilter) IsAllowed(d *Directive) bool {
	_, ok := d.Type.(*StructDecl)
	return ok
}

//
// PublicStructFilter
//

var _ Filter = (*InterfaceFilter)(nil)

func (*InterfaceFilter) IsAllowed(d *Directive) bool {
	_, ok := d.Type.(*InterfaceDecl)
	return ok
}

//
// PublicFilter
//

var _ Filter = (*PublicFilter)(nil)

func (*PublicFilter) IsAllowed(d *Directive) bool {
	rune, _ := utf8.DecodeRuneInString(d.GetName())
	return unicode.IsUpper(rune)
}
