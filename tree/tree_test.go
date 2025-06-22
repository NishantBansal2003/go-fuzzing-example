package tree

import "testing"

// FuzzBuildTree fuzzes the BuildTree function.
func FuzzBuildTree(f *testing.F) {
	f.Fuzz(func(t *testing.T, depth int) {
		_ = BuildTree(depth)
	})
}
