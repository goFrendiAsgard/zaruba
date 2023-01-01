package numcmd

import "github.com/spf13/cobra"

var Cmd = &cobra.Command{
	Use:   "num",
	Short: "Number manipulation utilities. Use '--' for negative number arguments",
}

func Init() {
	Cmd.AddCommand(addCmd)
	Cmd.AddCommand(subtractCmd)
	Cmd.AddCommand(multiplyCmd)
	Cmd.AddCommand(divideCmd)
	Cmd.AddCommand(randomCmd)
	Cmd.AddCommand(rangeCmd)
	Cmd.AddCommand(validateIntCmd)
	Cmd.AddCommand(validateFloatCmd)
}
