package taskcmd

import "github.com/spf13/cobra"

var Cmd = &cobra.Command{
	Use:   "task",
	Short: "Task manipulation utilities",
}

func Init() {
	Cmd.AddCommand(addDependencyCmd)
	Cmd.AddCommand(addParentCmd)
	Cmd.AddCommand(getIconCmd)
	Cmd.AddCommand(isExistCmd)
	Cmd.AddCommand(setConfigCmd)
	Cmd.AddCommand(setEnvCmd)
	Cmd.AddCommand(syncEnvCmd)
}
