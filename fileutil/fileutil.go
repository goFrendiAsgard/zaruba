package fileutil

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
	"github.com/state-alchemists/zaruba/jsonutil"
	"github.com/state-alchemists/zaruba/strutil"
	"github.com/state-alchemists/zaruba/yamlstyler"
	"gopkg.in/yaml.v3"
)

type FileUtil struct {
	json *jsonutil.JsonUtil
}

func NewFileUtil(jsonUtil *jsonutil.JsonUtil) *FileUtil {
	return &FileUtil{
		json: jsonUtil,
	}
}

func (fileUtil *FileUtil) IsExist(filePath string) (exist bool, err error) {
	if _, statErr := os.Stat(filePath); statErr == nil {
		return true, nil
	} else if os.IsNotExist(statErr) {
		return false, nil
	} else {
		return false, statErr
	}
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

func (fileUtil *FileUtil) ReadLines(fileName string) (jsonString string, err error) {
	content, err := fileUtil.ReadText(fileName)
	if err != nil {
		return "[]", err
	}
	return fileUtil.json.FromStringList(strings.Split(content, "\n"))
}

func (fileUtil *FileUtil) WriteLines(fileName string, jsonString string, fileMode os.FileMode) (err error) {
	lines, err := fileUtil.json.ToStringList(jsonString)
	if err != nil {
		return err
	}
	content := strings.Join(lines, "\n")
	return fileUtil.WriteText(fileName, content, fileMode)
}

func (fileUtil *FileUtil) ReadEnv(fileName string) (jsonString string, err error) {
	envMap, err := godotenv.Read(fileName)
	if err != nil {
		return "", err
	}
	return fileUtil.json.FromStringDict(envMap)
}

func (fileUtil *FileUtil) ReadYaml(fileName string) (jsonString string, err error) {
	yamlString, err := fileUtil.ReadText(fileName)
	if err != nil {
		return "", err
	}
	return fileUtil.json.FromYaml(yamlString)
}

func (fileUtil *FileUtil) WriteYaml(fileName, jsonString string, fileMode os.FileMode) (err error) {
	yamlString, err := fileUtil.json.ToYaml(jsonString)
	if err != nil {
		return err
	}
	return fileUtil.WriteText(fileName, yamlString, fileMode)
}

func (fileUtil *FileUtil) ReadYamlNode(fileName string) (node *yaml.Node, err error) {
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

func (fileUtil *FileUtil) WriteYamlNode(fileName string, node *yaml.Node, fileMode os.FileMode, yamlStylers []yamlstyler.YamlStyler) (err error) {
	yamlContentB, err := yaml.Marshal(node)
	if err != nil {
		return err
	}
	yamlContent := string(yamlContentB)
	yamlLines := strings.Split(yamlContent, "\n")
	for _, styler := range yamlStylers {
		yamlLines = styler(yamlLines)
	}
	jsonString, err := fileUtil.json.FromStringList(yamlLines)
	if err != nil {
		return err
	}
	return fileUtil.WriteLines(fileName, jsonString, fileMode)
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

func (fileUtil *FileUtil) CopyR(sourceTemplatePath, destinationPath string) (err error) {
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
			absDestinationLocation := filepath.Join(absDestinationPath, relativeLocation)
			if info.IsDir() {
				fileMode := info.Mode()
				os.Mkdir(absDestinationLocation, fileMode)
				return nil
			}
			_, err = fileUtil.Copy(absSourceLocation, absDestinationLocation)
			return err
		},
	)
}

func (fileUtil *FileUtil) Generate(sourceTemplatePath, destinationPath string, replacementMapString string) (err error) {
	replacementMap, err := fileUtil.json.ToStringDict(replacementMapString)
	if err != nil {
		return err
	}
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
			absDestinationLocation := filepath.Join(absDestinationPath, strutil.StrReplace(relativeLocation, replacementMap))
			fileMode := info.Mode()
			if info.IsDir() {
				os.Mkdir(absDestinationLocation, fileMode)
				return nil
			}
			content, err := fileUtil.ReadText(absSourceLocation)
			if err != nil {
				return err
			}
			newContent := strutil.StrReplace(content, replacementMap)
			if newContent == content {
				_, err = fileUtil.Copy(absSourceLocation, absDestinationLocation)
				return err
			}
			return fileUtil.WriteText(absDestinationLocation, newContent, fileMode)
		},
	)
}
