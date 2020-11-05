package config

import (
	"testing"
)

func TestValidTaskParsedEnv(t *testing.T) {
	if err := setupValidProjectConfig(t); err != nil {
		return
	}
	expectation := map[string]map[string]string{
		"runApiGateway": {
			"HTTP_PORT":       "8080",
			"PROMETHEUS_PORT": "8081",
			"SOME_KEY":        "SOME_VALUE",
		},
		"runIntegrationTest": {},
		"core.runNodeJsService": {
			"SOME_KEY": "SOME_VALUE",
		},
	}
	for taskName, parsedEnvExpectation := range expectation {
		task, exists := validConf.Tasks[taskName]
		if !exists {
			t.Errorf("Task %s is not exist", taskName)
			continue
		}
		for key := range parsedEnvExpectation {
			if _, exists := task.ParsedEnv[key]; !exists {
				t.Errorf("Expecting env %s of %s but not found", key, taskName)
			}
		}
		for key, actualVal := range task.ParsedEnv {
			expectedVal, exists := parsedEnvExpectation[key]
			if !exists {
				t.Errorf("Env key %s of %s is not expected", key, taskName)
			}
			if expectedVal != actualVal {
				t.Errorf("Expecting env %s of %s to be %s but getting %s", key, taskName, expectedVal, actualVal)
			}
		}
	}
}

func TestValidTaskGetEnv(t *testing.T) {
	if err := setupValidProjectConfig(t); err != nil {
		return
	}
	expectation := map[string]map[string]string{
		"runApiGateway": {
			"HTTP_PORT":                   "8080",
			"PROMETHEUS_PORT":             "8081",
			"SOME_KEY":                    "SOME_VALUE",
			"API_GATEWAY_HTTP_PORT":       "8080",
			"API_GATEWAY_PROMETHEUS_PORT": "8081",
		},
		"runIntegrationTest": {
			"API_GATEWAY_HTTP_PORT":       "8080",
			"API_GATEWAY_PROMETHEUS_PORT": "8081",
		},
		"core.runNodeJsService": {
			"SOME_KEY":                    "SOME_VALUE",
			"API_GATEWAY_HTTP_PORT":       "8080",
			"API_GATEWAY_PROMETHEUS_PORT": "8081",
		},
	}
	for taskName, parsedEnvExpectation := range expectation {
		task, exists := validConf.Tasks[taskName]
		if !exists {
			t.Errorf("Task %s is not exist", taskName)
			continue
		}
		for key, expectedVal := range parsedEnvExpectation {
			if actualVal := task.GetEnv(key); actualVal != expectedVal {
				t.Errorf("Expecting env %s of %s to be %s but getting %s", key, taskName, expectedVal, actualVal)
			}
		}
	}
}

func TestValidTaskParsedConfig(t *testing.T) {
	if err := setupValidProjectConfig(t); err != nil {
		return
	}
	expectation := map[string]map[string]string{
		"runApiGateway": {
			"checkPort": "8080",
			"checkHost": "localhost",
		},
		"runIntegrationTest": {},
		"core.runNodeJsService": {
			"checkPort": "3000",
			"checkHost": "localhost",
		},
	}
	for taskName, parsedConfigExpectation := range expectation {
		task, exists := validConf.Tasks[taskName]
		if !exists {
			t.Errorf("Task %s is not exist", taskName)
			continue
		}
		for key := range parsedConfigExpectation {
			if _, exists := task.ParsedConfig[key]; !exists {
				t.Errorf("Expecting config %s of %s but not found", key, taskName)
			}
		}
		for key, actualVal := range task.ParsedConfig {
			expectedVal, exists := parsedConfigExpectation[key]
			if !exists {
				t.Errorf("Config %s of %s is not expected", key, taskName)
			}
			if expectedVal != actualVal {
				t.Errorf("Expecting config %s of %s to be %s but getting %s", key, taskName, expectedVal, actualVal)
			}
		}
	}
}

func TestValidTaskParsedLConfig(t *testing.T) {
	if err := setupValidProjectConfig(t); err != nil {
		return
	}
	expectation := map[string]map[string][]string{
		"runApiGateway": {
			"tags":        {"api", "nodejs"},
			"requirement": {"nodejs-v11", "npm-v6.5.0"},
		},
		"runIntegrationTest": {},
		"core.runNodeJsService": {
			"tags":        {"nodejs", "service"},
			"requirement": {"nodejs-v11", "npm-v6.5.0"},
		},
	}
	for taskName, parsedLConfigExpectation := range expectation {
		task, exists := validConf.Tasks[taskName]
		if !exists {
			t.Errorf("Task %s is not exist", taskName)
			continue
		}
		for key := range parsedLConfigExpectation {
			if _, exists := task.ParsedLConfig[key]; !exists {
				t.Errorf("Expecting config %s of %s but not found", key, taskName)
			}
		}
		for key, actualVals := range task.ParsedLConfig {
			expectedVals, exists := parsedLConfigExpectation[key]
			if !exists {
				t.Errorf("Config %s of %s is not expected", key, taskName)
				continue
			}
			if len(expectedVals) != len(actualVals) {
				t.Errorf("Expecting config %s of %s to contains %d element but getting %d element", key, taskName, len(expectedVals), len(actualVals))
				continue
			}
			for index, expectedVal := range expectedVals {
				actualVal := actualVals[index]
				if actualVal != expectedVal {
					t.Errorf("Expecting config %s[%d] of %s to be %s but getting %s", key, index, taskName, expectedVal, actualVal)
				}
			}
		}
	}
}

