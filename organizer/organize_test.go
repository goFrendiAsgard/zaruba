package organizer

import (
	"io/ioutil"
	"path"
	"strings"
	"testing"
	"time"
)

func TestOrganize(t *testing.T) {
	project := path.Join("..", "playground", "projects", "test-organize")
	// start watcher and wait for a while
	Organize(project)
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
