package generator_test

import (
	"testing"

	"github.com/typical-go/typical-code-generator/generator"
)

func TestReadme(t *testing.T) {
	readme := generator.ReadMe{
		ProjectTitle:       "Some Title",
		ProjectDescription: "Some Description",
	}

	readme.Generate("path")
}
