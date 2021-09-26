package file

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/state-alchemists/zaruba/utility"
	"gopkg.in/yaml.v3"
)

func Copy(src, dst string) (byteCount int64, err error) {
	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()
	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

func ReadText(fileName string) (text string, err error) {
	contentB, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	return string(contentB), nil
}

func WriteText(fileName string, text string, fileMode os.FileMode) (err error) {
	if err := os.MkdirAll(filepath.Dir(fileName), fileMode); err != nil {
		return err
	}
	return ioutil.WriteFile(fileName, []byte(text), fileMode)
}

func ReadLines(fileName string) (lines []string, err error) {
	content, err := ReadText(fileName)
	if err != nil {
		return []string{}, err
	}
	return strings.Split(content, "\n"), nil
}

func WriteLines(fileName string, lines []string, fileMode os.FileMode) (err error) {
	content := strings.Join(lines, "\n")
	return WriteText(fileName, content, fileMode)
}

func ReadYaml(fileName string) (node *yaml.Node, err error) {
	var nodeObj yaml.Node
	node = &nodeObj
	yamlScript, err := ReadText(fileName)
	if err != nil {
		return node, err
	}
	if err = yaml.Unmarshal([]byte(yamlScript), node); err != nil {
		return node, err
	}
	// set default kind for node
	if node.Kind == 0 {
		node.Kind = yaml.DocumentNode
	}
	// set default content for mode
	if len(node.Content) == 0 {
		node.Content = []*yaml.Node{
			{Kind: yaml.MappingNode, Content: []*yaml.Node{}},
		}
	}
	return node, err
}

type YamlLinesPreprocessors func([]string) []string

func YamlTwoSpace(yamlLines []string) (newYamlLines []string) {
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

func YamlFixEmoji(yamlLines []string) (newYamlLines []string) {
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

func YamlAddLineBreakForTwoSpaceIndented(yamlLines []string) (newYamlLines []string) {
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

func WriteYaml(fileName string, node *yaml.Node, fileMode os.FileMode, yamlLinesPreprocessors []YamlLinesPreprocessors) (err error) {
	yamlContentB, err := yaml.Marshal(node)
	if err != nil {
		return err
	}
	yamlContent := string(yamlContentB)
	yamlLines := strings.Split(yamlContent, "\n")
	for _, preprocessor := range yamlLinesPreprocessors {
		yamlLines = preprocessor(yamlLines)
	}
	return WriteLines(fileName, yamlLines, fileMode)
}

func ListDir(dirPath string) (fileNames []string, err error) {
	fileNames = []string{}
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return fileNames, err
	}
	for _, f := range files {
		fileNames = append(fileNames, f.Name())
	}
	return fileNames, nil
}

func Generate(sourceTemplatePath, destinationPath string, replacementMap map[string]string) (err error) {
	absSourceTemplatePath, err := filepath.Abs(sourceTemplatePath)
	if err != nil {
		return err
	}
	absDestinationPath, err := filepath.Abs(destinationPath)
	if err != nil {
		return err
	}
	util := utility.NewUtil()
	return filepath.Walk(absSourceTemplatePath,
		func(absSourceLocation string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			relativeLocation := absSourceLocation[len(absSourceTemplatePath):]
			absDestinationLocation := filepath.Join(absDestinationPath, util.Str.Replace(relativeLocation, replacementMap))
			fileMode := info.Mode()
			if info.IsDir() {
				os.Mkdir(absDestinationLocation, fileMode)
				return nil
			}
			content, err := ReadText(absSourceLocation)
			if err != nil {
				return err
			}
			newContent := util.Str.Replace(content, replacementMap)
			if newContent == content {
				_, err = Copy(absSourceLocation, absDestinationLocation)
				return err
			}
			return WriteText(absDestinationLocation, newContent, fileMode)
		},
	)
}
