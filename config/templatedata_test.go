package config

import (
	"path/filepath"
	"testing"
)

var td *TaskData

func setupTaskData(t *testing.T) (err error) {
	if td != nil {
		return err
	}
	tdConf, err := NewConfig("../test_resource/valid/zaruba.yaml")
	if err != nil {
		t.Error(err)
		return err
	}
	kwargs := []string{
		"alchemist::flamel::name=Nicholas Flamel",
		"alchemist::flamel::age=90",
		"alchemist::dumbledore::name=Dumbledore",
		"alchemist::dumbledore::age=90",
		"barbarian::sonya::name=Sonya",
		"alchemist::=unknown",
	}
	for _, kwarg := range kwargs {
		if err = tdConf.AddKwargs(kwarg); err != nil {
			return err
		}
	}
	td = NewTaskData(tdConf.Tasks["runApiGateway"])
	return err
}

func TestTaskDataGetWorkPath(t *testing.T) {
	if err := setupTaskData(t); err != nil {
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
	if err := setupTaskData(t); err != nil {
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
	if err := setupTaskData(t); err != nil {
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
	if err := setupTaskData(t); err != nil {
		return
	}
	if _, err := td.GetConfig("checkPort"); err != nil {
		t.Error("config checkPort does not exist")
	}
}

func TestTaskDataGetAllConfig(t *testing.T) {
	if err := setupTaskData(t); err != nil {
		return
	}
	if _, err := td.GetAllConfig(); err != nil {
		t.Error(err)
	}
}

func TestTaskDataGetLConfig(t *testing.T) {
	if err := setupTaskData(t); err != nil {
		return
	}
	if _, exist := td.GetLConfig("tags"); exist != nil {
		t.Error("lconfig tags does not exist")
	}
}

func TestTaskDataGetAllLConfig(t *testing.T) {
	if err := setupTaskData(t); err != nil {
		return
	}
	if _, err := td.GetAllLConfig(); err != nil {
		t.Error(err)
	}
}

func TestTaskDataGetKwarg(t *testing.T) {
	if err := setupTaskData(t); err != nil {
		return
	}
	if _, exist := td.GetKwarg("alchemist::flamel::age"); exist != nil {
		t.Error("kwarg alchemist::flamel::age does not exist")
	}
}

func TestTaskDataGetAllKwargs(t *testing.T) {
	if err := setupTaskData(t); err != nil {
		return
	}
	if _, err := td.GetAllKwargs(); err != nil {
		t.Error(err)
	}
}

func TestTaskDataGetEnv(t *testing.T) {
	if err := setupTaskData(t); err != nil {
		return
	}
	if _, exist := td.GetEnv("HTTP_PORT"); exist != nil {
		t.Error("env HTTP_PORT does not exist")
	}
}

func TestTaskDataGetAllEnv(t *testing.T) {
	if err := setupTaskData(t); err != nil {
		return
	}
	if _, err := td.GetAllEnv(); err != nil {
		t.Error(err)
	}
}

func TestTaskDataKwargsGetSubKeys(t *testing.T) {
	if err := setupTaskData(t); err != nil {
		return
	}
	subkeys := td.GetKwargSubKeys("alchemist")
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

func TestTaskDataKwargsGetValue(t *testing.T) {
	if err := setupTaskData(t); err != nil {
		return
	}
	actual, err := td.GetKwarg("alchemist", "flamel", "name")
	if err != nil {
		t.Error(err)
	}
	expected := "Nicholas Flamel"
	if actual != expected {
		t.Errorf("%s expected, but getting %s", expected, actual)
	}
}

func TestTaskDataGetTask(t *testing.T) {
	if err := setupTaskData(t); err != nil {
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
	if err := setupTaskData(t); err != nil {
		return
	}
	_, err := td.GetTask("invalidTask")
	if err == nil {
		t.Error("Error expected, but no error found")
	}
}
