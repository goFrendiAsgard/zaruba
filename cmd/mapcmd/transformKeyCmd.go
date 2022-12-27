package mapcmd

import (
	"fmt"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/dsl"
	"github.com/state-alchemists/zaruba/output"
)

var transformKeyPrefix string
var transformKeySuffix string
var transformKeyTransformation []string

var transformKeyExample = `
> zaruba map transformKey '{"server": "localhost", "port": 3306}' -p=DB_ -s=_DEV -t=upperSnake
{"DB_PORT_DEV":3306,"DB_SERVER_DEV":"localhost"}
`

var transformKeyCmd = &cobra.Command{
	Use:     "transformKey <jsonMap>",
	Short:   "Transform map keys",
	Example: transformKeyExample,
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 1)
		jsonMap, prefix, suffix := args[0], transformKeyPrefix, transformKeySuffix
		transformers := []func(string) string{}
		util := dsl.NewDSLUtil()
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
				cmdHelper.Exit(cmd, logger, decoration, fmt.Errorf("invalid transformer %s", alias))
			}
		}
		newJsonMap, err := util.Json.Map.TransformKeys(jsonMap, prefix, suffix, transformers...)
		if err != nil {
			cmdHelper.Exit(cmd, logger, decoration, err)
		}
		fmt.Println(newJsonMap)
	},
}
