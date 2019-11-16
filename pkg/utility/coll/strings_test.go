package coll_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/utility/coll"
)

func TestStrings(t *testing.T) {
	var coll coll.Strings
	coll.Add("hello")
	coll.Add("world")
	require.EqualValues(t, []string{"hello", "world"}, coll)
}
