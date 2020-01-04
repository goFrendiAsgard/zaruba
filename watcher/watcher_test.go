package watcher

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/state-alchemists/zaruba/config"
	"github.com/state-alchemists/zaruba/file"
)

func TestWatch(t *testing.T) {
	baseTestPath := config.GetTestDir()
	testPath := filepath.Join(baseTestPath, "testWatch")

	if err := file.Copy("../test-resource/project", testPath); err != nil {
		t.Errorf("[ERROR] Cannot copy test-case: %s", err)
		return
	}
	if err := file.Copy("../test-resource/zaruba.config.megazord.yaml", filepath.Join(testPath, "zaruba.config.yaml")); err != nil {
		t.Errorf("[ERROR] Cannot copy zaruba.config.yaml: %s", err)
		return
	}

	triggerFilePath := filepath.Join(testPath, "services/greeter/trigger.txt")
	megazordGatewayFilePath := filepath.Join(testPath, "megazord/gateway")
	megazordGreeterFilePath := filepath.Join(testPath, "megazord/greeter")

	// Watch project should succeed
	stopChan := make(chan bool)
	errChan := make(chan error)
	go Watch(testPath, errChan, stopChan)
	time.Sleep(time.Millisecond * 500)
	os.Create(triggerFilePath)
	time.Sleep(time.Millisecond * 300)
	stopChan <- true
	err := <-errChan
	if err != nil {
		t.Errorf("[ERROR] Cannot watch: %s", err)
	}

	triggerFileInfo, err := os.Stat(triggerFilePath)
	if err != nil {
		t.Errorf("[ERROR] Cannot get stat of %s: %s", triggerFilePath, err)
		return
	}

	megazordGatewayFileInfo, err := os.Stat(megazordGatewayFilePath)
	if err != nil {
		t.Errorf("[ERROR] Cannot get stat of %s: %s", megazordGatewayFilePath, err)
		return
	}

	megazordGreeterFileInfo, err := os.Stat(megazordGreeterFilePath)
	if err != nil {
		t.Errorf("[ERROR] Cannot get stat of %s: %s", megazordGreeterFilePath, err)
		return
	}

	// megazord's gateway should be older than trigger.
	if !(megazordGatewayFileInfo.ModTime().Before(triggerFileInfo.ModTime())) {
		t.Errorf(
			"[UNEXPECTED] %s (%s) should be older than %s (%s)",
			megazordGatewayFilePath, megazordGatewayFileInfo.ModTime(),
			triggerFilePath, triggerFileInfo.ModTime(),
		)
	}

	// Trigger should be older than megazord's greeter.
	if !(megazordGreeterFileInfo.ModTime().After(triggerFileInfo.ModTime())) {
		t.Errorf(
			"[UNEXPECTED] %s (%s) should be older than %s (%s)",
			triggerFilePath, triggerFileInfo.ModTime(),
			megazordGreeterFilePath, megazordGreeterFileInfo.ModTime(),
		)
	}

	// log.txt should contains `pre-organize`, `organize`, and `post-organize`
	testFilePath := filepath.Join(testPath, "log.txt")
	testByteContent, err := ioutil.ReadFile(testFilePath)
	if err != nil {
		t.Errorf("[ERROR] Cannot read %s: %s", testFilePath, err)
	} else {
		testContent := strings.Trim(string(testByteContent), "\n")
		expected := "pre-organize\norganize\npost-organize\npre-organize\norganize\npost-organize"
		if testContent != expected {
			t.Errorf("[UNEXPECTED] content should be `%s`, get: `%s`", expected, testContent)
		}
	}

}
