package cmd

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/output"
)

var serveExample = `
> zaruba serve
Serving /home/user on HTTP port 8080
You can open http://localhost:8080

> zaruba serve /home/user/docs
Serving /home/user/docs on HTTP port 8080
You can open http://localhost:8080

> zaruba serve /home/user/docs 3000
Serving /home/user/docs on HTTP port 3000
You can open http://localhost:3000
`

var serveCmd = &cobra.Command{
	Use:     "serve [location [port]]",
	Short:   "Serve static website",
	Example: serveExample,
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 0)
		location := "."
		if len(args) >= 1 {
			location = args[0]
		}
		absLocation, err := filepath.Abs(location)
		if err != nil {
			cmdHelper.Exit(cmd, logger, decoration, err)
		}
		port := 8080
		if len(args) >= 2 {
			port, err = strconv.Atoi(args[1])
			if err != nil {
				cmdHelper.Exit(cmd, logger, decoration, err)
			}
		}
		http.Handle("/", http.FileServer(http.Dir(absLocation)))
		logger.Printf("Serving %s on HTTP port %d\n", absLocation, port)
		logger.Printf("You can open http://localhost:%d\n", port)
		portStr := fmt.Sprintf(":%d", port)
		if err := http.ListenAndServe(portStr, nil); err != nil {
			cmdHelper.Exit(cmd, logger, decoration, err)
		}
	},
}
