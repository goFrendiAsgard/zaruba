package hook

import (
	"io/ioutil"
	"path"
	"strings"
	"testing"
	"time"
	"os"

	"github.com/state-alchemists/zaruba/dir"
	"github.com/state-alchemists/zaruba/config"
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

	// inspect repos/classifiers/pre.txt
	expectedPreContent := "pre"
	preContent, err := readGeneratedFile(project, path.Join("repos", "classifiers", "pre.txt"))
	if err != nil {
		t.Error(err)
	} else if strings.Trim(expectedPreContent, "\n") != strings.Trim(preContent, "\n") {
		t.Errorf("Expected:\n%s\nActual:\n%s", expectedPreContent, preContent)
	}

	// inspect repos/classifiers/post.txt
	expectedPostContent := "post"
	postContent, err := readGeneratedFile(project, path.Join("repos", "classifiers", "post.txt"))
	if err != nil {
		t.Error(err)
	} else if strings.Trim(expectedPostContent, "\n") != strings.Trim(postContent, "\n") {
		t.Errorf("Expected:\n%s\nActual:\n%s", expectedPostContent, postContent)
	}

	// inspect services/sentiment/classifiers/classifier.py
	_, err = readGeneratedFile(project, path.Join("services", "sentiment", "classifiers", "classifier.py"))
	if err != nil {
		t.Error(err)
	}

	// inspect services/sentiment/post.txt
	_, err = readGeneratedFile(project, path.Join("services", "sentiment", "post.txt"))
	if err != nil {
		t.Error(err)
	}

	// inspect services/image/classifiers/classifier.py
	_, err = readGeneratedFile(project, path.Join("services", "image", "classifiers", "classifier.py"))
	if err != nil {
		t.Error(err)
	}

	// inspect services/image/post.txt
	_, err = readGeneratedFile(project, path.Join("services", "image", "post.txt"))
	if err != nil {
		t.Error(err)
	}
}

func readGeneratedFile(project, filepath string) (string, error) {
	data, err := ioutil.ReadFile(path.Join(project, filepath))
	return string(data), err
}