func TestValidTaskStartCommand(t *testing.T) {
	if err := setupValidProjectConfig(t); err != nil {
		return
	}
	expectation := map[string][]string{
		"runIntegrationTest":       {"npm", "start"},
		"serveStaticFiles":         {"python", "-m", "http.server", "9000"},
		"core.runStaticWebService": {"python", "-m", "http.server", "8080"},
	}
	for taskName, expectedArgs := range expectation {
		task, exists := validConf.Tasks[taskName]
		if !exists {
			t.Errorf("Task %s is not exist", taskName)
			continue
		}
		startCmd, exist, err := task.GetStartCmd()
		if err != nil {
			t.Error(err)
			continue
		}
		if !exist {
			t.Errorf("Not exist")
		}
		actualArgs := startCmd.Args
		if len(actualArgs) != len(expectedArgs) {
			t.Errorf("Expecting args of %s to contains %d element but getting %d element", taskName, len(expectedArgs), len(actualArgs))
			continue
		}
		for index, expectedArg := range expectedArgs {
			actualArg := actualArgs[index]
			if actualArg != expectedArg {
				t.Errorf("Expecting Args[%d] of %s to be %s but getting %s", index, taskName, expectedArg, actualArg)
			}
		}
	}
}

func TestValidTaskCheckCommand(t *testing.T) {
	if err := setupValidProjectConfig(t); err != nil {
		return
	}
	expectation := map[string][]string{
		"runApiGateway":         {"sh", "-c", "until nc -z localhost 8080; do sleep 1; done"},
		"core.runNodeJsService": {"sh", "-c", "until nc -z localhost 3000; do sleep 1; done"},
	}
	for taskName, expectedArgs := range expectation {
		task, exists := validConf.Tasks[taskName]
		if !exists {
			t.Errorf("Task %s is not exist", taskName)
			continue
		}
		checkCmd, exist, err := task.GetCheckCmd()
		if err != nil {
			t.Error(err)
			continue
		}
		if !exist {
			t.Errorf("Not exist")
		}
		actualArgs := checkCmd.Args
		if len(actualArgs) != len(expectedArgs) {
			t.Errorf("Expecting args of %s to contains %d element but getting %d element", taskName, len(expectedArgs), len(actualArgs))
			continue
		}
		for index, expectedArg := range expectedArgs {
			actualArg := actualArgs[index]
			if actualArg != expectedArg {
				t.Errorf("Expecting Args[%d] of %s to be %s but getting %s", index, taskName, expectedArg, actualArg)
			}
		}
	}
}

func TestInvalidTaskExtension(t *testing.T) {
	conf, err := NewConfig("../test_resource/invalidTaskExtension.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	if err = conf.Init(); err == nil {
		t.Errorf("Error expected")
	}
}

func TestInvalidTaskConfigBrokenTemplate(t *testing.T) {
	conf, err := NewConfig("../test_resource/invalidTaskConfigBrokenTemplate.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	if err = conf.Init(); err == nil {
		t.Errorf("Error expected")
	}
}

func TestInvalidTaskConfigNonExecutableTemplate(t *testing.T) {
	conf, err := NewConfig("../test_resource/invalidTaskConfigNonExecutableTemplate.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	if err = conf.Init(); err == nil {
		t.Errorf("Error expected")
	}
}

func TestInvalidTaskLConfigBrokenTemplate(t *testing.T) {
	conf, err := NewConfig("../test_resource/invalidTaskLConfigBrokenTemplate.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	if err = conf.Init(); err == nil {
		t.Errorf("Error expected")
	}
}

func TestInvalidTaskLConfigNonExecutableTemplate(t *testing.T) {
	conf, err := NewConfig("../test_resource/invalidTaskLConfigNonExecutableTemplate.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	if err = conf.Init(); err == nil {
		t.Errorf("Error expected")
	}
}

func TestInvalidTaskUndefinedCommandWithoutParent(t *testing.T) {
	conf, err := NewConfig("../test_resource/invalidTaskUndefinedCommandWithoutParent.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	if err = conf.Init(); err != nil {
		t.Error(err)
	}
	task, exists := conf.Tasks["doSomething"]
	if !exists {
		t.Errorf("Task not found")
		return
	}
	if _, _, err = task.GetStartCmd(); err == nil {
		t.Errorf("Error expected")
	}
	if _, _, err = task.GetCheckCmd(); err == nil {
		t.Errorf("Error expected")
	}
}

func TestInvalidTaskBrokenCommand(t *testing.T) {
	conf, err := NewConfig("../test_resource/invalidTaskBrokenCommand.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	if err = conf.Init(); err != nil {
		t.Error(err)
	}
	task, exists := conf.Tasks["doSomething"]
	if !exists {
		t.Errorf("Task not found")
		return
	}
	if _, _, err = task.GetStartCmd(); err == nil {
		t.Errorf("Error expected")
	}
	if _, _, err = task.GetCheckCmd(); err == nil {
		t.Errorf("Error expected")
	}
}
