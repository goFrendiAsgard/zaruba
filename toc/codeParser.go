package toc

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/state-alchemists/zaruba/dsl"
)

func ParseCode(util *dsl.DSLUtil, dirPath, content string) (newContent string, err error) {
	codeTagPattern := regexp.MustCompile(`(?m)(?s)<!--startCode(.*?)-->(.*?)<!--endCode-->`)
	matches := codeTagPattern.FindAllStringSubmatch(content, -1)
	if len(matches) == 0 {
		return content, nil
	}
	newContent = content
	for _, match := range matches {
		attributeStr, oldCode := match[1], match[2]
		lang, src, cmdStr := getCodeTagAttributes(attributeStr)
		srcContent, err := getSrcContent(util, dirPath, src, cmdStr)
		if err != nil {
			return content, err
		}
		cmdResult, err := execCommand(dirPath, cmdStr)
		if err != nil {
			return content, err
		}
		newContent = strings.Replace(
			newContent,
			fmt.Sprintf(`<!--startCode%s-->%s<!--endCode-->`, attributeStr, oldCode),
			strings.Join([]string{
				fmt.Sprintf("<!--startCode%s-->", attributeStr),
				"__Code__",
				fmt.Sprintf("```%s", lang),
				srcContent,
				"```",
				"__Output__",
				"```",
				cmdResult,
				"```",
				"<!--endCode-->",
			}, "\n"),
			-1,
		)
	}
	return newContent, err
}

func getSrcContent(util *dsl.DSLUtil, dirPath, src string, cmdStr string) (srcContent string, err error) {
	if src != "" {
		return util.File.ReadText(filepath.Join(dirPath, src))
	}
	return cmdStr, nil
}

func execCommand(dirPath, cmdStr string) (result string, err error) {
	cmd := exec.Command("bash", "-c", cmdStr)
	cmd.Dir = dirPath
	out, err := cmd.Output()
	return string(out), err
}

func getCodeTagAttributes(attributeStr string) (lang, src, cmd string) {
	lang = getTagAttribute(attributeStr, "lang")
	if lang == "" {
		lang = "sh"
	}
	src = getTagAttribute(attributeStr, "src")
	cmd = getTagAttribute(attributeStr, "cmd")
	if cmd == "" {
		if src != "" {
			cmd = fmt.Sprintf("%s %s", lang, src)
		} else {
			cmd = `echo "undefined cmd"`
		}
	}
	return lang, src, cmd
}
