package typicore

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type ArcheType struct {
	Source  string
	Version string
	Package string
}

func (t *ArcheType) Download(parentPath string) (err error) {
	url := fmt.Sprintf("https://raw.githubusercontent.com/%s/%s/%s/prepare.go",
		t.Source, t.Version, t.Package)

	return downloadFile(parentPath+"/prepare.go", url)
}

func downloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
