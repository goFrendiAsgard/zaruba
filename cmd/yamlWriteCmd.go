package cmd

import (
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/file"
	"github.com/state-alchemists/zaruba/output"
	"gopkg.in/yaml.v2"
)

var yamlWriteCmd = &cobra.Command{
	Use:   "write <fileName> <obj>",
	Short: "Write obj to file as YAML",
	Run: func(cmd *cobra.Command, args []string) {
		commandName := cmd.Name()
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(commandName, logger, decoration, args, 2)
		fileName := args[0]
		var interfaceContent interface{}
		if err := json.Unmarshal([]byte(args[1]), &interfaceContent); err != nil {
			exit(commandName, logger, decoration, err)
		}
		resultB, err := yaml.Marshal(interfaceContent)
		if err != nil {
			exit(commandName, logger, decoration, err)
		}
		if err := file.WriteText(fileName, string(resultB), 0755); err != nil {
			exit(commandName, logger, decoration, err)
		}
	},
}
