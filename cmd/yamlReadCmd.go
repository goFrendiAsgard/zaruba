package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/file"
	"github.com/state-alchemists/zaruba/output"
	"gopkg.in/yaml.v2"
)

var yamlReadCmd = &cobra.Command{
	Use:   "read <fileName>",
	Short: "Read yaml from file",
	Run: func(cmd *cobra.Command, args []string) {
		commandName := cmd.Name()
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(commandName, logger, decoration, args, 1)
		fileName := args[0]
		yamlScript, err := file.ReadText(fileName)
		if err != nil {
			exit(commandName, logger, decoration, err)
		}
		// get yaml
		var interfaceContent interface{}
		if err = yaml.Unmarshal([]byte(yamlScript), &interfaceContent); err != nil {
			exit(commandName, logger, decoration, err)
		}
		interfaceContent = convertYamlObj(interfaceContent)
		resultB, err := json.Marshal(interfaceContent)
		if err != nil {
			exit(commandName, logger, decoration, err)
		}
		fmt.Println(string(resultB))
	},
}
