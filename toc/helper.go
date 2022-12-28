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

func splitContentByTag(util *dsl.DSLUtil, startTag, endTag, content string) (beforeTag, inTag, afterTag string) {
	r := regexp.MustCompile(fmt.Sprintf("(?s)(.*)%s(.*)%s(.*)", startTag, endTag))
	matches := r.FindStringSubmatch(content)
	if len(matches) > 3 {
		beforeTag = matches[1]
		inTag = matches[2]
		afterTag = matches[3]
	}
	return beforeTag, inTag, afterTag
}

func replaceTag(util *dsl.DSLUtil, startTag, endTag, content, replacement string) (newContent string) {
	beforeTag, _, afterTag := splitContentByTag(util, startTag, endTag, content)
	if len(replacement) > 0 {
		if replacement[0] != '\n' {
			replacement = "\n" + replacement
		}
		if replacement[len(replacement)-1] != '\n' {
			replacement = replacement + "\n"
		}
	}
	return strings.Join([]string{
		beforeTag,
		startTag,
		replacement,
		endTag,
		afterTag,
	}, "")
}
