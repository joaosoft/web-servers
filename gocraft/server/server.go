package server

import (
	"context"
	"fmt"
	"net/http"
	"web-servers/domain/server"
	"web-servers/gocraft/routes"

	"github.com/gocraft/web"
)

type Server struct {
	App    *http.Server
	Router *web.Router
	Port   int
}

type Context struct{}

func New(port int) server.IServer {
	router := web.New(Context{})
	server := &Server{
		App: &http.Server{
			Addr:    fmt.Sprintf(":%d", port),
			Handler: router,
		},
		Router: router,
		Port:   port,
	}

	return server
}

func (s *Server) Start() (err error) {
	routes.Init(s.Router)
	return s.App.ListenAndServe()
}

func (s *Server) Stop() (err error) {
	return s.App.Shutdown(context.Background())
}
