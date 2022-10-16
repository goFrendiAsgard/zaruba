package taskcmd

import "github.com/spf13/cobra"

var Cmd = &cobra.Command{
	Use:   "task",
	Short: "Task manipulation utilities",
}

func Init() {
	Cmd.AddCommand(addDependenciesCmd)
	Cmd.AddCommand(addParentsCmd)
	Cmd.AddCommand(getIconCmd)
	Cmd.AddCommand(isExistCmd)
	Cmd.AddCommand(setConfigsCmd)
	Cmd.AddCommand(setEnvsCmd)
	Cmd.AddCommand(syncEnvCmd)
}
