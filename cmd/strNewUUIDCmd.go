package cmd

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

var strNewUUIDCmd = &cobra.Command{
	Use:   "newUUID",
	Short: "Generate new UUID string",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(uuid.NewString())
	},
}
