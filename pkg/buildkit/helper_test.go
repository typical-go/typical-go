package buildkit_test

// NOTE: remove command because problem with github action
// func TestCommandExist(t *testing.T) {
// 	testcases := []struct {
// 		name     string
// 		expected bool
// 	}{
// 		{"go", true},
// 		{"", false},
// 		{"invalid-command", false},
// 	}

// 	ctx := context.Background()
// 	for _, tt := range testcases {
// 		require.Equal(t, tt.expected, buildkit.AvailableCommand(ctx, tt.name),
// 			"%s expected %t", tt.name, tt.expected)
// 	}
// }
