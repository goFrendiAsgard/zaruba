package timecmd

import "github.com/spf13/cobra"

var Cmd = &cobra.Command{
	Use:   "time",
	Short: "Base32 manipulation utilities",
}

func Init() {
	Cmd.AddCommand(toStringCmd)
	Cmd.AddCommand(toTimestamp)
}
