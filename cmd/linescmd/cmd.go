package linescmd

import "github.com/spf13/cobra"

var Cmd = &cobra.Command{
	Use:   "lines",
	Short: "Lines manipulation utilities",
}

func Init() {
	Cmd.AddCommand(fillCmd)
	Cmd.AddCommand(getIndexCmd)
	Cmd.AddCommand(insertAfterCmd)
	Cmd.AddCommand(insertBeforeCmd)
	Cmd.AddCommand(readCmd)
	Cmd.AddCommand(replaceCmd)
	Cmd.AddCommand(submatchCmd)
	Cmd.AddCommand(writeCmd)
}
