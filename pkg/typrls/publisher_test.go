package typrls_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typrls"
)

func TestPublisher(t *testing.T) {
	testcases := []struct {
		testName string
		typrls.Publisher
		context     *typrls.Context
		expectedErr string
	}{
		{
			Publisher:   typrls.NewPublisher(func(*typrls.Context) error { return errors.New("some-error") }),
			expectedErr: "some-error",
		},
		{
			Publisher: typrls.NewPublisher(func(*typrls.Context) error { return nil }),
		},
		{
			Publisher: typrls.Publishers{
				typrls.NewPublisher(func(*typrls.Context) error { return nil }),
				typrls.NewPublisher(func(*typrls.Context) error { return errors.New("some-error") }),
			},
			expectedErr: "some-error",
		},
		{
			Publisher: typrls.Publishers{
				typrls.NewPublisher(func(*typrls.Context) error { return errors.New("some-error") }),
				typrls.NewPublisher(func(*typrls.Context) error { return nil }),
			},
			expectedErr: "some-error",
		},
		{
			Publisher: typrls.Publishers{
				typrls.NewPublisher(func(*typrls.Context) error { return nil }),
				typrls.NewPublisher(func(*typrls.Context) error { return nil }),
			},
		},
	}

	for _, tt := range testcases {
		t.Run(tt.testName, func(t *testing.T) {
			err := tt.Publish(tt.context)
			if tt.expectedErr != "" {
				require.EqualError(t, err, tt.expectedErr)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
