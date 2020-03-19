package helloworld_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/examples/generate-mock/mock_helloworld"

	"github.com/golang/mock/gomock"
	"github.com/typical-go/typical-go/examples/generate-mock/helloworld"
)

func TestPrint(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var s strings.Builder
	app := helloworld.New().WithWriter(&s)

	greeter := mock_helloworld.NewMockGreeter(ctrl)
	greeter.EXPECT().Greet().Return("some-word")

	app.Print(greeter)
	require.Equal(t, "some-word\n", s.String())

}
