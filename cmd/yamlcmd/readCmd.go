package yamlcmd

import (
	"fmt"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/dsl"
	"github.com/state-alchemists/zaruba/output"
)

var readExample = `
> cat book.yaml
id: 1
title: Doraemon

> zaruba yaml read book.yaml
{"id":1,"title":"Doraemon"}
`

var readCmd = &cobra.Command{
	Use:     "read <yamlFileName>",
	Short:   "Read YAML file as JSON map or list",
	Example: readExample,
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 1)
		fileName := args[0]
		util := dsl.NewDSLUtil()
		jsonString, err := util.File.ReadYaml(fileName)
		if err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
		fmt.Println(jsonString)
	},
}
