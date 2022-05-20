package jsoncmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/core"
	"github.com/state-alchemists/zaruba/output"
)

var printCmd = &cobra.Command{
	Use:     "print <mapOrList> [jsonFileName]",
	Short:   "Print JSON map or list",
	Aliases: []string{"write"},
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 1)
		jsonString := args[0]
		util := core.NewCoreUtil()
		var obj interface{}
		err := json.Unmarshal([]byte(jsonString), &obj)
		if err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
		prettyJsonBytes, err := json.MarshalIndent(obj, "", "  ")
		if err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
		prettyJsonString := string(prettyJsonBytes)
		if len(args) > 1 {
			jsonFileName := args[1]
			if err = util.File.WriteText(jsonFileName, prettyJsonString, 0755); err != nil {
				cmdHelper.Exit(cmd, args, logger, decoration, err)
			}
			return
		}
		fmt.Println(prettyJsonString)
	},
}
