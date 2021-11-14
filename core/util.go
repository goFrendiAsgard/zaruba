package core

import (
	"bytes"
	"path/filepath"
	"regexp"
	"strings"

	booleanutil "github.com/state-alchemists/zaruba/core/booleanutil"
	"github.com/state-alchemists/zaruba/core/fileutil"
	"github.com/state-alchemists/zaruba/core/strutil"
)

type Util struct {
	Str  *strutil.StrUtil
	File *fileutil.FileUtil
	Bool *booleanutil.BooleanUtil
}

func NewUtil() *Util {
	util := &Util{
		Str:  strutil.NewStrutil(),
		File: fileutil.NewFileUtil(),
		Bool: booleanutil.NewBooleanUtil(),
	}
	return util
}

func (util *Util) GetDefaultAppName(location string) (appName string, err error) {
	absPath, err := filepath.Abs(location)
	if err != nil {
		return "", err
	}
	baseName := filepath.Base(absPath)
	pattern := regexp.MustCompile(`[^A-Za-z0-9]`)
	spacedBaseName := (pattern.ReplaceAllString(baseName, " "))
	titledBaseName := strings.Title(spacedBaseName)
	appName = strings.ReplaceAll(titledBaseName, " ", "")
	if len(appName) > 0 {
		bts := []byte(appName)
		lc := bytes.ToLower([]byte{bts[0]})
		rest := bts[1:]
		appName = string(bytes.Join([][]byte{lc, rest}, nil))
	}
	return appName, err
}
