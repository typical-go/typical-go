package typast

import (
	"strings"
)

// Ast responsible to store declaration
type Ast struct {
	decls []*Declaration
}

// Append return DeclStore with appended decls
func (c *Ast) Append(decls ...*Declaration) *Ast {
	c.decls = append(c.decls, decls...)
	return c
}

// EachDecl to handle each declaration
func (c *Ast) EachDecl(fn DeclFunc) (err error) {
	for _, decl := range c.decls {
		if err = fn(decl); err != nil {
			return
		}
	}
	return
}

// EachAnnotation to handle each annotation
func (c *Ast) EachAnnotation(name string, declType DeclType, fn AnnotationFunc) (err error) {
	return c.EachDecl(func(decl *Declaration) (err error) {
		annotation := getAnnotation(decl.Annotations, name)
		if annotation != nil {
			if decl.Type == declType {
				if err = fn(decl, annotation); err != nil {
					return
				}
			} else {
				// log.Warnf("[%s] has no effect to %s:%s", name, declType, decl.SourceName)
				// TODO: give some hint
			}
		}
		return
	})
}

func getAnnotation(a []*Annotation, name string) *Annotation {
	for _, a := range a {
		if strings.ToLower(name) == strings.ToLower(a.Name) {
			return a
		}
	}
	return nil
}
