package mapcmd

import "github.com/spf13/cobra"

var Cmd = &cobra.Command{
	Use:   "map",
	Short: "JsonMap manipulation utilities",
}

func Init() {
	Cmd.AddCommand(getCmd)
	Cmd.AddCommand(getKeysCmd)
	Cmd.AddCommand(mergeCmd)
	Cmd.AddCommand(rangeKeyCmd)
	Cmd.AddCommand(setCmd)
	transformKeyCmd.Flags().StringVarP(&transformKeyPrefix, "prefix", "p", "", "key prefix")
	transformKeyCmd.Flags().StringVarP(&transformKeySuffix, "suffix", "s", "", "key suffix")
	transformKeyCmd.Flags().StringArrayVarP(&transformKeyTransformation, "transformation", "t", []string{}, "transformation (e.g., '-t upper', '-t lower', '-t upperSnake', -t 'camel', '-t kebab', '-t pascal', '-t snake')")
	Cmd.AddCommand(transformKeyCmd)
	Cmd.AddCommand(validateCmd)
	Cmd.AddCommand(toStringMapCmd)
	Cmd.AddCommand(toVariedStringMapCmd)

}
