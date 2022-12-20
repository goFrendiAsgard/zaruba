package fileutil

import (
	"io"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
	"github.com/state-alchemists/zaruba/jsonutil"
	jsonHelper "github.com/state-alchemists/zaruba/jsonutil/helper"
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

func (fileUtil *FileUtil) CopyFile(src, dst string) (byteCount int64, err error) {
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

func (fileUtil *FileUtil) ReadStringList(fileName string) (stringList []string, err error) {
	stringList = []string{}
	content, err := fileUtil.ReadText(fileName)
	if err != nil {
		return stringList, err
	}
	return strings.Split(content, "\n"), nil
}

func (fileUtil *FileUtil) ReadLines(fileName string) (jsonStrList string, err error) {
	stringList, err := fileUtil.ReadStringList(fileName)
	if err != nil {
		return "[]", err
	}
	return fileUtil.json.FromStringList(stringList)
}

func (fileUtil *FileUtil) WriteStringList(fileName string, stringList []string, fileMode os.FileMode) (err error) {
	content := strings.Join(stringList, "\n")
	return fileUtil.WriteText(fileName, content, fileMode)
}

func (fileUtil *FileUtil) WriteLines(fileName string, jsonStrList string, fileMode os.FileMode) (err error) {
	lines, err := fileUtil.json.ToStringList(jsonStrList)
	if err != nil {
		return err
	}
	return fileUtil.WriteStringList(fileName, lines, fileMode)
}

func (fileUtil *FileUtil) ReadEnv(fileName string) (jsonMap string, err error) {
	envMap, err := godotenv.Read(fileName)
	if err != nil {
		return "", err
	}
	return fileUtil.json.FromStringDict(envMap)
}

func (fileUtil *FileUtil) ReadYaml(fileName string) (jsonAny string, err error) {
	yamlString, err := fileUtil.ReadText(fileName)
	if err != nil {
		return "", err
	}
	return fileUtil.json.FromYaml(yamlString)
}

func (fileUtil *FileUtil) WriteYaml(fileName, jsonAny string, fileMode os.FileMode) (err error) {
	yamlString, err := fileUtil.json.ToYaml(jsonAny)
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

func (fileUtil *FileUtil) List(dirPath string) (fileNames []string, err error) {
	fileNames = []string{}
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return fileNames, err
	}
	for _, fileInfo := range files {
		fileNames = append(fileNames, fileInfo.Name())
	}
	return fileNames, nil
}

func (fileUtil *FileUtil) Copy(sourcePath, destinationPath string) (err error) {
	absSourcePath, err := filepath.Abs(sourcePath)
	if err != nil {
		return err
	}
	absDestinationPath, err := filepath.Abs(destinationPath)
	if err != nil {
		return err
	}
	return filepath.Walk(absSourcePath,
		func(absSourceLocation string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			relativeLocation := absSourceLocation[len(absSourcePath):]
			absDestinationLocation := filepath.Join(absDestinationPath, relativeLocation)
			if info.IsDir() {
				fileMode := info.Mode()
				os.Mkdir(absDestinationLocation, fileMode)
				return nil
			}
			_, err = fileUtil.CopyFile(absSourceLocation, absDestinationLocation)
			return err
		},
	)
}

func (fileUtil *FileUtil) Walk(dirPath string) (relativeChildPaths []string, err error) {
	absDirPath, err := filepath.Abs(dirPath)
	if err != nil {
		return relativeChildPaths, err
	}
	relativeChildPaths = []string{}
	err = filepath.Walk(absDirPath,
		func(absLocation string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			relativeLocation := absLocation[len(absDirPath):]
			relativeChildPaths = append(relativeChildPaths, relativeLocation)
			return nil
		},
	)
	return relativeChildPaths, err
}

func (fileUtil *FileUtil) Generate(sourceTemplatePath, destinationPath string, jsonMapReplacement string) (err error) {
	replacementMap, absSourceTemplatePath, absDestinationPath, err := fileUtil.preparePathAndReplacementMap(sourceTemplatePath, destinationPath, jsonMapReplacement)
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
				_, err = fileUtil.CopyFile(absSourceLocation, absDestinationLocation)
				return err
			}
			return fileUtil.WriteText(absDestinationLocation, newContent, fileMode)
		},
	)
}

