package cmd

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

var newUUIDStringCmd = &cobra.Command{
	Use:   "newUUIDString",
	Short: "generate new UUID string",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(uuid.NewString())
	},
}

func init() {
	rootCmd.AddCommand(newUUIDStringCmd)
}
