package mapcmd

import (
	"fmt"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/dsl"
	"github.com/state-alchemists/zaruba/output"
)

var toVariedStringMapCmd = &cobra.Command{
	Use:   "toVariedStringMap <jsonMap> [keys...]",
	Short: "Transform to string map with various combination (original, kebab-case, camelCase, PascalCase, snake_case, lower case, UPPER CASE, UPPER_SNAKE_CASE, \"double quoted\", 'single quoted')",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 1)
		jsonMap, keys := args[0], args[1:]
		util := dsl.NewDSLUtil()
		newJsonMap, err := util.Json.Map.ToVariedStringMap(jsonMap, keys...)
		if err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
		fmt.Println(newJsonMap)
	},
}
