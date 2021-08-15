package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/output"
	"gopkg.in/yaml.v2"
)

var yamlPrintCmd = &cobra.Command{
	Use:   "print <obj>",
	Short: "Print obj as YAML",
	Run: func(cmd *cobra.Command, args []string) {
		commandName := cmd.Name()
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(commandName, logger, decoration, args, 1)
		var interfaceContent interface{}
		if err := json.Unmarshal([]byte(args[0]), &interfaceContent); err != nil {
			exit(commandName, logger, decoration, err)
		}
		resultB, err := yaml.Marshal(interfaceContent)
		if err != nil {
			exit(commandName, logger, decoration, err)
		}
		fmt.Println(string(resultB))
	},
}
