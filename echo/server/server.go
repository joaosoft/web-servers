package server

import (
	"fmt"
	"github.com/joaosoft/web-servers/domain/server"
	"github.com/joaosoft/web-servers/echo/routes"

	"github.com/labstack/echo"
)

type Server struct {
	App  *echo.Echo
	Port int
}

func New(port int) server.IServer {
	server := &Server{
		App:  echo.New(),
		Port: port,
	}
	server.App.HideBanner = true

	return server
}

func (s *Server) Start() (err error) {
	routes.Init(s.App)
	return s.App.Start(fmt.Sprintf(":%d", s.Port))
}

func (s *Server) Stop() (err error) {
	return s.App.Close()
}
