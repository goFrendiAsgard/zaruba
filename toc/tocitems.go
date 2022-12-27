package toc

import (
	"regexp"
	"strings"
)

type TocItems []*TocItem

func (tocItems TocItems) AsLinks(level int, dirPath string) (str string, err error) {
	lines := []string{}
	for _, tocItem := range tocItems {
		line, err := tocItem.AsLinks(level, dirPath)
		if err != nil {
			return str, err
		}
		lines = append(lines, line)
	}
	return strings.Join(lines, "\n"), nil
}

func (tocItems TocItems) RenderNewContent() (err error) {
	for _, tocItem := range tocItems {
		if err := tocItem.RenderNewContent(); err != nil {
			return err
		}
	}
	return nil
}

func NewTocItems(toc *Toc, parent *TocItem, level int, lines []string) (tocItems TocItems, err error) {
	tocItems = []*TocItem{}
	completePattern, _ := regexp.Compile(`^([ \t]*)[\*-] \[(.*)\]\((.*)\).*$`)
	minimalPattern, _ := regexp.Compile(`^([ \t]*)[\*-] (.*)$`)
	baseIndentation := ""
	var lastItem *TocItem
	childLines := []string{}
	firstMatch := true
	for _, line := range lines {
		match := completePattern.FindStringSubmatch(line)
		if len(match) == 0 {
			// complete pattern failed, try minimal pattern
			match = minimalPattern.FindStringSubmatch(line)
		}
		// complete/minimal pattern match
		if len(match) > 0 {
			indentation := match[1]
			title := match[2]
			fileLocation := ""
			if len(match) > 3 {
				fileLocation = match[3]
			}
			// get baseIndentation
			if firstMatch {
				baseIndentation = indentation
				lastItem = NewTocItem(toc, parent, level, title, fileLocation)
				tocItems = append(tocItems, lastItem)
				firstMatch = false
				continue
			}
			// childLines
			if len(indentation) > len(baseIndentation) {
				childLines = append(childLines, line)
				continue
			}
			// new item
			if len(indentation) >= len(baseIndentation) {
				// wrap up previous lastItem
				if len(childLines) > 0 {
					// make sure we only do this when childLines exist
					lastItem.Children, err = NewTocItems(toc, lastItem, level+1, childLines)
					if err != nil {
						return tocItems, err
					}
					childLines = []string{}
				}
				lastItem.SetNewFileLocation()
				// set current line as last item
				lastItem = NewTocItem(toc, parent, level, title, fileLocation)
				tocItems = append(tocItems, lastItem)
			}
		}
	}
	// there are unfinished businesses
	if lastItem != nil {
		if len(childLines) > 0 {
			// make sure we only do this when childLines exist
			lastItem.Children, err = NewTocItems(toc, lastItem, level+1, childLines)
			if err != nil {
				return tocItems, err
			}
		}
		lastItem.SetNewFileLocation()
	}
	return tocItems, err
}
