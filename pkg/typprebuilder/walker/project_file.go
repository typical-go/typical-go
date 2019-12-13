package walker

// ProjectFile of walk analysis
type ProjectFile struct {
	Name string
	Mock bool
}

// IsEmpty is true if empty truct
func (f *ProjectFile) IsEmpty() bool {
	return !f.Mock
}
