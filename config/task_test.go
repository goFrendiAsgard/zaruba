package config

import (
	"testing"
)

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
			actualVal, err := task.GetEnv(NewTaskData(task), key)
			if err != nil {
				t.Error(err)
			}
			if actualVal != expectedVal {
				t.Errorf("Expecting env %s of %s to be %s but getting %s", key, taskName, expectedVal, actualVal)
			}
		}
	}
}

func TestValidTaskGetConfig(t *testing.T) {
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
			expectedVal := parsedConfigExpectation[key]
			actualVal, err := task.GetConfig(NewTaskData(task), key)
			if err != nil {
				t.Error(err)
			}
			if actualVal != expectedVal {
				t.Errorf("Expecting config %s of %s to be %s but getting %s", key, taskName, expectedVal, actualVal)
			}
		}
	}
}

func TestValidTaskGetLConfig(t *testing.T) {
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
		for key, expectedVals := range parsedLConfigExpectation {
			actualVals, err := task.GetLConfig(NewTaskData(task), key)
			if err != nil {
				t.Error(err)
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
		startCmd, exist, err := task.GetStartCmd(make(chan error))
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
		checkCmd, exist, err := task.GetCheckCmd(make(chan error))
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

func TestValidTaskDependencies(t *testing.T) {
	conf, err := NewConfig("../test_resource/dependencies.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	task := conf.Tasks["buyMilk"]
	expectedDependencies := []string{"driveCar", "haveMoney"}
	actualDependencies := task.GetDependencies()
	if len(actualDependencies) != len(expectedDependencies) {
		t.Errorf("Expecting dependencies of %s to contains %d element but getting %d element", task.Name, len(expectedDependencies), len(actualDependencies))
		return
	}
	for index, expectedDependency := range expectedDependencies {
		actualDependency := actualDependencies[index]
		if actualDependency != expectedDependency {
			t.Errorf("Expecting Dependencies[%d] of %s to be %s but getting %s", index, task.Name, expectedDependency, actualDependency)
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
	if err = conf.Init(); err != nil {
		t.Error(err)
		return
	}
	task := conf.Tasks["job"]
	if _, parseError := task.GetConfig(NewTaskData(task), "someKey"); parseError == nil {
		t.Errorf("Error expected")
	}
}

func TestInvalidTaskConfigNonExecutableTemplate(t *testing.T) {
	conf, err := NewConfig("../test_resource/invalidTaskConfigNonExecutableTemplate.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	if err = conf.Init(); err != nil {
		t.Error(err)
		return
	}
	task := conf.Tasks["job"]
	if _, parseError := task.GetConfig(NewTaskData(task), "someKey"); parseError == nil {
		t.Errorf("Error expected")
	}
}

func TestInvalidTaskLConfigBrokenTemplate(t *testing.T) {
	conf, err := NewConfig("../test_resource/invalidTaskLConfigBrokenTemplate.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	if err = conf.Init(); err != nil {
		t.Error(err)
		return
	}
	task := conf.Tasks["job"]
	if _, parseError := task.GetLConfig(NewTaskData(task), "someKey"); parseError == nil {
		t.Errorf("Error expected")
	}
}

func TestInvalidTaskLConfigNonExecutableTemplate(t *testing.T) {
	conf, err := NewConfig("../test_resource/invalidTaskLConfigNonExecutableTemplate.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	if err = conf.Init(); err != nil {
		t.Error(err)
		return
	}
	task := conf.Tasks["job"]
	if _, parseError := task.GetLConfig(NewTaskData(task), "someKey"); parseError == nil {
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
	if _, _, err = task.GetStartCmd(make(chan error)); err == nil {
		t.Errorf("Error expected")
	}
	if _, _, err = task.GetCheckCmd(make(chan error)); err == nil {
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
	if _, _, err = task.GetStartCmd(make(chan error)); err == nil {
		t.Errorf("Error expected")
	}
	if _, _, err = task.GetCheckCmd(make(chan error)); err == nil {
		t.Errorf("Error expected")
	}
}
