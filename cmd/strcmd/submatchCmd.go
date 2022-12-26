package strcmd

import (
	"fmt"

	"github.com/state-alchemists/zaruba/dsl"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/output"
)

var submatchLong = `
Return submatch of a string matching the pattern.

Example:

String   : 'abcddefghi'
Pattern  : 'abc(d+)(.*)'

Submatch : [
"abcdefghi",   # match the whole pattern
"dd",          # match 'd+'
"efghi"        # match '.*'
]
`

var submatchExample = `
> zaruba str submatch 'abcdefghi' 'abc(d+)(.*)'
["abcdefghi","dd","efghi"]
`

var submatchCmd = &cobra.Command{
	Use:     "submatch <string> <pattern>",
	Short:   "Return submatch matching the pattern",
	Long:    submatchLong,
	Example: submatchExample,
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 2)
		text, pattern := args[0], args[1]
		util := dsl.NewDSLUtil()
		submatch, err := util.Str.Submatch(text, pattern)
		if err != nil {
			cmdHelper.Exit(cmd, logger, decoration, err)
		}
		jsonSubmatch, err := util.Json.FromStringList(submatch)
		if err != nil {
			cmdHelper.Exit(cmd, logger, decoration, err)
		}
		fmt.Println(jsonSubmatch)
	},
}
