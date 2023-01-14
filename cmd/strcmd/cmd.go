package strcmd

import "github.com/spf13/cobra"

var Cmd = &cobra.Command{
	Use:   "str",
	Short: "String manipulation utilities",
}

func Init() {
	Cmd.AddCommand(addPrefixCmd)
	Cmd.AddCommand(currentTimeCmd)
	Cmd.AddCommand(doubleQuoteCmd)
	Cmd.AddCommand(fullIndentCmd)
	Cmd.AddCommand(getIndentationCmd)
	Cmd.AddCommand(indentCmd)
	Cmd.AddCommand(newNameCmd)
	Cmd.AddCommand(newUuidCmd)
	Cmd.AddCommand(padLeftCmd)
	Cmd.AddCommand(padRightCmd)
	Cmd.AddCommand(repeatCmd)
	Cmd.AddCommand(replaceCmd)
	Cmd.AddCommand(singleQuoteCmd)
	Cmd.AddCommand(splitCmd)
	Cmd.AddCommand(submatchCmd)
	Cmd.AddCommand(toCamelCmd)
	Cmd.AddCommand(toKebabCmd)
	Cmd.AddCommand(toLowerCmd)
	Cmd.AddCommand(toPascalCmd)
	Cmd.AddCommand(toPluralCmd)
	Cmd.AddCommand(toSingularCmd)
	Cmd.AddCommand(toSnakeCmd)
	Cmd.AddCommand(toTitleCmd)
	Cmd.AddCommand(toUpperCmd)
	Cmd.AddCommand(toUpperSnakeCmd)
	Cmd.AddCommand(toUrlPath)
}
