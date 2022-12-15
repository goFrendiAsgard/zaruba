package cmd

import (
	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/dsl"
	"github.com/state-alchemists/zaruba/output"
)

var generateCmd = &cobra.Command{
	Use:   "generate <sourceTemplatePath> <destinationPath> [jsonMapReplacement]",
	Short: "Generate a directory based on sourceTemplate and jsonMapReplacement",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 2)
		sourceTemplatePath, destinationPath := args[0], args[1]
		jsonMapReplacement := "{}"
		if len(args) > 2 {
			jsonMapReplacement = args[2]
		}
		util := dsl.NewDSLUtil()
		if err := util.File.Generate(sourceTemplatePath, destinationPath, jsonMapReplacement); err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
	},
}
