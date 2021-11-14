package yamlStyler

import (
	"regexp"
	"strconv"
	"strings"
)

type YamlStyler func([]string) []string

func TwoSpaces(yamlLines []string) (newYamlLines []string) {
	newYamlLines = []string{}
	indentationRex := regexp.MustCompile(`^(\s*)(.*)$`)
	for _, line := range yamlLines {
		indentationMatch := indentationRex.FindStringSubmatch(line)
		indentation := indentationMatch[1]
		content := indentationMatch[2]
		halfIndentation := indentation[:len(indentation)/2]
		newYamlLines = append(newYamlLines, halfIndentation+content)
	}
	return newYamlLines
}

func FixEmoji(yamlLines []string) (newYamlLines []string) {
	newYamlLines = []string{}
	quotedEmojiRex := regexp.MustCompile(`"\\U[0-9A-F]+"`)
	for _, line := range yamlLines {
		contentB := []byte(line)
		line = string(quotedEmojiRex.ReplaceAllFunc(contentB, func(sByte []byte) (resultByte []byte) {
			result, _ := strconv.Unquote(string(sByte))
			return []byte(result)
		}))
		newYamlLines = append(newYamlLines, line)
	}
	return newYamlLines
}

func AddLineBreak(yamlLines []string) (newYamlLines []string) {
	newYamlLines = []string{}
	indentationRex := regexp.MustCompile(`^(\s*)(.*)$`)
	previousIndentation := ""
	previousContent := ""
	for _, line := range yamlLines {
		indentationMatch := indentationRex.FindStringSubmatch(line)
		indentation := indentationMatch[1]
		content := indentationMatch[2]
		if len(previousIndentation) != len(indentation) && len(indentation) <= 2 && !strings.HasPrefix(previousContent, "includes:") {
			newYamlLines = append(newYamlLines, "")
		}
		previousIndentation = indentation
		previousContent = content
		newYamlLines = append(newYamlLines, line)
	}
	return newYamlLines
}
