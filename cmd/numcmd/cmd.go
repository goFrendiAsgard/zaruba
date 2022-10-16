package numcmd

import "github.com/spf13/cobra"

var Cmd = &cobra.Command{
	Use:   "num",
	Short: "Number manipulation utilities",
}

func Init() {
	Cmd.AddCommand(rangeCmd)
	Cmd.AddCommand(validateIntCmd)
	Cmd.AddCommand(validateFloatCmd)
}
