package yamlcmd

import (
	"fmt"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/dsl"
	"github.com/state-alchemists/zaruba/output"
)

var printExample = `
> zaruba yaml print '{"id": 1, "title": "Doraemon"}'
id: 1
title: Doraemon

> zaruba yaml print '{"id": 1, "title": "Doraemon"}' book.yaml
> cat book.yaml
id: 1
title: Doraemon

`
var printCmd = &cobra.Command{
	Use:     "print <jsonAny> [yamlFileName]",
	Short:   "Print JSON map or list as YAML",
	Example: printExample,
	Aliases: []string{"write"},
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 1)
		jsonAny := args[0]
		util := dsl.NewDSLUtil()
		if len(args) > 1 {
			yamlFileName := args[1]
			if err := util.File.WriteYaml(yamlFileName, jsonAny, 0755); err != nil {
				cmdHelper.Exit(cmd, args, logger, decoration, err)
			}
			return
		}
		yamlString, err := util.Json.ToYaml(jsonAny)
		if err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
		fmt.Println(yamlString)
	},
}
