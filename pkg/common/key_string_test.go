package common_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/common"
)

func TestKeyString_SimpleFormat(t *testing.T) {
	testcases := []struct {
		common.KeyString
		sep    string
		result string
	}{
		{
			KeyString: common.KeyString{Key: "key", String: "string"},
			sep:       "+",
			result:    "key+string",
		},
		{
			KeyString: common.KeyString{Key: "hello", String: "world"},
			sep:       " ",
			result:    "hello world",
		},
	}
	for _, tt := range testcases {
		require.Equal(t, tt.result, tt.SimpleFormat(tt.sep))
	}
}

func TestKeyString_Format(t *testing.T) {
	testcases := []struct {
		common.KeyString
		formatter func(key, s string) string
		result    string
	}{
		{
			KeyString: common.KeyString{Key: "hello", String: "world"},
			formatter: func(key, s string) string {
				return fmt.Sprintf("%s is key; %s is string", key, s)
			},
			result: "hello is key; world is string",
		},
	}
	for _, tt := range testcases {
		require.Equal(t, tt.result, tt.Format(tt.formatter))
	}
}
