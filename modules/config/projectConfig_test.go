package config

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/state-alchemists/zaruba/modules/file"
)

func TestLoadProjectConfig(t *testing.T) {
	baseTestPath := GetTestDir()
	testPath := filepath.Join(baseTestPath, "testProjectConfig")
	if err := file.Copy("../../test-resource/testProjectConfig", testPath); err != nil {
		t.Errorf("[ERROR] Cannot copy test-case: %s", err)
		return
	}

	// load project config
	p, err := CreateProjectConfig(testPath)
	if err != nil {
		t.Errorf("[ERROR] Cannot load config: %s", err)
		return
	}

	// get component by name
	rmq, err := p.GetComponentByName("rmq")
	if err != nil {
		t.Errorf("[ERROR] Cannot get component rmq: %s", err)
		return
	}
	gopher1, err := p.GetComponentByName("gopher1")
	if err != nil {
		t.Errorf("[ERROR] Cannot get component gopher1: %s", err)
		return
	}
	_, err = p.GetComponentByName("gopher2")
	if err != nil {
		t.Errorf("[ERROR] Cannot get component gopher2: %s", err)
		return
	}

	// get component by labels: scenario:test
	selectors := []string{"scenario:test"}
	components, err := p.GetComponentsBySelectors(selectors)
	testComponentsExist(t, selectors, components, []string{"gopher1-test", "gopher2-test"})

	// get component by labels: language:go
	selectors = []string{"language:go"}
	components, err = p.GetComponentsBySelectors(selectors)
	testComponentsExist(t, selectors, components, []string{"gopher1-test", "gopher2-test", "gopher1", "gopher2"})

	// get component by labels: scenario:integration-test
	selectors = []string{"scenario:integration-test"}
	components, err = p.GetComponentsBySelectors(selectors)
	testComponentsExist(t, selectors, components, []string{"gopher1-test"})

	// get component by labels: scenario:end-to-end-test
	selectors = []string{"scenario:end-to-end-test"}
	components, err = p.GetComponentsBySelectors(selectors)
	testComponentsExist(t, selectors, components, []string{"gopher1-test"})

	// get component by names or labels
	selectors = []string{"scenario:test", "rmq"}
	components, err = p.GetComponentsBySelectors(selectors)
	if err != nil {
		t.Errorf("[ERROR] Cannot get component language:go")
	}
	testComponentsExist(t, selectors, components, []string{"gopher1-test", "gopher2-test", "rmq"})

	// test each component
	testGopher1Component(t, testPath, gopher1)
	testRmqComponent(t, testPath, rmq)

	// test sortedLinkSource
	sortedLinkSources := p.GetSortedLinkSources()
	expectedSortedLinkSources := []string{"/tmp/zaruba-test/testProjectConfig/libraries/go/transport"}
	for index, source := range expectedSortedLinkSources {
		if source != sortedLinkSources[index] {
			t.Errorf("[UNEXPECTED] sortedLinkSource[%d] should be `%s`, get `%s`", index, source, sortedLinkSources[index])
		}
	}

	// test cascaded config
	expected := filepath.Join(testPath, "./services/gopher1")
	if err != nil {
		t.Errorf("[ERROR] Cannot get component: %s", err)
	} else if gopher1.GetLocation() != expected {
		t.Errorf("[UNEXPECTED] gopher1.Location should be `%s`, but contains `%s`", expected, gopher1.GetLocation())
	}

	// test subRepoPrefixMap
	expected = "services/gopher1"
	if subRepoPrefixMap := p.GetSubrepoPrefixMap(testPath); subRepoPrefixMap["gopher1"] != expected {
		t.Errorf("[UNEXPECTED] subRepoPrefixMap[\"gopher1\"] should be `%s`, but contains: `%s`", expected, subRepoPrefixMap["gopher1"])
	}

	// test YAML conversion
	yaml, err := p.ToYaml()
	if err != nil {
		t.Errorf("[ERROR] Cannot convert config into YAML: %s", err)
	} else if yaml == "" {
		t.Errorf("[UNEXPECTED] YAML is empty")
	}

}

