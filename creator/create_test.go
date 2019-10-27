package creator

import (
	"io/ioutil"
	"os"
	"path"
	"testing"
)

func TestCreate(t *testing.T) {
	os.Setenv("ZARUBA_TEMPLATE_DIR", "playground/templates")
	os.Setenv("sender", "sender@gmail.com")
	os.Setenv("receiver", "receiver@gmail.com")
	parentpath := "playground/projects/test-create"
	Create("test", parentpath)

	// inspect readme.txt
	expectedReadmeContent := "# Test\nThis is a {{ test }}"
	readmeContent, err := readGeneratedFile(parentpath, "readme.txt")
	if err != nil {
		t.Errorf("%#v", err)
	}
	if readmeContent != expectedReadmeContent {
		t.Errorf("Expected:\n%s\nActual:\n%s", expectedReadmeContent, readmeContent)
	}

	// inspect email.txt
	expectedEmailContent := "from: sender@gmail.com\nto: receiver@gmail.com\nHello,\nThis is an email from sender@gmail.com to receiver@gmail.com"
	emailContent, err := readGeneratedFile(parentpath, "email.txt")
	if err != nil {
		t.Errorf("%#v", err)
	}
	if emailContent != expectedEmailContent {
		t.Errorf("Expected:\n%s\nActual:\n%s", expectedEmailContent, emailContent)
	}

	// inspect hello.txt
	expectedHelloContent := "hello world"
	helloContent, err := readGeneratedFile(parentpath, "hello.txt")
	if err != nil {
		t.Errorf("%#v", err)
	}
	if helloContent != expectedHelloContent {
		t.Errorf("Expected:\n%s\nActual:\n%s", expectedHelloContent, helloContent)
	}
}

func readGeneratedFile(parentpath, filepath string) (string, error) {
	data, err := ioutil.ReadFile(path.Join(parentpath, filepath))
	return string(data), err
}
