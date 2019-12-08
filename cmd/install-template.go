package cmd

import (
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/template"
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
			templateDir = urlParts[len(urlParts)-1]
		} else {
			templateDir = args[1]
		}
		// invoke action
		cwd, _ := os.Getwd()
		log.Printf("[INFO] Invoking install-template. cwd: %s, git-url: %s, template-dir: %s", cwd, gitURL, templateDir)
		if err := template.Install(gitURL, templateDir); err != nil {
			log.Fatal("[ERROR] ", err)
		}
	},
}
