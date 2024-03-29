package filecmd

import (
	"github.com/spf13/cobra"

	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/dsl"
	"github.com/state-alchemists/zaruba/output"
)

var copyExample = `
> ls myDir
main.py    module

> zaruba file copy myDir otherDir

> ls otherDir
main.py    module
`

var copyCmd = &cobra.Command{
	Use:     "copy <strSourcePath> <strDestinationPath>",
	Short:   "Copy files/folders recursively",
	Example: copyExample,
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 2)
		strSourcePath, strDestinationPath := args[0], args[1]
		util := dsl.NewDSLUtil()
		if err := util.File.Copy(strSourcePath, strDestinationPath); err != nil {
			cmdHelper.Exit(cmd, logger, decoration, err)
		}
	},
}
