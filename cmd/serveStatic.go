package cmd

import (
	"fmt"
	"path/filepath"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/output"
)

var serveStaticCmd = &cobra.Command{
	Use:   "serveStatic <location> <port>",
	Short: "Serve static files in location at a specified port",
	Run: func(cmd *cobra.Command, args []string) {
		commandName := cmd.Name()
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(commandName, logger, decoration, args, 1)
		absLocation, err := filepath.Abs(args[0])
		if err != nil {
			exit(commandName, logger, decoration, err)
		}
		port := 8080
		if len(args) == 1 || args[1] != "" {
			port, err = strconv.Atoi(args[1])
			if err != nil {
				exit(commandName, logger, decoration, err)
			}
		}
		e := echo.New()
		e.Use(middleware.Logger())
		e.Static("/", absLocation)
		e.Start(fmt.Sprintf(":%d", port))
	},
}

func init() {
	rootCmd.AddCommand(serveStaticCmd)
}
