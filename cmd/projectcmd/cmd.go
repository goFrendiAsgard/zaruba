package projectcmd

import "github.com/spf13/cobra"

var Cmd = &cobra.Command{
	Use:   "project",
	Short: "Project manipulation utilities",
}

func Init() {
	Cmd.AddCommand(addTaskIfNotExistCmd)
	Cmd.AddCommand(includeCmd)
	Cmd.AddCommand(setValueCmd)
	Cmd.AddCommand(showLogCmd)
	Cmd.AddCommand(syncEnvCmd)
	Cmd.AddCommand(syncEnvFilesCmd)
}
