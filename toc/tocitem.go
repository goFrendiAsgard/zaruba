package toc

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type TocItem struct {
	Level           int
	Title           string
	KebabCaseTitle  string
	OldFileLocation string
	NewFileLocation string
	Children        TocItems
	Parent          *TocItem
	Toc             *Toc
}

func (tocItem *TocItem) RenderNewContent() (err error) {
	tocHeader, err := tocItem.GetTocHeader()
	if err != nil {
		return err
	}
	dirPath := filepath.Dir(tocItem.NewFileLocation)
	tocSubtopic, err := tocItem.Children.AsLinks(tocItem.Level+1, dirPath)
	if err != nil {
		return err
	}
	if tocSubtopic != "" {
		tocSubtopic = strings.Join([]string{"# Subtopics", tocSubtopic}, "\n")
	}
	util := tocItem.Toc.Util
	oldFileExist, oldFileExistErr := util.File.IsExist(tocItem.OldFileLocation)
	// old file location is not defined or not exist
	if tocItem.OldFileLocation == "" || !oldFileExist || oldFileExistErr != nil {
		if err := tocItem.RenderNewContentToNewFile(dirPath, tocHeader, tocSubtopic); err != nil {
			return err
		}
		return tocItem.Children.RenderNewContent()
	}
	// old file location is defined
	if err := tocItem.RenderNewContentFromOldFile(dirPath, tocHeader, tocSubtopic); err != nil {
		return err
	}
	// also render children
	return tocItem.Children.RenderNewContent()
}

func (tocItem *TocItem) RenderNewContentFromOldFile(dirPath, tocHeader, tocSubtopic string) (err error) {
	util := tocItem.Toc.Util
	oldFileContent, err := util.File.ReadText(tocItem.OldFileLocation)
	if err != nil {
		return err
	}
	newFileContent := replaceTag(startTocHeaderTag, endTocHeaderTag, oldFileContent, tocHeader)
	newFileContent = replaceTag(startTocSubtopicTag, endTocSubtopicTag, newFileContent, tocSubtopic)
	newFileContent = replaceLink(tocItem.Toc.LinkReplacement, tocItem.OldFileLocation, tocItem.NewFileLocation, newFileContent)
	newFileContent, err = tocItem.Toc.ParseCodeTag(newFileContent)
	if err != nil {
		return err
	}
	if err := util.File.WriteText(tocItem.NewFileLocation, newFileContent, 0755); err != nil {
		return err
	}
	if tocItem.OldFileLocation != tocItem.NewFileLocation {
		oldFileExist, oldFileExistErr := util.File.IsExist(tocItem.OldFileLocation)
		if oldFileExist && oldFileExistErr == nil {
			return os.Remove(tocItem.OldFileLocation)
		}
	}
	return nil
}

func (tocItem *TocItem) RenderNewContentToNewFile(dirPath, tocHeader, tocSubtopic string) (err error) {
	newFileContent := strings.Join([]string{
		startTocHeaderTag,
		tocHeader,
		endTocHeaderTag,
		"",
		fmt.Sprintf("> TODO: Write about `%s`.", tocItem.Title),
		"",
		tocItem.GetNewTaggedSubtopicContent(tocSubtopic),
	}, "\n")
	newFileContent, err = tocItem.Toc.ParseCodeTag(newFileContent)
	if err != nil {
		return err
	}
	return tocItem.Toc.Util.File.WriteText(tocItem.NewFileLocation, newFileContent, 0755)
}

func (tocItem *TocItem) GetNewTaggedSubtopicContent(tocSubtopic string) string {
	if tocSubtopic == "" {
		return strings.Join([]string{startTocSubtopicTag, endTocSubtopicTag}, "\n")
	}
	return strings.Join([]string{startTocSubtopicTag, tocSubtopic, endTocSubtopicTag}, "\n")
}

