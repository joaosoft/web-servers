package server

import (
	"fmt"
	"github.com/joaosoft/web-servers/domain/server"
	"github.com/joaosoft/web-servers/mux/routes"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	App    *http.Server
	Router *mux.Router
	Port   int
}

func New(port int) server.IServer {
	router := mux.NewRouter()
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
	return s.App.Close()
}
