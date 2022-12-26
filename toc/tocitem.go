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
	CamelCaseTitle  string
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
	// old file location is not defined
	if tocItem.OldFileLocation == "" {
		if err := tocItem.RenderNewContentToNewFile(tocHeader, tocSubtopic); err != nil {
			return err
		}
	}
	// old file location is defined
	if tocItem.OldFileLocation != "" {
		if err := tocItem.RenderNewContentFromOldFile(tocHeader, tocSubtopic); err != nil {
			return err
		}
	}
	// also render children
	return tocItem.Children.RenderNewContent()
}

func (tocItem *TocItem) RenderNewContentFromOldFile(tocHeader, tocSubtopic string) (err error) {
	util := tocItem.Toc.Util
	fmt.Println(tocItem.OldFileLocation)
	oldFileContent, err := util.File.ReadText(tocItem.OldFileLocation)
	if err != nil {
		return err
	}
	newFileContent := replaceTag(util, startTocHeaderTag, endTocHeaderTag, oldFileContent, tocHeader)
	newFileContent = replaceTag(util, startTocSubtopicTag, endTocSubtopicTag, newFileContent, tocSubtopic)
	if err := util.File.WriteText(tocItem.NewFileLocation, newFileContent, 0755); err != nil {
		return err
	}
	if tocItem.OldFileLocation != tocItem.NewFileLocation {
		return os.Remove(tocItem.OldFileLocation)
	}
}

func (tocItem *TocItem) RenderNewContentToNewFile(tocHeader, tocSubtopic string) (err error) {
	newFileContent := strings.Join([]string{
		startTocHeaderTag,
		tocHeader,
		endTocHeaderTag,
		"",
		fmt.Sprintf("> TODO: Write about `%s`.", tocItem.Title),
		"",
		startTocSubtopicTag,
		tocSubtopic,
		endTocSubtopicTag,
	}, "\n")
	return tocItem.Toc.Util.File.WriteText(tocItem.NewFileLocation, newFileContent, 0755)
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
		pathList = append([]string{parent.CamelCaseTitle}, pathList...)
		parent = parent.Parent
	}
	pathList = append([]string{tocDirPath}, pathList...)
	tocItem.NewFileLocation = filepath.Join(pathList...)
}

func (tocItem *TocItem) getNewFileName() (fileName string) {
	if len(tocItem.Children) == 0 {
		return fmt.Sprintf("%s.md", tocItem.CamelCaseTitle)
	}
	return filepath.Join(tocItem.CamelCaseTitle, "README.md")
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
		CamelCaseTitle:  toc.Util.Str.ToCamel(title),
		OldFileLocation: oldFileLocation,
		Parent:          parent,
		Toc:             toc,
	}
	return tocItem
}