func testGopher1Component(t *testing.T, testPath string, gopher1 *Component) {
	// check gopher1's properties
	if gopher1.GetName() != "gopher1" {
		t.Errorf("[UNEXPECTED] gopher1's name should be `gopher1`, but it contains `%s`", gopher1.GetName())
	}
	if gopher1.GetSymbol() != "" {
		t.Errorf("[UNEXPECTED] gopher1's symbol should be ``, but it contains `%s`", gopher1.GetSymbol())
	}
	if gopher1.GetRuntimeSymbol() == "" {
		t.Errorf("[UNEXPECTED] gopher1's runtimeSymbol should not be empty, but it contains `%s`", gopher1.GetRuntimeSymbol())
	}
	// expected runtimeLocation
	expectedRuntimeLocation := filepath.Join(testPath, "./services/gopher1")
	if gopher1.GetRuntimeLocation() != expectedRuntimeLocation {
		t.Errorf("[UNEXPECTED] gopher1's runtimeLocation should not be `%s`, but it contains `%s`", expectedRuntimeLocation, gopher1.GetRuntimeLocation())
	}
	if gopher1.GetColor() == 0 {
		t.Errorf("[UNEXPECTED] gopher1's color should be `0`, but it contains `%d`", gopher1.GetColor())
	}
	if gopher1.GetStartCommand() != "go build && ./app" {
		t.Errorf("[UNEXPECTED] gopher1's start command should be `go build && ./app`, but it contains `%s`", gopher1.GetStartCommand())
	}
	if gopher1.GetStartCommand() != "go build && ./app" {
		t.Errorf("[UNEXPECTED] gopher1's runtime start command should be `go build && ./app`, but it contains `%s`", gopher1.GetStartCommand())
	}
	if gopher1.GetRuntimeCommand() != gopher1.GetStartCommand() {
		t.Errorf("[UNEXPECTED] gopher1's runtime command should be `%s`, but it contains `%s`", gopher1.GetStartCommand(), gopher1.GetRuntimeCommand())
	}
	if len(gopher1.GetRuntimeName()) < 12 {
		t.Errorf("[UNEXPECTED] gopher1's runtimeReadinessURL should be at least 12 character, but it contains `%s`", gopher1.GetRuntimeName())
	}
	// test component runtime environment
	expectedEnv := map[string]string{
		"GOPHER1_HTTP_PORT": "3011",
		"GOPHER2_HTTP_PORT": "3012",
		"RMQ_PASS":          "toor",
		"RMQ_PORT":          "5772",
		"RMQ_USER":          "root",
		"gopher1":           "0.0.0.0",
		"gopher2":           "0.0.0.0",
		"rmq":               GetDockerHost(),
	}
	env := gopher1.GetRuntimeEnv()
	for name, expectedValue := range expectedEnv {
		if value := env[name]; value != expectedValue {
			t.Errorf("[UNEXPECTED] env `%s` should be `%s`, but contains `%s`", name, expectedValue, value)
		}
	}
}

func testRmqComponent(t *testing.T, testPath string, rmq *Component) {
	if rmq.GetStartCommand() != "" {
		t.Errorf("[UNEXPECTED] rmq's start command should be ``, but it contains `%s`", rmq.GetStartCommand())
	}
	if rmq.GetRuntimeCommand() == "" {
		t.Errorf("[UNEXPECTED] rmq's runtime run command should not be empty, but it contains `%s`", rmq.GetRuntimeCommand())
	}
	if rmq.GetRuntimeSymbol() != "🐇" {
		t.Errorf("[UNEXPECTED] rmq's runtimeSymbol should be `🐇`, but it contains `%s`", rmq.GetRuntimeSymbol())
	}
	// expected runtimeLocation
	expectedRuntimeLocation := testPath
	if rmq.GetRuntimeLocation() != expectedRuntimeLocation {
		t.Errorf("[UNEXPECTED] gopher1's runtimeLocation should not be `%s`, but it contains `%s`", expectedRuntimeLocation, rmq.GetRuntimeLocation())
	}
	if rmq.GetReadinessURL() != "http://${rmq}:${RMQ_API_PORT}" {
		t.Errorf("[UNEXPECTED] rmq's readinessURL should be `http://${rmq}:${RMQ_API_PORT}`, but it contains `%s`", rmq.GetReadinessURL())
	}
	expectedReadinessURL := fmt.Sprintf("http://%s:15772", GetDockerHost())
	if rmq.GetRuntimeReadinessURL() != expectedReadinessURL {
		t.Errorf("[UNEXPECTED] rmq's runtimeReadinessURL should be `%s`, but it contains `%s`", expectedReadinessURL, rmq.GetRuntimeReadinessURL())
	}
	if len(rmq.GetRuntimeName()) < 12 {
		t.Errorf("[UNEXPECTED] rmq's runtimeReadinessURL should be at least 12 character, but it contains `%s`", rmq.GetRuntimeName())
	}
	if rmq.GetRuntimeReadinessCheckCommand() != "if [ $(docker inspect -f '{{.State.Running}}' \"testProjectConfig-rmq\") = true ]; then echo ok; else echo notOk 1>&2; fi;" {
		t.Errorf("[UNEXPECTED] rmq's runtimeCheckCommand should be `if [ $(docker inspect -f '{{.State.Running}}' \"testProjectConfig-rmq\") = true ]; then echo ok; else echo notOk 1>&2; fi;`, but it contains `%s`", rmq.GetRuntimeReadinessCheckCommand())
	}
}

func testComponentsExist(t *testing.T, selectors []string, components map[string]*Component, names []string) {
	componentCount := 0
	for range components {
		componentCount++
	}
	if componentCount != len(names) {
		t.Errorf("[UNEXPECTED] expected component keys are %#v, but fetched components are %#v", names, components)
	}
	for _, name := range names {
		nameExists := false
		for componentName := range components {
			if componentName == name {
				nameExists = true
			}
		}
		if !nameExists {
			t.Errorf("[UNEXPECTED] component %s is not exists on %#v when fetched with selector %#v", name, components, selectors)
		}
	}
}
