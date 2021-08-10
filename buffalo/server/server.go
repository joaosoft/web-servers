package server

import (
	"fmt"
	"web-servers/buffalo/routes"
	"web-servers/domain/server"

	"github.com/gobuffalo/buffalo"
)

type Server struct {
	App  *buffalo.App
	Port int
}

func New(port int) server.IServer {
	server := &Server{
		App: buffalo.New(
			buffalo.Options{
				Addr:     fmt.Sprintf(":%d", port),
				LogLevel: "error",
			},
		),
		Port: port,
	}
	return server
}

func (s *Server) Start() (err error) {
	routes.Init(s.App)
	return s.App.Serve()
}

func (s *Server) Stop() (err error) {
	return s.App.Stop(nil)
}
