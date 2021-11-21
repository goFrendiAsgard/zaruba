package cmd

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/output"
)

var serveCmd = &cobra.Command{
	Use:   "serve [location [port]]",
	Short: "Serve static files in location at a specified port",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(cmd, logger, decoration, args, 0)
		location := "."
		if len(args) >= 1 {
			location = args[0]
		}
		absLocation, err := filepath.Abs(location)
		if err != nil {
			exit(cmd, args, logger, decoration, err)
		}
		port := 8080
		if len(args) >= 2 {
			port, err = strconv.Atoi(args[1])
			if err != nil {
				exit(cmd, args, logger, decoration, err)
			}
		}
		http.Handle("/", http.FileServer(http.Dir(absLocation)))
		logger.Printf("Serving %s on HTTP port %d\n", absLocation, port)
		logger.Printf("You can open http://localhost:%d\n", port)
		portStr := fmt.Sprintf(":%d", port)
		if err := http.ListenAndServe(portStr, nil); err != nil {
			exit(cmd, args, logger, decoration, err)
		}
	},
}
