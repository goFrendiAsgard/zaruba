package dsl

import (
	"github.com/state-alchemists/zaruba/booleanutil"
	"github.com/state-alchemists/zaruba/dictutil"
	"github.com/state-alchemists/zaruba/fileutil"
	"github.com/state-alchemists/zaruba/jsonutil"
	"github.com/state-alchemists/zaruba/listutil"
	"github.com/state-alchemists/zaruba/pathutil"
	"github.com/state-alchemists/zaruba/strutil"
)

type DSLUtil struct {
	Str     *strutil.StrUtil
	List    *listutil.ListUtil
	Dict    *dictutil.DictUtil
	File    *fileutil.FileUtil
	Bool    *booleanutil.BooleanUtil
	Path    *pathutil.PathUtil
	Json    *jsonutil.JsonUtil
	Project *ProjectUtil
}

func NewDSLUtil() *DSLUtil {
	strUtil := strutil.NewStrUtil()
	listUtil := listutil.NewListUtil()
	dictUtil := dictutil.NewDictUtil()
	jsonUtil := jsonutil.NewJsonUtil(strUtil)
	fileUtil := fileutil.NewFileUtil(jsonUtil)
	projectUtil := NewProjectUtil(fileUtil, jsonUtil)
	pathUtil := pathutil.NewPathUtil(jsonUtil)
	booleanUtil := booleanutil.NewBooleanUtil()
	coreUtil := &DSLUtil{
		Str:     strUtil,
		List:    listUtil,
		Dict:    dictUtil,
		File:    fileUtil,
		Bool:    booleanUtil,
		Path:    pathUtil,
		Json:    jsonUtil,
		Project: projectUtil,
	}
	return coreUtil
}
