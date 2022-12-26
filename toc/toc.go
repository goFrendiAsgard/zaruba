package toc

import (
	"path/filepath"
	"strings"

	"github.com/state-alchemists/zaruba/dsl"
)

type Toc struct {
	FileLocation string
	FileContent  string
	Items        TocItems
	Util         *dsl.DSLUtil
}

func (toc *Toc) RenderNewContent() (err error) {
	beforeTocLines, _, afterTocLines := splitContent(toc.Util, toc.FileContent)
	if err != nil {
		return err
	}
	dirPath := filepath.Dir(toc.FileLocation)
	newContent, err := renderContent(dirPath, beforeTocLines, afterTocLines, toc.Items)
	if err != nil {
		return err
	}
	if err := toc.Util.File.WriteText(toc.FileLocation, newContent, 0755); err != nil {
		return err
	}
	return toc.Items.RenderNewContent()
}

func NewToc(fileLocation string) (toc *Toc, err error) {
	absFileLocation := fileLocation
	if !filepath.IsAbs(absFileLocation) {
		absFileLocation, err = filepath.Abs(absFileLocation)
		if err != nil {
			return toc, err
		}
	}
	util := dsl.NewDSLUtil()
	fileContent, err := util.File.ReadText(absFileLocation)
	if err != nil {
		return toc, err
	}
	toc = &Toc{
		FileLocation: absFileLocation,
		FileContent:  fileContent,
		Util:         util,
	}
	_, tocContent, _ := splitContent(util, fileContent)
	tocLines := strings.Split(tocContent, "\n")
	toc.Items, err = NewTocItems(toc, nil, 0, tocLines)
	if err != nil {
		return toc, err
	}
	return toc, err
}
