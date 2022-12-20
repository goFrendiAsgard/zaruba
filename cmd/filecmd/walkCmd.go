package filecmd

import (
	"fmt"

	"github.com/spf13/cobra"

	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/dsl"
	"github.com/state-alchemists/zaruba/output"
)

var walkExample = `
> ls myDir
a.txt   b.txt   c
> ls myDir/c
d.txt   e.txt

> zaruba file walk myDir
/a.txt
/b.txt
/c
/c/d.txt
/c/e.txt
`

var walkCmd = &cobra.Command{
	Use:     "walk <strDirPath>",
	Short:   "List files/folder in a path, recursively",
	Example: walkExample,
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 1)
		strDirPath := args[0]
		util := dsl.NewDSLUtil()
		fileNames, err := util.File.Walk(strDirPath)
		if err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
		for _, fileName := range fileNames {
			fmt.Println(fileName)
		}
	},
}
