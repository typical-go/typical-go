package typast

import (
	log "github.com/sirupsen/logrus"
)

// Store responsible to store declaration
type Store struct {
	decls []*Declaration
}

// DeclFunc to handle declaration
type DeclFunc func(*Declaration) error

// AnnotationFunc to handle annotation
type AnnotationFunc func(decl *Declaration, ann *Annotation) error

// Append return DeclStore with appended decls
func (c *Store) Append(decls ...*Declaration) *Store {
	c.decls = append(c.decls, decls...)
	return c
}

// EachDecl to handle each declaration
func (c *Store) EachDecl(fn DeclFunc) (err error) {
	for _, decl := range c.decls {
		if err = fn(decl); err != nil {
			return
		}
	}
	return
}

// EachAnnotation to handle each annotation
func (c *Store) EachAnnotation(name string, declType DeclType, fn AnnotationFunc) (err error) {
	return c.EachDecl(func(decl *Declaration) (err error) {
		annotation := decl.Annotations.Get(name)
		if annotation != nil {
			if decl.Type == declType {
				if err = fn(decl, annotation); err != nil {
					return
				}
			} else {
				log.Warnf("[%s] has no effect to %s:%s", name, declType, decl.SourceName)
			}
		}
		return
	})
}
