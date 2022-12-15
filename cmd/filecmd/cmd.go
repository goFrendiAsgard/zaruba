package filecmd

import "github.com/spf13/cobra"

var Cmd = &cobra.Command{
	Use:   "file",
	Short: "File manipulation utilities",
}

func Init() {
	Cmd.AddCommand(copyCmd)
	Cmd.AddCommand(getLineCmd)
	getLineIndexDesiredPatternIndex = getLineIndexCmd.Flags().IntP("index", "i", -1, "desired pattern index")
	Cmd.AddCommand(getLineIndexCmd)
	insertAfterIndex = insertAfterCmd.Flags().IntP("index", "i", -1, "index")
	Cmd.AddCommand(insertAfterCmd)
	insertBeforeIndex = insertBeforeCmd.Flags().IntP("index", "i", 0, "index")
	Cmd.AddCommand(insertBeforeCmd)
	Cmd.AddCommand(listCmd)
	Cmd.AddCommand(readCmd)
	Cmd.AddCommand(walkCmd)
	replaceIndex = replaceCmd.Flags().IntP("index", "i", 0, "index")
	Cmd.AddCommand(replaceCmd)

}
