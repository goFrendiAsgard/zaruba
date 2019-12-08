package cmd

import (
	"log"
	"strings"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/installtemplate"
)

func init() {
	rootCmd.AddCommand(installTemplateCmd)
}

var installTemplateCmd = &cobra.Command{
	Use:   "install-template <git-url> [template-dir]",
	Short: "Install template",
	Long:  `Zaruba will install a template for you`,
	Run: func(cmd *cobra.Command, args []string) {
		// handle invalid parameter
		if len(args) < 1 {
			log.Fatal("[ERROR] template's Git URL is expected")
		}
		// get `gitURL` and `templateDir`
		gitURL := args[0]
		templateDir := ""
		if len(args) < 2 {
			urlParts := strings.Split(gitURL, "/")
			templateDir = urlParts[len(urlParts)-1]
		} else {
			templateDir = args[1]
		}
		// invoke action
		log.Printf("[INFO] Invoking install-template. git-url: %s, template-dir: %s", gitURL, templateDir)
		if err := installtemplate.Install(gitURL, templateDir); err != nil {
			log.Fatal("[ERROR]", err)
		}
	},
}
