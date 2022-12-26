package toc

import (
	"regexp"
	"strings"

	"github.com/state-alchemists/zaruba/dsl"
)

func splitContent(util *dsl.DSLUtil, content string) (beforeToc, tocSection, afterToc string) {
	lines := strings.Split(content, "\n")
	beforeTocIndex, afterTocIndex := 0, 0
	startToc, _ := regexp.Compile(`<!--startToc-->`)
	endToc, _ := regexp.Compile(`<!--endToc-->`)
	for index, line := range lines {
		if startToc.Match([]byte(line)) {
			beforeTocIndex = index
		}
		if endToc.Match([]byte(line)) {
			afterTocIndex = index
		}
	}
	beforeToc = strings.Join(lines[:beforeTocIndex], "\n")
	tocSection = strings.Join(lines[beforeTocIndex+1:afterTocIndex-1], "\n")
	afterToc = strings.Join(lines[afterTocIndex+1:], "\n")
	return beforeToc, tocSection, afterToc
}

func renderContent(dirPath, beforeTocContent, afterTocContent string, tocItems TocItems) (content string, err error) {
	tocSection := ""
	if len(tocItems) > 0 {
		tocSection, err = tocItems.AsLinks(0, dirPath)
		if err != nil {
			return "", err
		}
	}
	content = strings.Join([]string{
		beforeTocContent,
		"<!--startToc-->",
		tocSection,
		"<!--endToc-->",
		afterTocContent,
	}, "\n")
	return content, nil
}
