package server

import (
	"web-servers/domain/server"
	"web-servers/revel/app/tmp/run"

	"github.com/revel/revel"
)

type Server struct {
	App  *revel.ServerEngine
	Port int
}

func New(port int) server.IServer {
	server := &Server{
		App:  &revel.CurrentEngine,
		Port: port,
	}

	return server
}

func (s *Server) Start() (err error) {
	revel.Init("dev", "web-servers", "")
	run.Run(s.Port)
	return nil
}

func (s *Server) Stop() (err error) {
	_ = revel.StopServer(nil)
	return nil
}
