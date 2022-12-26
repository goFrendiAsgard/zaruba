package toc

import (
	"fmt"
	"path/filepath"
	"strings"
)

type TocItem struct {
	Level           int
	Title           string
	CamelCaseTitle  string
	FileLocation    string
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
	if tocItem.FileLocation == "" {
		return tocItem.RenderNewContentToNewFile(tocHeader, tocSubtopic)
	}
	return tocItem.RenderNewContentFromOldFile(tocHeader, tocSubtopic)
}

func (tocItem *TocItem) RenderNewContentFromOldFile(tocHeader, tocSubtopic string) (err error) {
	return nil
}

func (tocItem *TocItem) RenderNewContentToNewFile(tocHeader, tocSubtopic string) (err error) {
	content := strings.Join([]string{
		"<!--startTocHeader-->",
		tocHeader,
		"<!--endTocHeader-->",
		"",
		fmt.Sprintf("> TODO: Write about `%s`.", tocItem.Title),
		"",
		"<!--startTocSubtopic-->",
		tocSubtopic,
		"<!--endTocSubtopic-->",
	}, "\n")
	if err := tocItem.Toc.Util.File.WriteText(tocItem.NewFileLocation, content, 0755); err != nil {
		return err
	}
	return tocItem.Children.RenderNewContent()
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
	pathList := []string{tocItem.getFileName()}
	parent := tocItem.Parent
	for parent != nil {
		pathList = append([]string{parent.CamelCaseTitle}, pathList...)
		parent = parent.Parent
	}
	pathList = append([]string{tocDirPath}, pathList...)
	tocItem.NewFileLocation = filepath.Join(pathList...)
}

func (tocItem *TocItem) getFileName() (fileName string) {
	if len(tocItem.Children) == 0 {
		return fmt.Sprintf("%s.md", tocItem.CamelCaseTitle)
	}
	return filepath.Join(tocItem.CamelCaseTitle, "README.md")
}

func NewTocItem(toc *Toc, parent *TocItem, level int, title, fileLocation string) (tocItem *TocItem) {
	tocItem = &TocItem{
		Level:          level,
		Title:          title,
		CamelCaseTitle: toc.Util.Str.ToCamel(title),
		FileLocation:   fileLocation,
		Parent:         parent,
		Toc:            toc,
	}
	return tocItem
}
