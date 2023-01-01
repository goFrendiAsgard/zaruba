package base32cmd

import "github.com/spf13/cobra"

var Cmd = &cobra.Command{
	Use:   "base32",
	Short: "Base32 manipulation utilities",
}

func Init() {
	Cmd.AddCommand(decodeCmd)
	Cmd.AddCommand(encodeCmd)
}
