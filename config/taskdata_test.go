package config

import (
	"path/filepath"
	"testing"
)

func setupTaskData(t *testing.T) (td *TaskData, err error) {
	project, err := NewProject("../test_resource/valid/zaruba.yaml")
	if err != nil {
		t.Error(err)
		return td, err
	}
	values := []string{
		"alchemist::flamel::name=Nicholas Flamel",
		"alchemist::flamel::age=90",
		"alchemist::dumbledore::name=Dumbledore",
		"alchemist::dumbledore::age=90",
		"barbarian::sonya::name=Sonya",
		"alchemist::=unknown",
	}
	for _, value := range values {
		if err = project.AddValue(value); err != nil {
			return td, err
		}
	}
	td = NewTaskData(project.Tasks["runApiGateway"])
	return td, err
}

func TestTaskDataGetWorkPath(t *testing.T) {
	td, err := setupTaskData(t)
	if err != nil {
		return
	}
	workPath := td.WorkPath
	absWorkPath := td.GetWorkPath(workPath)
	if workPath != absWorkPath {
		t.Errorf("Expecting %s, get %s", workPath, absWorkPath)
	}
	subPath := "./boom"
	absSubPath := td.GetWorkPath(subPath)
	expected := filepath.Join(workPath, subPath)
	if absSubPath != expected {
		t.Errorf("Expecting %s, get %s", expected, absSubPath)
	}
}

func TestTaskDataGetBasePath(t *testing.T) {
	td, err := setupTaskData(t)
	if err != nil {
		return
	}
	basePath := td.BasePath
	absBasePath := td.GetBasePath(basePath)
	if basePath != absBasePath {
		t.Errorf("Expecting %s, get %s", basePath, absBasePath)
	}
	subPath := "./boom"
	absSubPath := td.GetBasePath(subPath)
	expected := filepath.Join(basePath, subPath)
	if absSubPath != expected {
		t.Errorf("Expecting %s, get %s", expected, absSubPath)
	}
}

func TestTaskDataGetRelativePath(t *testing.T) {
	td, err := setupTaskData(t)
	if err != nil {
		return
	}
	basePath := td.BasePath
	absBasePath := td.GetRelativePath(basePath)
	if basePath != absBasePath {
		t.Errorf("Expecting %s, get %s", basePath, absBasePath)
	}
	subPath := "./boom"
	absSubPath := td.GetRelativePath(subPath)
	expected := filepath.Join(basePath, subPath)
	if absSubPath != expected {
		t.Errorf("Expecting %s, get %s", expected, absSubPath)
	}
}

func TestTaskDataGetConfig(t *testing.T) {
	td, err := setupTaskData(t)
	if err != nil {
		return
	}
	if _, err := td.GetConfig("checkPort"); err != nil {
		t.Error("config checkPort does not exist")
	}
}

func TestTaskDataGetConfigs(t *testing.T) {
	td, err := setupTaskData(t)
	if err != nil {
		return
	}
	if _, err := td.GetConfigs(); err != nil {
		t.Error(err)
	}
}

func TestTaskDataGetLConfig(t *testing.T) {
	td, err := setupTaskData(t)
	if err != nil {
		return
	}
	if _, exist := td.GetLConfig("tags"); exist != nil {
		t.Error("lconfig tags does not exist")
	}
}

func TestTaskDataGetLConfigs(t *testing.T) {
	td, err := setupTaskData(t)
	if err != nil {
		return
	}
	if _, err := td.GetLConfigs(); err != nil {
		t.Error(err)
	}
}

func TestTaskDataGetValue(t *testing.T) {
	td, err := setupTaskData(t)
	if err != nil {
		return
	}
	if _, exist := td.GetValue("alchemist::flamel::age"); exist != nil {
		t.Error("value alchemist::flamel::age does not exist")
	}
}

func TestTaskDataGetValues(t *testing.T) {
	td, err := setupTaskData(t)
	if err != nil {
		return
	}
	if _, err := td.GetValues(); err != nil {
		t.Error(err)
	}
}

func TestTaskDataGetEnv(t *testing.T) {
	td, err := setupTaskData(t)
	if err != nil {
		return
	}
	if _, exist := td.GetEnv("HTTP_PORT"); exist != nil {
		t.Error("env HTTP_PORT does not exist")
	}
}

func TestTaskDataGetEnvs(t *testing.T) {
	td, err := setupTaskData(t)
	if err != nil {
		return
	}
	if _, err := td.GetEnvs(); err != nil {
		t.Error(err)
	}
}

func TestTaskDataValuesGetSubKeys(t *testing.T) {
	td, err := setupTaskData(t)
	if err != nil {
		return
	}
	subkeys := td.GetSubValueKeys("alchemist")
	if len(subkeys) != 2 {
		t.Errorf("Subkeys length should be 2, but currently contains %#v", subkeys)
	}
	flamelFound := false
	dumbledoreFound := false
	for _, subkey := range subkeys {
		if subkey == "flamel" {
			flamelFound = true
		} else if subkey == "dumbledore" {
			dumbledoreFound = true
		}
	}
	if !flamelFound {
		t.Errorf("flamel not found in %#v", subkeys)
	}
	if !dumbledoreFound {
		t.Errorf("dumbledore not found in %#v", subkeys)
	}
}

func TestTaskDataValuesGetValue(t *testing.T) {
	td, err := setupTaskData(t)
	if err != nil {
		return
	}
	actual, err := td.GetValue("alchemist", "flamel", "name")
	if err != nil {
		t.Error(err)
	}
	expected := "Nicholas Flamel"
	if actual != expected {
		t.Errorf("%s expected, but getting %s", expected, actual)
	}
}

func TestTaskDataGetTask(t *testing.T) {
	td, err := setupTaskData(t)
	if err != nil {
		return
	}
	expected := "9000"
	other, err := td.GetTask("serveStaticFiles")
	if err != nil {
		t.Error(err)
	}
	actual, err := other.GetConfig("port")
	if err != nil {
		t.Error(err)
	}
	if actual != expected {
		t.Errorf("%s expected, but getting %s", expected, actual)
	}
}

func TestTaskDataGetInvalidTask(t *testing.T) {
	td, err := setupTaskData(t)
	if err != nil {
		return
	}
	_, err = td.GetTask("invalidTask")
	if err == nil {
		t.Error("Error expected, but no error found")
	}
}

func TestTaskDataReadFile(t *testing.T) {
	td, err := setupTaskData(t)
	if err != nil {
		return
	}
	content, err := td.ReadFile("../resources/a.txt")
	if err != nil {
		t.Error(err)
		return
	}
	expected := "{{ .GetEnv \"HTTP_PORT\" }}"
	if content != expected {
		t.Errorf("%s is expected, but getting %s", expected, content)
	}
}

func TestTaskDataListDir(t *testing.T) {
	td, err := setupTaskData(t)
	if err != nil {
		return
	}
	directoryList, err := td.ListDir("../resources")
	if err != nil {
		t.Error(err)
		return
	}
	if len(directoryList) != 2 {
		t.Errorf("directorryList should contain two item, currently it is %#v", directoryList)
	}
}

func TestTaskDataParseFile(t *testing.T) {
	td, err := setupTaskData(t)
	if err != nil {
		return
	}
	content, err := td.ParseFile("../resources/a.txt")
	if err != nil {
		t.Error(err)
		return
	}
	expected, err := td.GetEnv("HTTP_PORT")
	if err != nil {
		t.Error(err)
		return
	}
	if content != expected {
		t.Errorf("%s is expected, but getting %s", expected, content)
	}
}