func (fileUtil *FileUtil) GetLinesSubmatch(filePath, jsonStrListPattern string, desiredPatternIndex int) (matchIndex int, jsonSubmatch string, err error) {
	absDestinationPath, err := filepath.Abs(filePath)
	if err != nil {
		return -1, "[]", err
	}
	stringList, err := fileUtil.ReadStringList(absDestinationPath)
	if err != nil {
		return -1, "[]", err
	}
	patterns, patternErr := jsonHelper.ToStringList(jsonStrListPattern)
	if patternErr != nil {
		patterns = []string{jsonStrListPattern}
	}
	matchIndex, submatch, err := strutil.StrGetLineSubmatch(stringList, patterns, desiredPatternIndex)
	if err != nil {
		return -1, "[]", err
	}
	jsonSubmatch, err = jsonHelper.FromStringList(submatch)
	return matchIndex, jsonSubmatch, err
}

func (fileUtil *FileUtil) preparePathAndReplacementMap(sourceTemplatePath, destinationPath string, replacementMapString string) (replacementMap jsonHelper.StringDict, absSourceTemplatePath, absDestinationPath string, err error) {
	replacementMap, err = fileUtil.json.ToStringDict(replacementMapString)
	if err != nil {
		return replacementMap, absSourceTemplatePath, absDestinationPath, err
	}
	absSourceTemplatePath, err = filepath.Abs(sourceTemplatePath)
	if err != nil {
		return replacementMap, absSourceTemplatePath, absDestinationPath, err
	}
	absDestinationPath, err = filepath.Abs(destinationPath)
	return replacementMap, absSourceTemplatePath, absDestinationPath, err
}

func (fileUtil *FileUtil) ReplaceLineAtIndex(destinationFilePath string, content string, index int) (err error) {
	return fileUtil.replaceLineAtIndex(destinationFilePath, content, index, "REPLACE")
}

func (fileUtil *FileUtil) InsertLineAfterIndex(destinationFilePath string, content string, index int) (err error) {
	return fileUtil.replaceLineAtIndex(destinationFilePath, content, index, "AFTER")
}

func (fileUtil *FileUtil) InsertLineBeforeIndex(destinationFilePath string, content string, index int) (err error) {
	return fileUtil.replaceLineAtIndex(destinationFilePath, content, index, "BEFORE")
}

func (fileUtil *FileUtil) replaceLineAtIndex(destinationFilePath string, content string, index int, mode string) (err error) {
	absDestinationPath, err := filepath.Abs(destinationFilePath)
	if err != nil {
		return err
	}
	stringList, err := fileUtil.ReadStringList(absDestinationPath)
	if err != nil {
		return err
	}
	destinationMode, err := fileUtil.getFileMode(destinationFilePath)
	if err != nil {
		return err
	}
	if index < 0 {
		index = len(stringList) + index
	}
	replacements := []string{content}
	switch mode {
	case "BEFORE":
		replacements = append(replacements, stringList[index])
	case "AFTER":
		replacements = append([]string{stringList[index]}, replacements...)
	case "REPLACE":
	default:
	}
	newStringList, err := strutil.StrReplaceLineAtIndex(stringList, index, replacements)
	if err != nil {
		return err
	}
	return fileUtil.WriteStringList(absDestinationPath, newStringList, destinationMode)
}

func (fileUtil *FileUtil) getFileMode(filePath string) (fileMode fs.FileMode, err error) {
	fileStat, err := os.Stat(filePath)
	if err != nil {
		return fileMode, err
	}
	fileMode = fileStat.Mode()
	return fileMode, err
}
