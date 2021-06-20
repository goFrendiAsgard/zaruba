package util

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func SetProjectValue(fileName, key, value string) (err error) {
	if key == "" {
		return fmt.Errorf("key cannot be empty")
	}
	if value == "" {
		return fmt.Errorf("value cannot be empty")
	}
	fileContentB, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}
	configMap := map[string]string{}
	if err := yaml.Unmarshal(fileContentB, &configMap); err != nil {
		return err
	}
	configMap[key] = value
	newFileContentB, err := yaml.Marshal(configMap)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(fileName, newFileContentB, 0755)
}
