package hook

import (
	"io/ioutil"
	"os"
	"path"
	"strings"
	"testing"
	"time"

	"github.com/state-alchemists/zaruba/config"
	"github.com/state-alchemists/zaruba/dir"
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

func TestHookRunAction(t *testing.T) {
	project := path.Join("..", "playground", "projects", "test-hook")
	// start watcher and wait for a while

	allDirPaths, err := dir.GetAllDirPaths(project)
	if err != nil {
		t.Error(err)
	}
	// create hookConfig
	hc, err := NewCascadedConfig(allDirPaths)
	if err != nil {
		t.Error(err)
	}
	sortedKeys := hc.GetSortedKeys()
	shell := config.GetShell()
	environ := os.Environ()
	for _, key := range sortedKeys {
		if err = hc.RunAction(shell, environ, key); err != nil {
			t.Error(err)
		}
	}
	time.Sleep(1 * time.Second)

	// repos/classifiers
	// - repos/classifiers/readme.txt
	expectedReadmeContent := "v 0.0.0"
	readmeContent, err := readFile(path.Join(project, "repos", "classifiers", "readme.txt"))
	if err != nil {
		t.Error(err)
	} else if strings.Trim(expectedReadmeContent, "\n") != strings.Trim(readmeContent, "\n") {
		t.Errorf("Expected:\n%s\nActual:\n%s", expectedReadmeContent, readmeContent)
	}
	// - repos/classifiers/changelog.txt
	expectedChangelogContent := "v 0.0.0"
	changelogContent, err := readFile(path.Join(project, "repos", "classifiers", "changelog.txt"))
	if err != nil {
		t.Error(err)
	} else if strings.Trim(expectedChangelogContent, "\n") != strings.Trim(changelogContent, "\n") {
		t.Errorf("Expected:\n%s\nActual:\n%s", expectedChangelogContent, changelogContent)
	}

	// services/classifiers/sentiment
	// - services/sentiment/classifiers/readme.txt
	expectedReadmeContent = "v 0.0.0"
	readmeContent, err = readFile(path.Join(project, "services", "sentiment", "classifiers", "readme.txt"))
	if err != nil {
		t.Error(err)
	} else if strings.Trim(expectedReadmeContent, "\n") != strings.Trim(readmeContent, "\n") {
		t.Errorf("Expected:\n%s\nActual:\n%s", expectedReadmeContent, readmeContent)
	}
	// - services/sentiment/classifiers/changelog.txt
	expectedChangelogContent = "v 0.0.0"
	changelogContent, err = readFile(path.Join(project, "services", "sentiment", "classifiers", "changelog.txt"))
	if err != nil {
		t.Error(err)
	} else if strings.Trim(expectedChangelogContent, "\n") != strings.Trim(changelogContent, "\n") {
		t.Errorf("Expected:\n%s\nActual:\n%s", expectedChangelogContent, changelogContent)
	}
	// - services/sentiment/classifiers/output.txt
	_, err = readFile(path.Join(project, "services", "sentiment", "output.txt"))
	if err != nil {
		t.Error(err)
	}

	// services/classifiers/image
	// - services/image/classifiers/readme.txt
	expectedReadmeContent = "v 0.0.0"
	readmeContent, err = readFile(path.Join(project, "services", "image", "classifiers", "readme.txt"))
	if err != nil {
		t.Error(err)
	} else if strings.Trim(expectedReadmeContent, "\n") != strings.Trim(readmeContent, "\n") {
		t.Errorf("Expected:\n%s\nActual:\n%s", expectedReadmeContent, readmeContent)
	}
	// - services/image/classifiers/changelog.txt
	expectedChangelogContent = "v 0.0.0"
	changelogContent, err = readFile(path.Join(project, "services", "image", "classifiers", "changelog.txt"))
	if err != nil {
		t.Error(err)
	} else if strings.Trim(expectedChangelogContent, "\n") != strings.Trim(changelogContent, "\n") {
		t.Errorf("Expected:\n%s\nActual:\n%s", expectedChangelogContent, changelogContent)
	}
	// - services/image/classifiers/output.txt
	_, err = readFile(path.Join(project, "services", "image", "output.txt"))
	if err != nil {
		t.Error(err)
	}

}

func readFile(filepath string) (string, error) {
	data, err := ioutil.ReadFile(filepath)
	return string(data), err
}
