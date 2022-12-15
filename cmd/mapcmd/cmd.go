package mapcmd

import "github.com/spf13/cobra"

var Cmd = &cobra.Command{
	Use:   "map",
	Short: "JsonMap manipulation utilities",
}

func Init() {
	Cmd.AddCommand(getCmd)
	Cmd.AddCommand(getKeysCmd)
	Cmd.AddCommand(mergeCmd)
	Cmd.AddCommand(rangeKeyCmd)
	Cmd.AddCommand(setCmd)
	Cmd.AddCommand(transformKeyCmd)
	Cmd.AddCommand(validateCmd)
	Cmd.AddCommand(toStringMapCmd)
	Cmd.AddCommand(toVariedStringMapCmd)

}
