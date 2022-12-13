package jsoncmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/dsl"
	"github.com/state-alchemists/zaruba/output"
)

var printPretty *bool

var printCmd = &cobra.Command{
	Use:     "print <mapOrList> [jsonFileName]",
	Short:   "Print JSON map or list",
	Aliases: []string{"write"},
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 1)
		jsonRawString := args[0]
		util := dsl.NewDSLUtil()
		var obj interface{}
		err := json.Unmarshal([]byte(jsonRawString), &obj)
		if err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
		jsonString, err := printCmdGetJsonString(obj, *printPretty)
		if err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
		if len(args) > 1 {
			jsonFileName := args[1]
			if err = util.File.WriteText(jsonFileName, jsonString, 0755); err != nil {
				cmdHelper.Exit(cmd, args, logger, decoration, err)
			}
			return
		}
		fmt.Println(jsonString)
	},
}

func printCmdGetJsonString(obj interface{}, printPretty bool) (jsonString string, err error) {
	var jsonBytes []byte
	if !printPretty {
		jsonBytes, err = json.Marshal(obj)
	} else {
		jsonBytes, err = json.MarshalIndent(obj, "", "  ")
	}
	if err != nil {
		return "", err
	}
	jsonString = string(jsonBytes)
	return jsonString, err
}
