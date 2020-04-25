package core

import (
	"app/transport"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/gin-gonic/gin"
)

// mainApp implementation, comply with App
type mainApp struct {
	readinessMux     sync.Mutex
	readiness        bool
	livenessMux      sync.Mutex
	liveness         bool
	httpPort         int
	logger           *log.Logger
	router           *gin.Engine
	globalPublisher  transport.Publisher
	localPublisher   transport.Publisher
	globalSubscriber transport.Subscriber
	localSubscriber  transport.Subscriber
	globalRPCServer  transport.RPCServer
	localRPCServer   transport.RPCServer
	globalRPCClient  transport.RPCClient
	localRPCClient   transport.RPCClient
}

// Setup application
func (app *mainApp) Setup(setupComponents []SetupComponent) {
	for _, setup := range setupComponents {
		setup()
	}
}

// Run application
func (app *mainApp) Run() {
	errChan := make(chan error)
	go app.globalSubscriber.Subscribe(errChan)
	go app.localSubscriber.Subscribe(errChan)
	go app.globalRPCServer.Serve(errChan)
	go app.localRPCServer.Serve(errChan)
	go app.router.Run(fmt.Sprintf(":%d", app.httpPort))
	app.logger.Println(fmt.Sprintf("Run at port %d", app.httpPort))
	app.liveness = true
	app.readiness = true
	// waiting for error
	err := <-errChan
	app.logger.Printf("[ERROR] %s", err)
	app.liveness = false
	app.readiness = false
	// waiting forever
	forever := make(chan bool)
	<-forever
}

// Logger get logger
func (app *mainApp) Logger() *log.Logger {
	return app.logger
}

// Router get router
func (app *mainApp) Router() *gin.Engine {
	return app.router
}

// GlobalPublisher get globalPublisher
func (app *mainApp) GlobalPublisher() transport.Publisher {
	return app.globalPublisher
}

// LocalPublisher get globalPublisher
func (app *mainApp) LocalPublisher() transport.Publisher {
	return app.localPublisher
}

// GlobalSubscriber get globalSubscriber
func (app *mainApp) GlobalSubscriber() transport.Subscriber {
	return app.globalSubscriber
}

// LocalSubscriber get globalSubscriber
func (app *mainApp) LocalSubscriber() transport.Subscriber {
	return app.localSubscriber
}

// GlobalRPCServer get globalRPCServer
func (app *mainApp) GlobalRPCServer() transport.RPCServer {
	return app.globalRPCServer
}

// LocalRPCServer get globalRPCServer
func (app *mainApp) LocalRPCServer() transport.RPCServer {
	return app.localRPCServer
}

// GlobalRPCClient get globalRPCClient
func (app *mainApp) GlobalRPCClient() transport.RPCClient {
	return app.globalRPCClient
}

// LocalRPCClient get globalRPCClient
func (app *mainApp) LocalRPCClient() transport.RPCClient {
	return app.localRPCClient
}

// Liveness get liveness of application
func (app *mainApp) Liveness() bool {
	return app.liveness
}

// SetLiveness set liveness of application
func (app *mainApp) SetLiveness(liveness bool) {
	app.livenessMux.Lock()
	defer app.livenessMux.Unlock()
	app.liveness = liveness
}

// Readiness get readiness of application
func (app *mainApp) Readiness() bool {
	return app.readiness
}

// SetReadiness set readiness of application
func (app *mainApp) SetReadiness(readiness bool) {
	app.readinessMux.Lock()
	defer app.readinessMux.Unlock()
	app.readiness = readiness
}

// CreateMainApp create application
func CreateMainApp(httpPort int, globalRmqConnectionString, localRmqConnectionString string) (app App) {
	logger := log.New(os.Stdout, "", log.LstdFlags)
	app = &mainApp{
		liveness:         false,
		readiness:        false,
		httpPort:         httpPort,
		logger:           logger,
		router:           gin.Default(),
		globalPublisher:  transport.CreateRmqPublisher(globalRmqConnectionString).SetLogger(logger),
		localPublisher:   transport.CreateRmqPublisher(localRmqConnectionString).SetLogger(logger),
		globalSubscriber: transport.CreateRmqSubscriber(globalRmqConnectionString).SetLogger(logger),
		localSubscriber:  transport.CreateRmqSubscriber(localRmqConnectionString).SetLogger(logger),
		globalRPCServer:  transport.CreateRmqRPCServer(globalRmqConnectionString).SetLogger(logger),
		localRPCServer:   transport.CreateRmqRPCServer(localRmqConnectionString).SetLogger(logger),
		globalRPCClient:  transport.CreateRmqRPCClient(globalRmqConnectionString).SetLogger(logger),
		localRPCClient:   transport.CreateRmqRPCClient(localRmqConnectionString).SetLogger(logger),
	}
	return app
}
