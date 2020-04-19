package main

import (
	"app/component/example"
	"app/component/monitoring"
	"app/config"
	"app/core"
	"fmt"
)

func main() {
	// create config and app
	config := config.CreateConfig()
	fmt.Println(config)
	app := core.CreateApplication(
		config.HTTPPort,
		config.GlobalRmqConnectionString,
		config.LocalRmqConnectionString,
	)
	// setup components
	app.Setup([]core.SetupComponent{
		monitoring.CreateSetup(app, config),        // setup monitoring
		example.CreateComponent(app, config).Setup, // setup example
	})
	// run
	app.Run()
}
