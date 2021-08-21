package tmplkit

import (
	"fmt"
	"strings"
)

type (
	SourceCode struct {
		Package   string
		Signature Signature
		Imports   *Imports
		LineCodes []fmt.Stringer
	}
	Function struct {
		Name      string
		Params    []Variable
		Returns   []Variable
		LineCodes []fmt.Stringer
	}
	Variable struct {
		Name string
		Type string
	}
	Import struct {
		PackagePath string
		AliasName   string
	}
	LineCode  string
	Comment   string
	Signature struct {
		TagName string
	}
)

//
// LineCode
//

func (l LineCode) String() string { return string(l) }

//
// Comment
//

func (c Comment) String() string { return "// " + string(c) }

//
// Func
//

var _ fmt.Stringer = (*Function)(nil)

func InitFunction(lines []fmt.Stringer) *Function {
	return &Function{
		Name:      "init",
		LineCodes: lines,
	}
}

func (f Function) String() string {
	var o strings.Builder
	o.WriteString("func ")
	o.WriteString(f.Name)

	o.WriteString(paramsString(f.Params))

	if len(f.Returns) > 0 {
		o.WriteString(paramsString(f.Returns))
	}
	o.WriteString("{\n")

	for _, lc := range f.LineCodes {
		o.WriteString("\t")
		o.WriteString(lc.String())
		o.WriteString("\n")
	}
	o.WriteString("}\n")
	return o.String()
}

func paramsString(vars []Variable) string {
	var o strings.Builder
	o.WriteString("(")
	for i, v := range vars {
		if i > 0 {
			o.WriteString(",")
		}
		o.WriteString(v.Name)
		o.WriteString(" ")
		o.WriteString(v.Type)
	}
	o.WriteString(")")
	return o.String()
}

//
// Signature
//

var _ fmt.Stringer = (*Signature)(nil)

func (s Signature) String() string {
	var o strings.Builder
	o.WriteString("/* DO NOT EDIT. This is code generated file")
	if s.TagName != "" {
		o.WriteString(" from '")
		o.WriteString(s.TagName)
		o.WriteString("' annotation")
	}
	o.WriteString(". */")
	return o.String()
}

//
// Source Code
//

func (s SourceCode) String() string {
	var o strings.Builder
	o.WriteString("package ")
	o.WriteString(s.Package)
	o.WriteString("\n\n")

	o.WriteString(s.Signature.String())
	o.WriteString("\n")

	if s.Imports != nil {
		o.WriteString(s.Imports.String())
	}

	for _, line := range s.LineCodes {
		o.WriteString(line.String())
		o.WriteString("\n")
	}

	return o.String()
}
