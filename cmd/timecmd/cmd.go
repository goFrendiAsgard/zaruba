package timecmd

import "github.com/spf13/cobra"

var Cmd = &cobra.Command{
	Use:   "time",
	Short: "Time related utilities",
}

func Init() {
	Cmd.AddCommand(nowCmd)
	Cmd.AddCommand(toStringCmd)
	Cmd.AddCommand(toTimestampCmd)
}
