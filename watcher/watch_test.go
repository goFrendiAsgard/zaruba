package watcher

import (
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
	"testing"
	"time"
)

func TestGetSortedKeys(t *testing.T) {
	hc := HookConfig{
		"/a/b/a": SingleHookConfig{
			Links: []string{"/a/b/c/a"},
		},
		"/a/b/d": SingleHookConfig{
			Links: []string{"/a/b/e"},
		},
		"/a/b/e": SingleHookConfig{},
		"/a/b/c": SingleHookConfig{
			Links: []string{"/a/b/d"},
		},
		"/a/b/b": SingleHookConfig{
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

func TestWatch(t *testing.T) {
	project := path.Join("..", "playground", "projects", "test-watch")
	stop := make(chan bool)
	// start watcher and wait for a while
	go Watch(project, stop)
	time.Sleep(1 * time.Second)
	// trigger changes
	log.Println("Trigger changes")
	os.MkdirAll(path.Join(project, "repos", "classifiers", "trigger"), os.ModePerm)
	time.Sleep(2 * time.Second)

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

	stop <- true
}

func readGeneratedFile(project, filepath string) (string, error) {
	data, err := ioutil.ReadFile(path.Join(project, filepath))
	return string(data), err
}
