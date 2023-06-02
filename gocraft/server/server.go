package server

import (
	"context"
	"fmt"
	"github.com/joaosoft/web-servers/domain/server"
	"github.com/joaosoft/web-servers/gocraft/routes"
	"net/http"

	"github.com/gocraft/web"
)

type Server struct {
	App     *http.Server
	Router  *web.Router
	Port    int
	Context Context
}

type Context struct{}

func New(port int) server.IServer {
	context := Context{}
	router := web.New(context)
	server := &Server{
		Context: context,
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
	routes.Init(s.Context, s.Router)
	return s.App.ListenAndServe()
}

func (s *Server) Stop() (err error) {
	return s.App.Shutdown(context.Background())
}
