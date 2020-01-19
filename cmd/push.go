package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/modules/pusher"
)

func init() {
	rootCmd.AddCommand(pushCmd)
}

var pushCmd = &cobra.Command{
	Use:   "push [project-dir]",
	Short: "Push from subtrees",
	Long:  `Zaruba will push to subtrees`,
	Run: func(cmd *cobra.Command, args []string) {
		// get projectDir
		projectDir := "."
		if len(args) > 0 {
			projectDir = args[0]
		}
		// invoke action
		if err := pusher.Push(projectDir); err != nil {
			log.Fatal("[ERROR] ", err)
		}
	},
}
