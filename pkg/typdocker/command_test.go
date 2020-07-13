package typdocker_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typdocker"
	"github.com/typical-go/typical-go/pkg/typgo"
)

var (
	redisV2 = &typdocker.Recipe{
		Version: "2",
		Services: typdocker.Services{
			"redis":  "redis-service",
			"webdis": "webdis-service",
		},
		Networks: typdocker.Networks{
			"webdis": "webdis-network",
		},
		Volumes: typdocker.Volumes{
			"redis": "redis-volume",
		},
	}
	pgV2 = &typdocker.Recipe{
		Version: "2",
		Services: typdocker.Services{
			"pg": "pg-service",
		},
		Networks: typdocker.Networks{
			"pg": "pg-network",
		},
		Volumes: typdocker.Volumes{
			"pg": "pg-volume",
		},
	}
)

func TestComposeRecipe(t *testing.T) {
	testcases := []struct {
		testName  string
		version   string
		composers []typdocker.Composer
		expected  string
	}{
		{
			testName: "",
			version:  "2",
			composers: []typdocker.Composer{
				redisV2,
			},
			expected: `version: "2"
services:
  redis: redis-service
  webdis: webdis-service
networks:
  webdis: webdis-network
volumes:
  redis: redis-volume
`,
		},
		{
			testName: "version not match",
			version:  "3",
			composers: []typdocker.Composer{
				redisV2,
			},
			expected: `version: "3"
services: {}
networks: {}
volumes: {}
`,
		},
		{
			testName: "multiple composer",
			version:  "2",
			composers: []typdocker.Composer{
				redisV2,
				pgV2,
			},
			expected: `version: "2"
services:
  pg: pg-service
  redis: redis-service
  webdis: webdis-service
networks:
  pg: pg-network
  webdis: webdis-network
volumes:
  pg: pg-volume
  redis: redis-volume
`,
		},
	}

	for _, tt := range testcases {

		t.Run(tt.testName, func(t *testing.T) {
			os.Remove("docker-compose.yml")
			defer os.Remove("docker-compose.yml")

			utility := &typdocker.Command{
				Version:   tt.version,
				Composers: tt.composers,
			}

			require.NoError(t, utility.Compose(&typgo.Context{}))

			b, _ := ioutil.ReadFile("docker-compose.yml")
			require.Equal(t, tt.expected, string(b))
		})
	}
}
