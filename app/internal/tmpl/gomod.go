package tmpl

// GoMod template
const GoMod = `module {{.Pkg}}

go 1.13

require github.com/typical-go/typical-go v{{.TypicalVersion}}
`
