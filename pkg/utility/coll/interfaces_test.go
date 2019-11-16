package coll_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/utility/coll"
)

func TestInterfaces(t *testing.T) {
	var coll coll.Interfaces
	coll.Add("some-item")
	coll.Add(88)
	coll.Add(3.14)
	require.EqualValues(t, []interface{}{"some-item", 88, 3.14}, coll)
}
