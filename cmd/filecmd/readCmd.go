package filecmd

import (
	"fmt"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/dsl"
	"github.com/state-alchemists/zaruba/output"
)

var readExample = `
> cat myFile.txt
a
b
c

> zaruba file read myFile.txt
a
b
c
`

var readCmd = &cobra.Command{
	Use:     "read <strFileName>",
	Short:   "Read text from file",
	Example: readExample,
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 1)
		strFileName := args[0]
		util := dsl.NewDSLUtil()
		content, err := util.File.ReadText(strFileName)
		if err != nil {
			cmdHelper.Exit(cmd, logger, decoration, err)
		}
		fmt.Println(content)
	},
}
