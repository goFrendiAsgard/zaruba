package jsoncmd

import "github.com/spf13/cobra"

var Cmd = &cobra.Command{
	Use:   "json",
	Short: "Json manipulation utilities",
}

func Init() {
	printPretty = printCmd.Flags().BoolP("pretty", "p", true, "pretty print")
	Cmd.AddCommand(printCmd)
}
