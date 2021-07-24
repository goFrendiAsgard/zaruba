package file

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/state-alchemists/zaruba/str"
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

func WriteYaml(fileName string, node *yaml.Node, fileMode os.FileMode) (err error) {
	yamlContentB, err := yaml.Marshal(node)
	if err != nil {
		return err
	}
	yamlContent := string(yamlContentB)
	yamlLines := strings.Split(yamlContent, "\n")
	rex := regexp.MustCompile(`^(\s*)(.*)$`)
	for lineIndex, line := range yamlLines {
		match := rex.FindStringSubmatch(line)
		indentation := match[1]
		content := match[2]
		halfIndentation := indentation[:len(indentation)/2]
		yamlLines[lineIndex] = halfIndentation + content
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
	return filepath.Walk(absSourceTemplatePath,
		func(absSourceLocation string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			relativeLocation := absSourceLocation[len(absSourceTemplatePath):]
			absDestinationLocation := filepath.Join(absDestinationPath, str.ReplaceByMap(relativeLocation, replacementMap))
			fileMode := info.Mode()
			if info.IsDir() {
				os.Mkdir(absDestinationLocation, fileMode)
				return nil
			}
			content, err := ReadText(absSourceLocation)
			if err != nil {
				return err
			}
			newContent := str.ReplaceByMap(content, replacementMap)
			if newContent == content {
				_, err = Copy(absSourceLocation, absDestinationLocation)
				return err
			}
			return WriteText(absDestinationLocation, newContent, fileMode)
		},
	)
}