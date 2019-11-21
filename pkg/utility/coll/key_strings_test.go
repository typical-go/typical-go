package coll_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/utility/coll"
)

func TestKeyStrings_Append(t *testing.T) {
	testcases := []struct {
		*coll.KeyStrings
		i []coll.KeyString
	}{
		{
			KeyStrings: new(coll.KeyStrings).Append(
				coll.KeyString{"key-1", "string-1"},
				coll.KeyString{"key-2", "string-2"},
				coll.KeyString{"key-3", "string-3"},
			),
			i: []coll.KeyString{
				coll.KeyString{"key-1", "string-1"},
				coll.KeyString{"key-2", "string-2"},
				coll.KeyString{"key-3", "string-3"},
			},
		},
		{
			KeyStrings: new(coll.KeyStrings).
				Append(coll.KeyString{"key-1", "string-1"}).
				Append(coll.KeyString{"key-2", "string-2"}).
				Append(coll.KeyString{"key-3", "string-3"}),
			i: []coll.KeyString{
				coll.KeyString{"key-1", "string-1"},
				coll.KeyString{"key-2", "string-2"},
				coll.KeyString{"key-3", "string-3"},
			},
		},
	}
	for i, tt := range testcases {
		require.EqualValues(t, tt.i, *tt.KeyStrings, i)
	}
}

func TestKeyString_SimpleFormat(t *testing.T) {
	testcases := []struct {
		coll.KeyString
		sep    string
		result string
	}{
		{
			KeyString: coll.KeyString{Key: "key", String: "string"},
			sep:       "+",
			result:    "key+string",
		},
		{
			KeyString: coll.KeyString{Key: "hello", String: "world"},
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
		coll.KeyString
		formatter func(key, s string) string
		result    string
	}{
		{
			KeyString: coll.KeyString{Key: "hello", String: "world"},
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
