package linescmd

import "github.com/spf13/cobra"

var Cmd = &cobra.Command{
	Use:   "lines",
	Short: "Lines manipulation utilities",
}

func Init() {
	getIndexDesiredPatternIndex = getIndexCmd.Flags().IntP("index", "i", -1, "desired pattern index")
	Cmd.AddCommand(getIndexCmd)
	insertAfterIndex = insertAfterCmd.Flags().IntP("index", "i", -1, "desired pattern index")
	Cmd.AddCommand(insertAfterCmd)
	insertBeforeIndex = insertBeforeCmd.Flags().IntP("index", "i", 0, "desired pattern index")
	Cmd.AddCommand(insertBeforeCmd)
	Cmd.AddCommand(readCmd)
	replaceIndex = replaceCmd.Flags().IntP("index", "i", 0, "desired pattern index")
	Cmd.AddCommand(replaceCmd)
	submatchDesiredPatternIndex = submatchCmd.Flags().IntP("index", "i", -1, "desired pattern index")
	Cmd.AddCommand(submatchCmd)
	Cmd.AddCommand(printCmd)
}
