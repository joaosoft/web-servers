package server

import (
	"fmt"
	"github.com/joaosoft/web-servers/domain/server"
	"github.com/joaosoft/web-servers/fiber/routes"

	fiber "github.com/gofiber/fiber/v2"
)

type Server struct {
	App  *fiber.App
	Port int
}

func New(port int) server.IServer {
	server := &Server{
		App:  fiber.New(fiber.Config{DisableStartupMessage: true}),
		Port: port,
	}

	return server
}

func (s *Server) Start() (err error) {
	routes.Init(s.App)
	return s.App.Listen(fmt.Sprintf(":%d", s.Port))
}

func (s *Server) Stop() (err error) {
	return s.App.Shutdown()
}
