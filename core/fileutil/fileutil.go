package fileutil

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/state-alchemists/zaruba/core/strutil"
	"github.com/state-alchemists/zaruba/core/yamlStyler"
	"gopkg.in/yaml.v3"
)

type FileUtil struct{}

func NewFileUtil() *FileUtil {
	return &FileUtil{}
}

func (fileUtil *FileUtil) Copy(src, dst string) (byteCount int64, err error) {
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

func (fileUtil *FileUtil) ReadText(fileName string) (text string, err error) {
	contentB, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	return string(contentB), nil
}

func (fileUtil *FileUtil) WriteText(fileName string, text string, fileMode os.FileMode) (err error) {
	if err := os.MkdirAll(filepath.Dir(fileName), fileMode); err != nil {
		return err
	}
	return ioutil.WriteFile(fileName, []byte(text), fileMode)
}

func (fileUtil *FileUtil) ReadLines(fileName string) (lines []string, err error) {
	content, err := fileUtil.ReadText(fileName)
	if err != nil {
		return []string{}, err
	}
	return strings.Split(content, "\n"), nil
}

func (fileUtil *FileUtil) WriteLines(fileName string, lines []string, fileMode os.FileMode) (err error) {
	content := strings.Join(lines, "\n")
	return fileUtil.WriteText(fileName, content, fileMode)
}

func (fileUtil *FileUtil) ReadYaml(fileName string) (node *yaml.Node, err error) {
	var nodeObj yaml.Node
	node = &nodeObj
	yamlScript, err := fileUtil.ReadText(fileName)
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

func (fileUtil *FileUtil) WriteYaml(fileName string, node *yaml.Node, fileMode os.FileMode, yamlStyler []yamlStyler.YamlStyler) (err error) {
	yamlContentB, err := yaml.Marshal(node)
	if err != nil {
		return err
	}
	yamlContent := string(yamlContentB)
	yamlLines := strings.Split(yamlContent, "\n")
	for _, preprocessor := range yamlStyler {
		yamlLines = preprocessor(yamlLines)
	}
	return fileUtil.WriteLines(fileName, yamlLines, fileMode)
}

func (fileUtil *FileUtil) ListDir(dirPath string) (fileNames []string, err error) {
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

func (fileUtil *FileUtil) Generate(sourceTemplatePath, destinationPath string, replacementMap map[string]string) (err error) {
	absSourceTemplatePath, err := filepath.Abs(sourceTemplatePath)
	if err != nil {
		return err
	}
	absDestinationPath, err := filepath.Abs(destinationPath)
	if err != nil {
		return err
	}
	strUtil := strutil.NewStrutil()
	return filepath.Walk(absSourceTemplatePath,
		func(absSourceLocation string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			relativeLocation := absSourceLocation[len(absSourceTemplatePath):]
			absDestinationLocation := filepath.Join(absDestinationPath, strUtil.Replace(relativeLocation, replacementMap))
			fileMode := info.Mode()
			if info.IsDir() {
				os.Mkdir(absDestinationLocation, fileMode)
				return nil
			}
			content, err := fileUtil.ReadText(absSourceLocation)
			if err != nil {
				return err
			}
			newContent := strUtil.Replace(content, replacementMap)
			if newContent == content {
				_, err = fileUtil.Copy(absSourceLocation, absDestinationLocation)
				return err
			}
			return fileUtil.WriteText(absDestinationLocation, newContent, fileMode)
		},
	)
}
