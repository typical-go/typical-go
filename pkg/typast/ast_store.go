package typast

import (
	"strings"
)

// ASTStore responsible to store filename, declaration and annotation
type ASTStore struct {
	filenames []string
	decls     []*Decl
}

// EachDecl to handle each declaration
func (c *ASTStore) EachDecl(fn DeclFunc) (err error) {
	for _, decl := range c.decls {
		if err = fn(decl); err != nil {
			return
		}
	}
	return
}

// EachAnnotation to handle each annotation
func (c *ASTStore) EachAnnotation(name string, declType DeclType, fn AnnotFunc) (err error) {
	return c.EachDecl(func(decl *Decl) (err error) {
		annots := ParseAnnots(decl)
		annot := getAnnot(annots, name)
		if annot != nil {
			if decl.Type == declType {
				if err = fn(decl, annot); err != nil {
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

func getAnnot(a []*Annotation, name string) *Annotation {
	for _, a := range a {
		if strings.ToLower(name) == strings.ToLower(a.TagName) {
			return a
		}
	}
	return nil
}
