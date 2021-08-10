package server

import (
	"fmt"
	"net/http"
	"web-servers/domain/server"
	"web-servers/mux/routes"

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
