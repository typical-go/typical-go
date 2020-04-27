package typast

import (
	"strings"
)

// DeclStore responsible to store declaration
type DeclStore struct {
	decls []*Decl
}

// Append return DeclStore with appended decls
func (c *DeclStore) Append(decls ...*Decl) *DeclStore {
	c.decls = append(c.decls, decls...)
	return c
}

// EachDecl to handle each declaration
func (c *DeclStore) EachDecl(fn DeclFunc) (err error) {
	for _, decl := range c.decls {
		if err = fn(decl); err != nil {
			return
		}
	}
	return
}

// EachAnnotation to handle each annotation
func (c *DeclStore) EachAnnotation(name string, declType DeclType, fn AnnotFunc) (err error) {
	return c.EachDecl(func(decl *Decl) (err error) {
		annotation := getAnnot(decl.Annots, name)
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

func getAnnot(a []*Annot, name string) *Annot {
	for _, a := range a {
		if strings.ToLower(name) == strings.ToLower(a.Name) {
			return a
		}
	}
	return nil
}
