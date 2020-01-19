package cmd

import (
	"log"
	"strings"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/modules/template"
)

func init() {
	rootCmd.AddCommand(installTemplateCmd)
}

var installTemplateCmd = &cobra.Command{
	Use:   "install-template <git-url> [template-dir]",
	Short: "Install template",
	Long:  `Zaruba will install a template`,
	Run: func(cmd *cobra.Command, args []string) {
		// handle invalid parameter
		if len(args) < 1 {
			log.Fatal("[ERROR] template's Git URL is expected, current arguments: ", args)
		}
		// get `gitURL` and `templateDir`
		gitURL := args[0]
		templateDir := ""
		if len(args) < 2 {
			urlParts := strings.Split(gitURL, "/")
			templateDir = strings.Split(urlParts[len(urlParts)-1], ".")[0]
		} else {
			templateDir = args[1]
		}
		// invoke action
		if err := template.Install(gitURL, templateDir); err != nil {
			log.Fatal("[ERROR] ", err)
		}
	},
}
