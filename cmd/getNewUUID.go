package cmd

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

var getNewUUIDCmd = &cobra.Command{
	Use:   "getNewUUID",
	Short: "Generate new UUID string",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(uuid.NewString())
	},
}

func init() {
	rootCmd.AddCommand(getNewUUIDCmd)
}
