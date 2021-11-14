package core

import (
	"github.com/state-alchemists/zaruba/booleanutil"
	"github.com/state-alchemists/zaruba/fileutil"
	"github.com/state-alchemists/zaruba/pathutil"
	"github.com/state-alchemists/zaruba/strutil"
)

type CoreUtil struct {
	Str  *strutil.StrUtil
	File *fileutil.FileUtil
	Bool *booleanutil.BooleanUtil
	Path *pathutil.PathUtil
}

func NewCoreUtil() *CoreUtil {
	coreUtil := &CoreUtil{
		Str:  strutil.NewStrutil(),
		File: fileutil.NewFileUtil(),
		Bool: booleanutil.NewBooleanUtil(),
		Path: pathutil.NewPathUtil(),
	}
	return coreUtil
}
