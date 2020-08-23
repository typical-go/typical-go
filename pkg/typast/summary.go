package typast

type (
	// Summary responsible to store filename, declaration and annotation
	Summary struct {
		Paths  []string
		Decls  []*Decl
		Annots []*Annot
	}
	// CheckFn check function
	CheckFn func(*Annot) bool
)

// AddDecl add declaration
func (s *Summary) AddDecl(file File, declType Type) {
	decl := &Decl{
		File: file,
		Type: declType,
	}
	s.Decls = append(s.Decls, decl)
	s.Annots = append(s.Annots, retrieveAnnots(decl)...)
}

// FindAnnot find annot
func (s *Summary) FindAnnot(checkFn CheckFn) []*Annot {
	var annots []*Annot
	for _, annot := range s.Annots {
		if checkFn(annot) {
			annots = append(annots, annot)
		}
	}
	return annots
}
