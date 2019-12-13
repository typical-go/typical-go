package walker

// ProjectFiles information
type ProjectFiles []ProjectFile

// Add item to files
func (f *ProjectFiles) Add(item ProjectFile) {
	*f = append(*f, item)
}

// Automocks return automocked filenames
func (f *ProjectFiles) Automocks() (filenames []string) {
	for _, file := range *f {
		if file.Mock {
			filenames = append(filenames, file.Name)
		}
	}
	return
}
