package server

import (
	"context"
	"fmt"
	"web-servers/domain/server"
	"web-servers/iris/routes"

	"github.com/kataras/iris/v12"
)

type Server struct {
	App  *iris.Application
	Port int
}

func New(port int) server.IServer {
	server := &Server{
		App:  iris.New(),
		Port: port,
	}

	return server
}

func (s *Server) Start() (err error) {
	routes.Init(s.App)
	return s.App.Run(
		iris.Addr(fmt.Sprintf(":%d", s.Port)),
		iris.WithOptimizations,
	)
}

func (s *Server) Stop() (err error) {
	return s.App.Shutdown(context.Background())
}
