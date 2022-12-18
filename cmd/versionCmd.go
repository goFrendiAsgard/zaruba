package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionExample = `
> zaruba version
`

var versionCmd = &cobra.Command{
	Use:     "version",
	Short:   "Show current version",
	Example: versionExample,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(ZarubaVersion)
	},
}
