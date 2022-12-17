package envcmd

import (
	"fmt"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/dsl"
	"github.com/state-alchemists/zaruba/output"
)

var getLong = `
Get current environment variables as jsonMap.

You can cascade the environment variable using --prefix flag.
This is useful if you have multiple environments (e.g., dev, staging, prod)
`

var getExample = `
> export SERVER=localhost
> export PORT=3306

> zaruba env get
{"SERVER": "localhost", "PORT": "3306"}

> export SERVER=localhost
> export STG_SERVER=stg.stalchmst.com
> export PROD_SERVER=stalchmst.com
> export PORT=3306

> zaruba env get --prefix=STG
{"SERVER": "stg.stalchmst.com", "PORT": "3306", "STG_SERVER": "stg.stalchmst.com", "PROD_SERVER": "stalchmst.com"}

> zaruba env get --prefix=PROD
{"SERVER": "stalchmst.com", "PORT": "3306", "STG_SERVER": "stg.stalchmst.com", "PROD_SERVER": "stalchmst.com"}

> zaruba env get --prefix=DEV
{"SERVER": "localhost", "PORT": "3306", "STG_SERVER": "stg.stalchmst.com", "PROD_SERVER": "stalchmst.com"}
`

var getPrefix *string
var getCmd = &cobra.Command{
	Use:     "get",
	Short:   "Get current environment variables as jsonMap",
	Long:    getLong,
	Example: getExample,
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 0)
		util := dsl.NewDSLUtil()
		jsonMapEnv, err := util.Json.Map.GetFromEnv()
		if err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
		if *getPrefix == "" {
			fmt.Println(jsonMapEnv)
			return
		}
		cascadedEnvMapStr, err := util.Json.Map.CascadePrefixKeys(jsonMapEnv, *getPrefix)
		if err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
		fmt.Println(cascadedEnvMapStr)
	},
}
