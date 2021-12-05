package cmd

import (
	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/core"
	"github.com/state-alchemists/zaruba/output"
)

var generateCmd = &cobra.Command{
	Use:   "generate <templateLocation> <destination> <replacementMap>",
	Short: "Make something based on template",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 3)
		sourceTemplatePath, destinationPath, replacementMapString := args[0], args[1], args[2]
		util := core.NewCoreUtil()
		if err := util.File.Generate(sourceTemplatePath, destinationPath, replacementMapString); err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
	},
}
