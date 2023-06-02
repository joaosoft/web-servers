package server

import (
	"context"
	"fmt"
	"github.com/joaosoft/web-servers/domain/server"
	"github.com/joaosoft/web-servers/martini/routes"
	"net/http"

	"github.com/go-martini/martini"
)

type (
	Server struct {
		App
		Router martini.Router
		Port   int
	}
	App struct {
		*martini.Martini
		*http.Server
	}
)

func New(port int) server.IServer {
	m := martini.New()
	router := martini.NewRouter()
	server := &Server{
		App: App{
			Martini: m,
			Server: &http.Server{
				Addr:    fmt.Sprintf(":%d", port),
				Handler: m,
			},
		},
		Router: router,
		Port:   port,
	}

	server.App.Use(martini.Recovery())
	server.App.Use(martini.Static("public"))
	server.App.MapTo(router, (*martini.Routes)(nil))
	server.App.Action(router.Handle)
	return server
}

func (s *Server) Start() (err error) {
	routes.Init(s.App.Martini, s.Router)
	return s.App.Server.ListenAndServe()
}

func (s *Server) Stop() (err error) {
	return s.App.Server.Shutdown(context.Background())
}
