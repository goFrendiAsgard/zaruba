package toc

import (
	"fmt"
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

func splitContentByTag(util *dsl.DSLUtil, startTag, endTag, content string) (beforeTag, inTag, afterTag string, isTagFound bool) {
	r := regexp.MustCompile(fmt.Sprintf("(?s)(.*)%s(.*)%s(.*)", startTag, endTag))
	matches := r.FindStringSubmatch(content)
	isTagFound = false
	if len(matches) > 3 {
		beforeTag = matches[1]
		inTag = matches[2]
		afterTag = matches[3]
		isTagFound = true
	}
	return beforeTag, inTag, afterTag, isTagFound
}

func replaceTag(util *dsl.DSLUtil, startTag, endTag, content, replacement string) (newContent string) {
	beforeTag, _, afterTag, isTagFound := splitContentByTag(util, startTag, endTag, content)
	if !isTagFound {
		return content
	}
	if len(replacement) > 0 {
		if replacement[0] != '\n' {
			replacement = "\n" + replacement
		}
		if replacement[len(replacement)-1] != '\n' {
			replacement = replacement + "\n"
		}
	}
	if replacement == "" {
		replacement = "\n"
	}
	return strings.Join([]string{
		beforeTag,
		startTag,
		replacement,
		endTag,
		afterTag,
	}, "")
}

func getTagAttribute(attributeStr string, attributeName string) (attributeVal string) {
	attributeDelimiterPattern := regexp.MustCompile(fmt.Sprintf(`%s[\s]*=[\s]*(.)`, attributeName))
	delimiterMatch := attributeDelimiterPattern.FindStringSubmatch(attributeStr)
	delimiter := "\""
	if len(delimiterMatch) > 0 {
		delimiter = delimiterMatch[1]
	}
	attributePattern := regexp.MustCompile(fmt.Sprintf(`%s[\s]*=[\s]*%s([^%s]*)%s.*`, attributeName, delimiter, delimiter, delimiter))
	match := attributePattern.FindStringSubmatch(attributeStr)
	if len(match) > 0 {
		return match[1]
	}
	return ""
}
