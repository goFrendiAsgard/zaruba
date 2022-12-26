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
	var lastLink *TocItem
	childLines := []string{}
	for index, line := range lines {
		match := completePattern.FindStringSubmatch(line)
		if len(match) == 0 {
			match = minimalPattern.FindStringSubmatch(line)
		}
		// skip anything that is not a list (could be an explanation)
		if len(match) == 0 {
			continue
		}
		indentation := match[1]
		title := match[2]
		fileLocation := ""
		if len(match) > 3 {
			fileLocation = match[3]
		}
		// get baseIndentation
		if index == 0 {
			baseIndentation = indentation
		}
		// childLines
		if len(indentation) > len(baseIndentation) {
			childLines = append(childLines, line)
		}
		// add childLines
		if indentation == baseIndentation || index == len(lines)-1 {
			if lastLink != nil {
				lastLink.Children, err = NewTocItems(toc, lastLink, level+1, childLines)
				if err != nil {
					return tocItems, err
				}
				lastLink.SetNewFileLocation()
			}
		}
		// add link
		if indentation == baseIndentation {
			childLines = []string{}
			lastLink = NewTocItem(toc, parent, level, title, fileLocation)
			tocItems = append(tocItems, lastLink)
		}
		// set link new location
		if indentation == baseIndentation || index == len(lines)-1 {
			if lastLink != nil {
				lastLink.SetNewFileLocation()
			}
		}
	}
	return tocItems, err
}
