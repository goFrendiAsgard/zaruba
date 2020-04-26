package main

import (
	"app/component/example"
	"app/component/landingpage"
	"app/component/monitoring"
	"app/config"
	"app/core"
	"fmt"
)

func main() {
	// create config and app
	config := config.CreateConfig()
	fmt.Println("CONFIG:", config.ToString())
	app := core.CreateMainApp(
		config.HTTPPort,
		config.GlobalRmqConnectionString,
		config.LocalRmqConnectionString,
	)
	// setup components
	app.Setup([]core.SetupComponent{
		landingpage.CreateSetup(app, config),       // setup landingPage
		monitoring.CreateSetup(app, config),        // setup monitoring
		example.CreateComponent(app, config).Setup, // setup example
	})
	// run
	app.Run()
}