func (tocItem *TocItem) GetTocHeader() (tocHeader string, err error) {
	parentLink, err := tocItem.GetParentLinksAsString()
	if err != nil {
		return "", err
	}
	tocHeader = strings.Join([]string{
		parentLink,
		fmt.Sprintf("# %s", tocItem.Title),
	}, "\n")
	return tocHeader, nil
}

func (tocItem *TocItem) AsLinks(level int, dirPath string) (links string, err error) {
	singleLink, err := tocItem.AsLink(dirPath)
	if err != nil {
		return "", err
	}
	lines := []string{fmt.Sprintf(`%s- %s`, strings.Repeat("  ", tocItem.Level-level), singleLink)}
	if len(tocItem.Children) != 0 {
		childrenLines, err := tocItem.Children.AsLinks(level, dirPath)
		if err != nil {
			return "", err
		}
		lines = append(lines, childrenLines)
	}
	return strings.Join(lines, "\n"), nil
}

func (tocItem *TocItem) AsLink(dirPath string) (link string, err error) {
	relativeUrl, err := filepath.Rel(dirPath, tocItem.NewFileLocation)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf(`[%s](%s)`, tocItem.Title, relativeUrl), nil
}

func (tocItem *TocItem) GetParentLinksAsString() (str string, err error) {
	dirPath := filepath.Dir(tocItem.NewFileLocation)
	parent := tocItem.Parent
	links := []string{}
	// add link to parent files
	for parent != nil {
		absUrl := parent.NewFileLocation
		relativeUrl, err := filepath.Rel(dirPath, absUrl)
		if err != nil {
			return "", err
		}
		link := fmt.Sprintf("[%s](%s)", parent.Title, relativeUrl)
		links = append([]string{link}, links...)
		parent = parent.Parent
	}
	// add link to TOC file
	absUrl := tocItem.Toc.FileLocation
	relativeUrl, err := filepath.Rel(dirPath, absUrl)
	if err != nil {
		return "", err
	}
	link := fmt.Sprintf("[🏠](%s)", relativeUrl)
	links = append([]string{link}, links...)
	return strings.Join(links, " > "), nil
}

func (tocItem *TocItem) SetNewFileLocation() {
	tocFileLocation := tocItem.Toc.FileLocation
	tocDirPath := filepath.Dir(tocFileLocation)
	pathList := []string{tocItem.getNewFileName()}
	parent := tocItem.Parent
	for parent != nil {
		pathList = append([]string{parent.KebabCaseTitle}, pathList...)
		parent = parent.Parent
	}
	pathList = append([]string{tocDirPath}, pathList...)
	tocItem.NewFileLocation = filepath.Join(pathList...)
}

func (tocItem *TocItem) getNewFileName() (fileName string) {
	if len(tocItem.Children) == 0 {
		return fmt.Sprintf("%s.md", tocItem.KebabCaseTitle)
	}
	return filepath.Join(tocItem.KebabCaseTitle, "README.md")
}

func (tocItem *TocItem) GetLinkReplacement() (linkReplacement map[string]string) {
	linkReplacement = map[string]string{}
	if tocItem.OldFileLocation != "" {
		linkReplacement[tocItem.OldFileLocation] = tocItem.NewFileLocation
	}
	for _, child := range tocItem.Children {
		childLinkReplacement := child.GetLinkReplacement()
		for oldFileLocation, newFileLocation := range childLinkReplacement {
			linkReplacement[oldFileLocation] = newFileLocation
		}
	}
	return linkReplacement
}

func NewTocItem(toc *Toc, parent *TocItem, level int, title, oldFileLocation string) (tocItem *TocItem) {
	if oldFileLocation != "" && !filepath.IsAbs(oldFileLocation) {
		tocFileLocation := toc.FileLocation
		tocDirPath := filepath.Dir(tocFileLocation)
		oldFileLocation = filepath.Join(tocDirPath, oldFileLocation)
	}
	tocItem = &TocItem{
		Level:           level,
		Title:           title,
		KebabCaseTitle:  toc.Util.Str.ToKebab(title),
		OldFileLocation: oldFileLocation,
		Parent:          parent,
		Toc:             toc,
	}
	return tocItem
}
