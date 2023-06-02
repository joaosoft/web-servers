package server

import (
	"context"
	"fmt"
	"github.com/joaosoft/web-servers/domain/server"
	"github.com/joaosoft/web-servers/httprouter/middlewares"
	"github.com/joaosoft/web-servers/httprouter/routes"
	"net/http"
)

type Server struct {
	App    *http.Server
	Router *middlewares.RouterWrapper
	Port   int
}

func New(port int) server.IServer {
	router := middlewares.NewRouterWrapper()
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
