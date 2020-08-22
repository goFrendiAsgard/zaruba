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

	// app component definitions
	logger := log.New(os.Stdout, "", log.LstdFlags)
	config := config.CreateConfig()
	logger.Println("CONFIG:", config.ToString())
	router := gin.Default()
	defaultRmqConnection, err := amqp.Dial(config.DefaultRmqConnectionString)
	if err != nil {
		logger.Fatal("[RmqConnection]", err)
		return
	}
	rpcServer := transport.CreateRmqRPCServer(logger, defaultRmqConnection)
	rpcClient := transport.CreateRmqRPCClient(logger, defaultRmqConnection)
	subscriber := transport.CreateRmqSubscriber(logger, defaultRmqConnection)
	publisher := transport.CreateRmqPublisher(logger, defaultRmqConnection)

	// app creation
	app := core.CreateMainApp(core.MainAppConfig{
		Logger:      logger,
		Router:      router,
		Subscribers: []transport.Subscriber{subscriber},
		RPCServers:  []transport.RPCServer{rpcServer},
		HTTPPort:    config.HTTPPort,
	})

	// app setup
	app.Setup([]core.Comp{
		defaultcomponent.CreateComponent(config, router), // setup landingPage
		monitoring.CreateComponent(config, app, router),  // setup monitoring
		example.CreateComponent(
			config, router, publisher, subscriber, rpcServer, rpcClient,
		), // setup example
	})

	// app execution
	app.Run()

}
