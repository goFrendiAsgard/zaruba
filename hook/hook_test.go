package hook

import(
	"strings"
	"testing"
)

func TestGetSortedKeys(t *testing.T) {
	hc := Config{
		"/a/b/a": SingleConfig{
			Links: []string{"/a/b/c/a", "/a/b/b/a"},
		},
		"/a/b/d": SingleConfig{
			Links: []string{"/a/b/e"},
		},
		"/a/b/e": SingleConfig{},
		"/a/b/c": SingleConfig{
			Links: []string{"/a/b/d"},
		},
		"/a/b/b": SingleConfig{
			Links: []string{"/a/b/c/b"},
		},
	}
	sortedKeys := hc.GetSortedKeys()
	mergedSortedKeys := strings.Join(sortedKeys, ", ")
	expectedMergedSortedKeys := "/a/b/a, /a/b/b, /a/b/c, /a/b/d, /a/b/e"
	if mergedSortedKeys != expectedMergedSortedKeys {
		t.Errorf("Expected:\n%s\nActual:\n%s", expectedMergedSortedKeys, mergedSortedKeys)
	}
}
