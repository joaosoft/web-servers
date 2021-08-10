package server

import (
	"fmt"
	"web-servers/domain/server"
	"web-servers/fiber/routes"

	fiber "github.com/gofiber/fiber"
)

type Server struct {
	App  *fiber.App
	Port int
}

func New(port int) server.IServer {
	server := &Server{
		App:  fiber.New(&fiber.Settings{DisableStartupMessage: true}),
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
