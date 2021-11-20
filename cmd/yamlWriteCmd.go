package cmd

import (
	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/core"
	"github.com/state-alchemists/zaruba/output"
)

var yamlWriteCmd = &cobra.Command{
	Use:   "write <fileName> <obj>",
	Short: "Write obj to file as YAML",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(cmd, logger, decoration, args, 2)
		fileName := args[0]
		jsonString := args[1]
		util := core.NewCoreUtil()
		if err := util.File.WriteYaml(fileName, jsonString, 0755); err != nil {
			exit(cmd, logger, decoration, err)
		}
	},
}
