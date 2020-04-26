package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var banner string = `
                     _           
 ______ _ _ __ _   _| |__   __ _ 
|_  / _  | '__| | | | '_ \ / _  |
 / / (_| | |  | |_| | |_) | (_| |
/___\__,_|_|   \__,_|_.__/ \__,_|
`

var rootCmd = &cobra.Command{
	Use:   "zaruba <action> [...args]",
	Short: "Zaruba is agnostic generator and task runner",
	Long:  `Zaruba will help you create project and maintain dependencies among components`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println(banner)
			fmt.Println("My name is Zaruba. I came to be when Garo came to be and I am forever with him.")
			cmd.Help()
		}
	},
}

// Execute basic action
func Execute() {
	rootCmd.Execute()
}
