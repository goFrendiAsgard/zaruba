package main

import (
	"app/component/defaultcomponent"
	"app/component/example"
	"app/component/monitoring"
	"app/config"
	"app/core"
	"app/transport"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
)

func main() {

	// component definitions
	logger := log.New(os.Stdout, "", log.LstdFlags)
	config := config.CreateConfig()
	logger.Println("CONFIG:", config.ToString())
	router := gin.Default()
	defaultRmqConnection, err := amqp.Dial(config.DefaultRmqConnectionString)
	if err != nil {
		logger.Fatal("[RmqConnection]", err)
	}
	rpcServer := transport.CreateRmqRPCServer(logger, defaultRmqConnection)
	rpcClient := transport.CreateRmqRPCClient(logger, defaultRmqConnection)
	subscriber := transport.CreateRmqSubscriber(logger, defaultRmqConnection)
	publisher := transport.CreateRmqPublisher(logger, defaultRmqConnection)

	// create app
	app := core.CreateMainApp(
		logger,
		router,
		[]transport.Subscriber{subscriber},
		[]transport.RPCServer{rpcServer},
		config.HTTPPort,
	)

	// setup components
	app.Setup([]core.SetupComponent{
		defaultcomponent.CreateSetup(config, router), // setup landingPage
		monitoring.CreateSetup(config, app, router),  // setup monitoring
		example.CreateComponent(
			config, router, publisher, subscriber, rpcServer, rpcClient,
		).Setup, // setup example
	})

	// run
	app.Run()

}
