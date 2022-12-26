package mapcmd

import (
	"fmt"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/dsl"
	"github.com/state-alchemists/zaruba/output"
)

var rangeKeyExample = `
> zaruba map rangeKey '{"server": "localhost", "port": 3306}'
server
port

> MAP={"server": "localhost", "port": 3306}
> for KEY in $(zaruba map rangeKey "$MAP")
  do
	VALUE=$(zaruba map get "$MAP" $KEY)
	echo "$KEY $VALUE"
  done

server localhost
port 3306
`

var rangeKeyCmd = &cobra.Command{
	Use:     "rangeKey <jsonMap>",
	Short:   "Print jsonMap keys",
	Example: rangeKeyExample,
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 1)
		jsonMap := args[0]
		util := dsl.NewDSLUtil()
		keys, err := util.Json.Map.GetKeys(jsonMap)
		if err != nil {
			cmdHelper.Exit(cmd, logger, decoration, err)
		}
		for _, key := range keys {
			fmt.Println(key)
		}
	},
}
