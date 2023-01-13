package toc

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/state-alchemists/zaruba/dsl"
)

type Toc struct {
	FileLocation    string
	FileContent     string
	Items           TocItems
	Util            *dsl.DSLUtil
	LinkReplacement map[string]string
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
	newContent := replaceTag(startTocTag, endTocTag, toc.FileContent, newTocSection)
	newContent = replaceLink(toc.LinkReplacement, toc.FileLocation, toc.FileLocation, newContent)
	newContent, err = toc.ParseCodeTag(newContent)
	if err != nil {
		return err
	}
	if err := toc.Util.File.WriteText(toc.FileLocation, newContent, 0755); err != nil {
		return err
	}
	// render toc items
	return toc.Items.RenderNewContent()
}

func (toc *Toc) ParseCodeTag(content string) (newContent string, err error) {
	codeTagPattern := regexp.MustCompile(`(?m)(?s)<!--startCode(.*?)-->(.*?)<!--endCode-->`)
	matches := codeTagPattern.FindAllStringSubmatch(content, -1)
	if len(matches) == 0 {
		return content, nil
	}
	newContent = content
	for _, match := range matches {
		attributeStr, oldCodeTagContent := match[1], match[2]
		lang, src, cmdStr := toc.getCodeTagAttributes(attributeStr)
		srcContent, err := toc.getCodeTagSrcContent(src, cmdStr)
		if err != nil {
			return content, err
		}
		cmdResult, err := toc.execCommand(cmdStr)
		if err != nil {
			return content, err
		}
		newContent = strings.Replace(
			newContent,
			fmt.Sprintf(`<!--startCode%s-->%s<!--endCode-->`, attributeStr, oldCodeTagContent),
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

func (toc *Toc) getCodeTagSrcContent(src string, cmdStr string) (srcContent string, err error) {
	util := toc.Util
	dirPath := filepath.Dir(toc.FileLocation)
	if src != "" {
		return util.File.ReadText(filepath.Join(dirPath, src))
	}
	return cmdStr, nil
}

func (toc *Toc) execCommand(cmdStr string) (result string, err error) {
	cmd := exec.Command("bash", "-c", cmdStr)
	cmd.Dir = filepath.Dir(toc.FileLocation)
	out, err := cmd.Output()
	return string(out), err
}

func (toc *Toc) getCodeTagAttributes(attributeStr string) (lang, src, cmd string) {
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
	_, tocContent, _, isTagFound := splitContentByTag(startTocTag, endTocTag, fileContent)
	if !isTagFound {
		return toc, fmt.Errorf("no tag found at '%s', expecting '%s' and '%s'", absFileLocation, startTocTag, endTocTag)
	}
	tocLines := strings.Split(tocContent, "\n")
	toc.Items, err = NewTocItems(toc, nil, 0, tocLines)
	if err != nil {
		return toc, err
	}
	toc.LinkReplacement = toc.Items.GetLinkReplacement()
	return toc, err
}
