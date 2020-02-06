package typcore_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/typical-go/typical-go/pkg/typcore"
)

func TestWriteEnv(t *testing.T) {
	var w strings.Builder
	typcore.WriteEnv(&w,
		[]string{"some-name1", "some-name2"},
		typcore.ConfigMap{
			"some-name1": typcore.ConfigDetail{
				Name:    "some-name1",
				Default: "some-default1",
				Value:   "some-value1",
			},
			"some-name2": typcore.ConfigDetail{
				Name:    "some-name2",
				Default: "some-default2",
				Value:   "",
				IsZero:  true,
			},
		})
	require.Equal(t, "some-name1=some-value1\nsome-name2=some-default2\n", w.String())
}
