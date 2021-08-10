package server

import (
	"context"
	"web-servers/beego/routes"
	"web-servers/domain/server"

	"github.com/astaxie/beego"
)

type Server struct {
	App  *beego.App
	Port int
}

func New(port int) server.IServer {
	beego.BConfig.Listen.HTTPPort = port
	server := &Server{
		App:  beego.NewApp(),
		Port: port,
	}
	return server
}

func (s *Server) Start() (err error) {
	routes.Init(s.App)
	s.App.Run()
	return nil
}

func (s *Server) Stop() (err error) {
	return s.App.Server.Shutdown(context.Background())
}
