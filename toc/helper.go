package toc

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strings"
)

var startTocTag = "<!--startToc-->"
var endTocTag = "<!--endToc-->"
var startTocHeaderTag = "<!--startTocHeader-->"
var endTocHeaderTag = "<!--endTocHeader-->"
var startTocSubtopicTag = "<!--startTocSubtopic-->"
var endTocSubtopicTag = "<!--endTocSubtopic-->"

func replaceLink(linkReplacement map[string]string, oldFileLocation, newFileLocation, content string) (newContent string) {
	if oldFileLocation == "" {
		return content
	}
	oldDirLocation := filepath.Dir(oldFileLocation)
	newDirLocation := filepath.Dir(newFileLocation)
	newContent = content
	for oldFileLocation, newFileLocation := range linkReplacement {
		relativeOldLink, err := filepath.Rel(oldDirLocation, oldFileLocation)
		if err != nil {
			continue
		}
		relativeNewLink, err := filepath.Rel(newDirLocation, newFileLocation)
		if err != nil {
			continue
		}
		newContent = strings.Replace(newContent, relativeOldLink, relativeNewLink, -1)
	}
	return newContent
}

func splitContentByTag(startTag, endTag, content string) (beforeTag, inTag, afterTag string, isTagFound bool) {
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

func replaceTag(startTag, endTag, content, replacement string) (newContent string) {
	beforeTag, _, afterTag, isTagFound := splitContentByTag(startTag, endTag, content)
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
