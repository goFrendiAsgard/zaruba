package envcmd

import "github.com/spf13/cobra"

var Cmd = &cobra.Command{
	Use:   "env",
	Short: "Environment manipulation utilities",
}

func Init() {
	getPrefix = getCmd.Flags().StringP("prefix", "p", "", "environment prefix")
	Cmd.AddCommand(getCmd)
	readPrefix = readCmd.Flags().StringP("prefix", "p", "", "environment prefix")
	Cmd.AddCommand(readCmd)
	printPrefix = printCmd.Flags().StringP("prefix", "p", "", "environment prefix")
	Cmd.AddCommand(printCmd)
}
