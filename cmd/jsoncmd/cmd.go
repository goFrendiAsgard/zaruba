package jsoncmd

import "github.com/spf13/cobra"

var Cmd = &cobra.Command{
	Use:   "json",
	Short: "Json manipulation utilities",
}

func Init() {
	Cmd.AddCommand(getCmd)
	printPretty = printCmd.Flags().BoolP("pretty", "p", true, "pretty print")
	Cmd.AddCommand(printCmd)
	Cmd.AddCommand(setCmd)
}
