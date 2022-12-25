package toc

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/state-alchemists/zaruba/dsl"
)

type Toc struct {
	isOldFileExist bool
	oldFileName    string
	oldDirPath     string
	oldStrList     []string
	newFileName    string
	newDirPath     string
	newContent     string
	children       []*Toc
	parent         *Toc
	util           *dsl.DSLUtil
}

func NewToc(location string) (toc *Toc, err error) {
	return newToc(location, nil)
}

func newToc(location string, parent *Toc) (toc *Toc, err error) {
	util := createUtil(parent)
	absLocation, err := filepath.Abs(location)
	if err != nil {
		return toc, err
	}
	oldDirPath, oldFileName := filepath.Split(absLocation)
	isOldFileExist := true
	if _, err := os.Stat(absLocation); errors.Is(err, os.ErrNotExist) {
		isOldFileExist = false
	}
	oldStrList := []string{}
	if isOldFileExist {
		jsonString, err := util.File.ReadLines(absLocation)
		if err != nil {
			return toc, err
		}
		oldStrList, err = util.Json.ToStringList(jsonString)
		if err != nil {
			return toc, err
		}
	}
	toc = &Toc{
		isOldFileExist: isOldFileExist,
		oldDirPath:     oldDirPath,
		oldFileName:    oldFileName,
		oldStrList:     oldStrList,
		util:           util,
		children:       []*Toc{},
		parent:         parent,
	}
	return toc, nil
}

func createUtil(parent *Toc) *dsl.DSLUtil {
	if parent == nil {
		return dsl.NewDSLUtil()
	}
	return parent.util
}
