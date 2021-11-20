package core

import (
	"github.com/state-alchemists/zaruba/booleanutil"
	"github.com/state-alchemists/zaruba/fileutil"
	"github.com/state-alchemists/zaruba/jsonutil"
	"github.com/state-alchemists/zaruba/pathutil"
	"github.com/state-alchemists/zaruba/strutil"
)

type CoreUtil struct {
	Str     *strutil.StrUtil
	File    *fileutil.FileUtil
	Bool    *booleanutil.BooleanUtil
	Path    *pathutil.PathUtil
	Json    *jsonutil.JsonUtil
	Project *ProjectUtil
}

func NewCoreUtil() *CoreUtil {
	jsonUtil := jsonutil.NewJsonUtil()
	fileUtil := fileutil.NewFileUtil(jsonUtil)
	projectUtil := NewProjectUtil(fileUtil)
	coreUtil := &CoreUtil{
		Str:     strutil.NewStrutil(),
		File:    fileUtil,
		Bool:    booleanutil.NewBooleanUtil(),
		Path:    pathutil.NewPathUtil(),
		Json:    jsonUtil,
		Project: projectUtil,
	}
	return coreUtil
}
