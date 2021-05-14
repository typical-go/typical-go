package app_test

import (
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/examples/typmock-sample/internal/app"
	"github.com/typical-go/typical-go/examples/typmock-sample/internal/generated/mock/greeter_mock"
)

func TestPrint(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var out strings.Builder

	greeter := greeter_mock.NewMockGreeter(ctrl)
	greeter.EXPECT().Greet().Return("some-word")

	app.Start(&out, greeter)
	require.Equal(t, "some-word\n", out.String())

}
