package jsoncmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/dsl"
	"github.com/state-alchemists/zaruba/output"
)

var printLong = `
Print json.
You can print the output into stdout or a file.
You can also define whether you want to "pretty print" or not (It is pretty print by default).
`

var printExample = `
> zaruba json print '{"id": 1, "title": "Doraemon"}'
{
  "id": 1,
  "title": "Doraemon"
}

> zaruba json print '{"id": 1, "title": "Doraemon"}' --pretty=false
{"id":1,"title":"Doraemon"}

> zaruba json print '{"id": 1, "title": "Doraemon"}' book.json
> cat book.json
{
  "id": 1,
  "title": "Doraemon"
}
`

var printPretty *bool
var printCmd = &cobra.Command{
	Use:     "print <jsonAny> [jsonFileName]",
	Short:   "Print json",
	Long:    printLong,
	Example: printExample,
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
