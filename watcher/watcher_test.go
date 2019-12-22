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
	if err := file.Copy("../test-resource/testOrganize.template", testPath); err != nil {
		t.Errorf("[ERROR] Cannot copy test-case: %s", err)
		return
	}

	// Watch project should succeed
	stopChan := make(chan bool)
	errChan := make(chan error)
	go Watch(testPath, errChan, stopChan)
	time.Sleep(time.Millisecond * 500)
	triggerFilePath := filepath.Join(testPath, "lib/b/watchTrigger.txt")
	os.Create(triggerFilePath)
	time.Sleep(time.Millisecond * 300)
	stopChan <- true
	err := <-errChan
	if err != nil {
		t.Errorf("[ERROR] Cannot watch: %s", err)
	}

	triggerFileInfo, err := os.Stat(triggerFilePath)

	// a.txt on D's content should be a
	aOnDFilePath := filepath.Join(testPath, "service/d/controller/c/lib/a/a.txt")
	aOnDByteContent, err := ioutil.ReadFile(aOnDFilePath)
	if err != nil {
		t.Errorf("[ERROR] Cannot read %s: %s", aOnDFilePath, err)
	} else {
		aOnDContent := strings.Trim(string(aOnDByteContent), "\n")
		if aOnDContent != "a" {
			t.Errorf("[UNEXPECTED] content should be `a`: %s", aOnDContent)
		}
		// a.txt on D's should be older than watchTrigger.txt
		aOnDFileInfo, err := os.Stat(aOnDFilePath)
		if err != nil {
			t.Errorf("[ERROR] Cannot get filestat of %s: %s", aOnDFilePath, err)
		} else if aOnDFileInfo.ModTime().Before(triggerFileInfo.ModTime()) {
			t.Errorf(
				"[UNEXPECTED] %s (%s) should be older than %s (%s)",
				aOnDFilePath, aOnDFileInfo.ModTime(), triggerFilePath, triggerFileInfo.ModTime(),
			)
		}
	}

	// a.txt on C's content should be a
	aOnCFilePath := filepath.Join(testPath, "controller/c/lib/a/a.txt")
	aByteContent, err := ioutil.ReadFile(aOnCFilePath)
	if err != nil {
		t.Errorf("[ERROR] Cannot read %s: %s", aOnCFilePath, err)
	} else {
		aContent := strings.Trim(string(aByteContent), "\n")
		if aContent != "a" {
			t.Errorf("[UNEXPECTED] content should be `a`: %s", aContent)
		}
		// a.txt on C's should be older than watchTrigger.txt
		aOnCFileInfo, err := os.Stat(aOnCFilePath)
		if err != nil {
			t.Errorf("[ERROR] Cannot get filestat of %s: %s", aOnCFilePath, err)
		} else if aOnCFileInfo.ModTime().After(triggerFileInfo.ModTime()) {
			t.Errorf(
				"[UNEXPECTED] %s (%s) should be older than %s (%s)",
				aOnCFilePath, aOnCFileInfo.ModTime(), triggerFilePath, triggerFileInfo.ModTime(),
			)
		}
	}

	// b.txt on D's content should be b
	bFilePath := filepath.Join(testPath, "service/d/controller/c/lib/b/b.txt")
	bByteContent, err := ioutil.ReadFile(bFilePath)
	if err != nil {
		t.Errorf("[ERROR] Cannot read %s: %s", bFilePath, err)
	} else {
		bContent := strings.Trim(string(bByteContent), "\n")
		if bContent != "b" {
			t.Errorf("[UNEXPECTED] content should be `b`: %s", bContent)
		}
		// b.txt should be older than watchTrigger.txt
		bFileInfo, err := os.Stat(bFilePath)
		if err != nil {
			t.Errorf("[ERROR] Cannot get filestat of %s: %s", bFilePath, err)
		} else if bFileInfo.ModTime().Before(triggerFileInfo.ModTime()) {
			t.Errorf(
				"[UNEXPECTED] %s (%s) should be older than %s (%s)",
				triggerFilePath, triggerFileInfo.ModTime(), bFilePath, bFileInfo.ModTime(),
			)
		}
	}

	// c.txt on D's content should be c
	cFilePath := filepath.Join(testPath, "service/d/controller/c/c.txt")
	cByteContent, err := ioutil.ReadFile(cFilePath)
	if err != nil {
		t.Errorf("[ERROR] Cannot read %s: %s", cFilePath, err)
	} else {
		cContent := strings.Trim(string(cByteContent), "\n")
		if cContent != "c" {
			t.Errorf("[UNEXPECTED] content should be `c`: %s", cContent)
		}
		// c.txt should be older than watchTrigger.txt
		cFileInfo, err := os.Stat(cFilePath)
		if err != nil {
			t.Errorf("[ERROR] Cannot get filestat of %s: %s", cFilePath, err)
		} else if cFileInfo.ModTime().Before(triggerFileInfo.ModTime()) {
			t.Errorf(
				"[UNEXPECTED] %s (%s) should be older than %s (%s)",
				triggerFilePath, triggerFileInfo.ModTime(), cFilePath, cFileInfo.ModTime(),
			)
		}

	}

}
