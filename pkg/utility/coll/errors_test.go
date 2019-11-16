package coll_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/utility/coll"
)

func TestErrors(t *testing.T) {
	var errors coll.Errors
	errors.Add(fmt.Errorf("error1"))
	errors.Add(nil)
	errors.Add(fmt.Errorf("error2"))
	errors.Add(fmt.Errorf("error3"))
	require.Equal(t, "error1; error2; error3", errors.Error())
}