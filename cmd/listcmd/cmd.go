package listcmd

import "github.com/spf13/cobra"

var Cmd = &cobra.Command{
	Use:   "list",
	Short: "List manipulation utilities",
}

func Init() {
	Cmd.AddCommand(appendCmd)
	Cmd.AddCommand(containCmd)
	Cmd.AddCommand(getCmd)
	Cmd.AddCommand(joinCmd)
	Cmd.AddCommand(lengthCmd)
	Cmd.AddCommand(mergeCmd)
	Cmd.AddCommand(rangeIndexCmd)
	Cmd.AddCommand(setCmd)
	Cmd.AddCommand(validateCmd)
}
