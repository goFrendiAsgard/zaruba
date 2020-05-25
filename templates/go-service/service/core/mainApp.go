package core

import (
	"app/transport"
	"fmt"
	"log"
	"sync"

	"github.com/gin-gonic/gin"
)

// mainApp implementation, comply with App
type mainApp struct {
	readinessMux sync.Mutex
	readiness    bool
	livenessMux  sync.Mutex
	liveness     bool
	httpPort     int
	router       *gin.Engine
	rpcServers   []transport.RPCServer
	subscribers  []transport.Subscriber
	logger       *log.Logger
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
	for _, rpcServer := range app.rpcServers {
		go rpcServer.Serve(errChan)
	}
	for _, subscriber := range app.subscribers {
		go subscriber.Subscribe(errChan)
	}
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
func CreateMainApp(logger *log.Logger, router *gin.Engine, subscribers []transport.Subscriber, rpcServers []transport.RPCServer, httpPort int) (app App) {
	app = &mainApp{
		liveness:    false,
		readiness:   false,
		httpPort:    httpPort,
		logger:      logger,
		router:      router,
		subscribers: subscribers,
		rpcServers:  rpcServers,
	}
	return app
}
