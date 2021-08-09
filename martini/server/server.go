package server

import (
	"os"
	"strconv"
	"web-servers/domain/server"
	"web-servers/martini/routes"

	"github.com/go-martini/martini"
)

type Server struct {
	App    *martini.Martini
	Router martini.Router
	Port   int
}

func New(port int) server.IServer {
	router := martini.NewRouter()
	server := &Server{
		App:    martini.New(),
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
	if err = os.Setenv("PORT", strconv.Itoa(s.Port)); err != nil {
		return err
	}

	routes.Init(s.App, s.Router)
	s.App.Run()

	return nil
}

func (s *Server) Stop() (err error) {
	return nil
}
