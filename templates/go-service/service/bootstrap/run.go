package bootstrap

import (
	"app/components"
	"fmt"
)

// Run everything
// TODO: change the implementation to fit your need
func Run(s *components.Setting) {
	go s.RPCServers.Main.Serve()
	go s.RPCServers.Secondary.Serve()
	go s.Subscribers.Main.Subscribe()
	go s.Router.Run(fmt.Sprintf(":%d", s.Ctx.Config.HTTPPort))
	s.Ctx.Config.Logger.Println(fmt.Sprintf("Run at port %d", s.Ctx.Config.HTTPPort))
	forever := make(chan bool)
	<-forever
}
