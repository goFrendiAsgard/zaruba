package previousval

import (
	"os"

	"github.com/state-alchemists/zaruba/dsl"
	yaml "gopkg.in/yaml.v3"
)

// Load load previous value
func Load(project *dsl.Project, fileName string) (err error) {
	if _, statErr := os.Stat(fileName); os.IsNotExist(statErr) {
		return nil
	}
	return project.AddValue(fileName)
}

// Save save non-secret input into file
func Save(project *dsl.Project, fileName string) (err error) {
	values := map[string]string{}
	projectValues := project.GetValues()
	for key, val := range projectValues {
		input, exists := project.Inputs[key]
		if !exists || input.Secret {
			continue
		}
		values[key] = val
	}
	b, err := yaml.Marshal(values)
	if err != nil {
		return err
	}
	f, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(b)
	return err
}
