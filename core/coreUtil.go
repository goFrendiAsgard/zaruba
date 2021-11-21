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
	strUtil := strutil.NewStrUtil()
	jsonUtil := jsonutil.NewJsonUtil(strUtil)
	fileUtil := fileutil.NewFileUtil(jsonUtil)
	projectUtil := NewProjectUtil(fileUtil, jsonUtil)
	pathUtil := pathutil.NewPathUtil(jsonUtil)
	booleanUtil := booleanutil.NewBooleanUtil()
	coreUtil := &CoreUtil{
		Str:     strUtil,
		File:    fileUtil,
		Bool:    booleanUtil,
		Path:    pathUtil,
		Json:    jsonUtil,
		Project: projectUtil,
	}
	return coreUtil
}
