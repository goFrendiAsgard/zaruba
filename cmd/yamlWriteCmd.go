package cmd

import (
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/core"
	"github.com/state-alchemists/zaruba/output"
	"gopkg.in/yaml.v3"
)

var yamlWriteCmd = &cobra.Command{
	Use:   "write <fileName> <obj>",
	Short: "Write obj to file as YAML",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(cmd, logger, decoration, args, 2)
		fileName := args[0]
		util := core.NewUtil()
		var interfaceContent interface{}
		if err := json.Unmarshal([]byte(args[1]), &interfaceContent); err != nil {
			exit(cmd, logger, decoration, err)
		}
		resultB, err := yaml.Marshal(interfaceContent)
		if err != nil {
			exit(cmd, logger, decoration, err)
		}
		if err := util.File.WriteText(fileName, string(resultB), 0755); err != nil {
			exit(cmd, logger, decoration, err)
		}
	},
}
