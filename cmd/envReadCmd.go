package cmd

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/core"
	"github.com/state-alchemists/zaruba/output"
)

var envReadCmd = &cobra.Command{
	Use:   "read <fileName> [prefix]",
	Short: "Read envmap from file",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(cmd, logger, decoration, args, 1)
		fileName := args[0]
		envMap, err := godotenv.Read(fileName)
		if err != nil {
			exit(cmd, logger, decoration, err)
		}
		util := core.NewCoreUtil()
		mapString := util.Json.FromInterface(envMap)
		if len(args) > 1 {
			prefix := args[1]
			mapString, err = util.Json.Map.CascadePrefixKeys(mapString, prefix)
			if err != nil {
				exit(cmd, logger, decoration, err)
			}
		}
		fmt.Println(mapString)
	},
}
