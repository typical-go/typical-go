package stmt

import (
	"github.com/typical-go/typical-go/typicore"
)

type DownloadPrepareFile struct {
	Metadata *typicore.ContextMetadata
}

func (i DownloadPrepareFile) Run() error {

	// goCommand := fmt.Sprintf("%s/bin/go", build.Default.GOROOT)
	// cmd := exec.Command(goCommand, "fmt", "./...")
	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stdout

	return i.Metadata.ArcheType.Download(i.Metadata.ProjectPath)
}
