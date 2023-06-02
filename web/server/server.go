package server

import (
	"github.com/joaosoft/web-servers/domain/server"
	"github.com/joaosoft/web-servers/web/routes"

	"github.com/joaosoft/web"
)

type Server struct {
	App  *web.Server
	Port int
}

func New(port int) server.IServer {
	server := &Server{
		Port: port,
	}

	server.App, _ = web.NewServer()

	return server
}

func (s *Server) Start() (err error) {
	routes.Init(s.App)
	return s.App.Start()
}

func (s *Server) Stop() (err error) {
	return s.App.Stop()
}
