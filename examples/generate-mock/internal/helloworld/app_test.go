package helloworld_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/golang/mock/gomock"
	"github.com/typical-go/typical-go/examples/generate-mock/internal/helloworld"
	"github.com/typical-go/typical-go/examples/generate-mock/internal/helloworld_mock"
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
