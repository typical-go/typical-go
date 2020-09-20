package helloworld_test

import (
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/examples/typmock-sample/internal/helloworld"
	"github.com/typical-go/typical-go/examples/typmock-sample/internal/helloworld_mock"
)

func TestPrint(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var debugger strings.Builder

	greeter := helloworld_mock.NewMockGreeter(ctrl)
	greeter.EXPECT().Greet().Return("some-word")

	helloworld.Main(greeter, &debugger)
	require.Equal(t, "some-word\n", debugger.String())

}
