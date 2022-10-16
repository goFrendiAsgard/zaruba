package strcmd

import (
	"fmt"

	"github.com/state-alchemists/zaruba/dsl"

	"github.com/spf13/cobra"
)

var newUuidCmd = &cobra.Command{
	Use:   "newUuid",
	Short: "Generate new UUID string",
	Run: func(cmd *cobra.Command, args []string) {
		util := dsl.NewDSLUtil()
		fmt.Println(util.Str.NewUUID())
	},
}
