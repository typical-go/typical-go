package typcfg

import "os"

// Load configuration from source file
func Load(source string) (m map[string]string, err error) {
	var (
		file *os.File
	)

	if file, err = os.Open(source); err != nil {
		return
	}
	defer file.Close()

	m = Read(file)
	for key, value := range m {
		if err = os.Setenv(key, value); err != nil {
			return
		}
	}

	return
}
