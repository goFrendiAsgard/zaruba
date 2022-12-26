package toc

import (
	"regexp"
	"strings"

	"github.com/state-alchemists/zaruba/dsl"
)

var startTocTag = "<!--startToc-->"
var endTocTag = "<!--endToc-->"
var startTocHeaderTag = "<!--startTocHeader-->"
var endTocHeaderTag = "<!--endTocHeader-->"
var startTocSubtopicTag = "<!--startTocSubtopic-->"
var endTocSubtopicTag = "<!--endTocSubtopic-->"

func splitContentByTag(util *dsl.DSLUtil, startTag, endTag, content string) (beforeTag, inTag, afterTag string) {
	lines := strings.Split(content, "\n")
	beforeTagIndex, afterTagIndex := -1, -1
	startTagPattern, _ := regexp.Compile(startTag)
	endTagPattern, _ := regexp.Compile(endTag)
	for index, line := range lines {
		if startTagPattern.Match([]byte(line)) {
			beforeTagIndex = index
		}
		if endTagPattern.Match([]byte(line)) {
			afterTagIndex = index
		}
	}
	if beforeTagIndex < 0 || afterTagIndex < 0 {
		// tag incomplete or not found
		return content, "", ""
	}
	beforeTag = strings.Join(lines[:beforeTagIndex], "\n")
	if afterTagIndex > beforeTagIndex+1 {
		// only process this if there is really something inside the tag
		inTag = strings.Join(lines[beforeTagIndex+1:afterTagIndex-1], "\n")
	}
	afterTag = strings.Join(lines[afterTagIndex+1:], "\n")
	return beforeTag, inTag, afterTag
}

func replaceTag(util *dsl.DSLUtil, startTag, endTag, content, replacement string) (newContent string) {
	beforeTag, _, afterTag := splitContentByTag(util, startTag, endTag, content)
	stringList := []string{}
	if beforeTag != "" {
		stringList = append(stringList, beforeTag)
	}
	stringList = append(stringList, startTag, replacement, endTag)
	if afterTag != "" {
		stringList = append(stringList, afterTag)
	}
	return strings.Join(stringList, "\n")
}
