package typgen

import "os"

type (
	SourceCoder interface {
		SourceCode() string
	}
	Comment string
)

func WriteSourceCode(filename string, srcs ...SourceCoder) error {
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		return err
	}

	for _, src := range srcs {
		f.WriteString(src.SourceCode())
		f.WriteString("\n")
	}
	return f.Close()
}

func (c Comment) SourceCode() string {
	return "// " + string(c)
}
