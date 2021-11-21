package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/core"
	"github.com/state-alchemists/zaruba/output"
)

var yamlPrintCmd = &cobra.Command{
	Use:   "print <obj>",
	Short: "Print obj as YAML",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(cmd, logger, decoration, args, 1)
		jsonString := args[0]
		util := core.NewCoreUtil()
		yamlString, err := util.Json.ToYaml(jsonString)
		if err != nil {
			exit(cmd, args, logger, decoration, err)
		}
		fmt.Println(yamlString)
	},
}
