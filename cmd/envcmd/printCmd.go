package envcmd

import (
	"fmt"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/dsl"
	"github.com/state-alchemists/zaruba/output"
)

var printLong = `
Print a jsonMap as environment variable declaration (key=value)

You can cascade the environment variable using --prefix flag.
This is useful if you have multiple environments (e.g., dev, staging, prod)
`

var printExample = `
Print a jsonMap as environment variable declaration
Get current environment variables as jsonMap.
    > zaruba env print '{"SERVER": "localhost", "PORT": "3306"}'
    SERVER="localhost"
    PORT="3306"

Using --prefix flag to cascade the environment.
    > zaruba env print '{"SERVER": "localhost", "PORT": "3306", "STG_SERVER": "stg.stalchmst.com", "PROD_SERVER": "stalchmst.com"}' --prefix=STG
    SERVER="stg.stalchmst.com"
    PORT="3306"
    STG_SERVER="stg.stalchmst.com"
    PROD_SERVER="stalchmst.com"
`
var printPrefix *string
var printCmd = &cobra.Command{
	Use:     "print <jsonMap> [fileName]",
	Short:   "Print a jsonMap as environment variable declarations",
	Long:    printLong,
	Example: printExample,
	Aliases: []string{"write"},
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 1)
		mapString := args[0]
		fileName := ""
		if len(args) > 1 {
			fileName = args[1]
		}
		var err error
		util := dsl.NewDSLUtil()
		if *printPrefix != "" {
			mapString, err = util.Json.Map.CascadePrefixKeys(mapString, *printPrefix)
			if err != nil {
				cmdHelper.Exit(cmd, args, logger, decoration, err)
			}
		}
		envString, err := util.Json.Map.ToEnvString(mapString)
		if err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
		if fileName != "" {
			if err := util.File.WriteText(fileName, envString, 0755); err != nil {
				cmdHelper.Exit(cmd, args, logger, decoration, err)
			}
		}
		fmt.Println(envString)
	},
}
