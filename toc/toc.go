package toc

import (
	"fmt"
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
	dirPath := filepath.Dir(toc.FileLocation)
	newTocSection := ""
	if len(toc.Items) > 0 {
		newTocSection, err = toc.Items.AsLinks(0, dirPath)
		if err != nil {
			return err
		}
	}
	newContent := replaceTag(toc.Util, startTocTag, endTocTag, toc.FileContent, newTocSection)
	newContent, err = ParseCode(toc.Util, dirPath, newContent)
	if err != nil {
		return err
	}
	if err := toc.Util.File.WriteText(toc.FileLocation, newContent, 0755); err != nil {
		return err
	}
	// render toc items
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
	_, tocContent, _, isTagFound := splitContentByTag(util, startTocTag, endTocTag, fileContent)
	if !isTagFound {
		return toc, fmt.Errorf("no tag found at '%s', expecting '%s' and '%s'", absFileLocation, startTocTag, endTocTag)
	}
	tocLines := strings.Split(tocContent, "\n")
	toc.Items, err = NewTocItems(toc, nil, 0, tocLines)
	if err != nil {
		return toc, err
	}
	return toc, err
}
