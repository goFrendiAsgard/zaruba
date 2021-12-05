package mapcmd

import (
	"fmt"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/core"
	"github.com/state-alchemists/zaruba/output"
)

var transformKeyPrefix string
var transformKeySuffix string
var transformKeyTransformation []string

var transformKeyCmd = &cobra.Command{
	Use:   "transformKey <jsonMap>",
	Short: "Transform map keys",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 1)
		mapString, prefix, suffix := args[0], transformKeyPrefix, transformKeySuffix
		transformers := []func(string) string{}
		util := core.NewCoreUtil()
		for _, alias := range transformKeyTransformation {
			switch alias {
			case "upper":
				transformers = append(transformers, util.Str.ToUpper)
			case "lower":
				transformers = append(transformers, util.Str.ToLower)
			case "upperSnake":
				transformers = append(transformers, util.Str.ToUpperSnake)
			case "camel":
				transformers = append(transformers, util.Str.ToCamel)
			case "kebab":
				transformers = append(transformers, util.Str.ToKebab)
			case "pascal":
				transformers = append(transformers, util.Str.ToPascal)
			case "snake":
				transformers = append(transformers, util.Str.ToSnake)
			default:
				cmdHelper.Exit(cmd, args, logger, decoration, fmt.Errorf("invalid transformer %s", alias))
			}
		}
		newMapString, err := util.Json.Map.TransformKeys(mapString, prefix, suffix, transformers...)
		if err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
		fmt.Println(newMapString)
	},
}

func init() {
	transformKeyCmd.Flags().StringVarP(&transformKeyPrefix, "prefix", "p", "", "key prefix")
	transformKeyCmd.Flags().StringVarP(&transformKeySuffix, "suffix", "s", "", "key suffix")
	transformKeyCmd.Flags().StringArrayVarP(&transformKeyTransformation, "transformation", "t", []string{}, "transfomration (e.g: '-t upper', '-t lower', '-t upperSnake', -t 'camel', '-t kebab', '-t pascal', '-t snake')")

}
