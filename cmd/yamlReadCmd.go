package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/core"
	"github.com/state-alchemists/zaruba/output"
)

var yamlReadCmd = &cobra.Command{
	Use:   "read <fileName>",
	Short: "Read yaml from file",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(cmd, logger, decoration, args, 1)
		fileName := args[0]
		util := core.NewCoreUtil()
		jsonString, err := util.File.ReadYaml(fileName)
		if err != nil {
			exit(cmd, logger, decoration, err)
		}
		fmt.Println(jsonString)
	},
}
